package cos

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk/auth"
	"github.com/tencentyun/cos-go-sdk/http"
	"github.com/tencentyun/cos-go-sdk/utils"
)

const ENDPOINT = "http://web.file.myqcloud.com/files/v1/"
const EXPIRES = 60 * 5

// 腾讯云 COS API 客户端
type Client struct {
	AppId     string        // 项目 ID
	SecretId  string        // 项目 Secret ID
	SecretKey string        // 项目 Secret Key
	Timeout   time.Duration // 客户端超时时长
}

// 目录搜索模式, 细分为:
//     Both         : 列举目录和文件
//     DirectoryOnly: 仅列举目录
//     FileOnly     : 仅列举文件
type ListPattern int

const (
	Both          ListPattern = iota // 列举目录和文件
	DirectoryOnly                    // 仅列举目录
	FileOnly                         // 仅列举文件
)

// 字符串化目录搜索模式, 可能的返回值为: eListBoth, eListDirOnly, eListFileOnly
func (p ListPattern) String() string {
	switch p {
	case Both:
		return "eListBoth"
	case DirectoryOnly:
		return "eListDirOnly"
	case FileOnly:
		return "eListFileOnly"
	default:
		return "eListBoth"
	}
}

// 目录搜索排序方式, 细分为:
//     Asc:  正序
//     Desc: 反序
type ListOrder int

const (
	Asc  ListOrder = iota // 正序
	Desc                  // 反序
)

func (l ListOrder) Int() int {
	switch l {
	case Asc:
		return 0
	case Desc:
		return 1
	default:
		return 0
	}
}

// 构造 COS API 客户端
//     appId:     项目 ID
//     secretId:  项目的 Secret ID
//     secretKey: 项目的 Secret Key
func NewClient(appId, secretId, secretKey string) *Client {
	rand.Seed(time.Now().Unix())

	return &Client{
		AppId:     appId,
		SecretId:  secretId,
		SecretKey: secretKey,
		Timeout:   10000 * time.Millisecond,
	}
}

// 设置客户端 HTTP 请求超时时长
func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

// COS API 请求参数封装
type Request struct {
	Op string `json:"op" url:"op"`
}

// COS API 返回结果封装
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 目录创建操作的请求参数封装
type CreateFolderRequest struct {
	Request
	BizAttr string `json:"biz_attr"`
}

// 目录创建操作的返回结果封装
type CreateFolderResponse struct {
	Response
	Data struct {
		Ctime        string `json:"ctime"`
		ResourcePath string `json:"resource_path"`
	} `json:"data"`
}

// 目录更新操作的请求参数封装
type UpdateFolderRequest struct {
	Request
	BizAttr string `json:"biz_attr"`
}

// 目录更新操作的返回结果封装
type UpdateFolderResponse struct {
	Response
}

// 目录查询操作的请求参数封装
type StatFolderRequest struct {
	Request
}

// 目录查询操作的返回结果封装
type StatFolderResponse struct {
	Response
	Data struct {
		Name      string `json:"name"`
		BizAttr   string `json:"biz_attr"`
		FileSize  int64  `json:"filesize"`
		Sha       string `json:"sha"`
		Ctime     string `json:"ctime"`
		Mtime     string `json:"Mtime"`
		AccessUrl string `json:"access_url"`
	} `json:"data"`
}

// 目录删除操作的请求参数封装
type DeleteFolderRequest struct {
	Request
}

// 目录删除操作的返回结果封装
type DeleteFolderResponse struct {
	Response
}

// 目录列举及搜索操作的请求参数封装
type ListFolderRequest struct {
	Op      string `json:"op" url:"op"`
	Pattern string `json:"pattern" url:"pattern"`
	Context string `json:"context" url:"context"`
	Num     int    `json:"num" url:"num"`
	Order   int    `json:"order" url:"order"`
}

// 目录列举及搜索操作的返回结果封装
type ListFolderResponse struct {
	Response
	Data struct {
		Context   string `json:"context"`
		HasMore   bool   `json:"has_more"`
		DirCount  int    `json:"dircount"`
		FileCount int    `json:"filecount"`
		Infos     []struct {
			Name      string `json:"name"`
			BizAttr   string `json:"biz_attr"`
			FileSize  int64  `json:"filesize"`
			FileLen   int64  `json:"filelen"`
			Sha       string `json:"sha"`
			Ctime     string `json:"ctime"`
			Mtime     string `json:"mtime"`
			AccessUrl string `json:"access_url"`
		} `json:"infos"`
	} `json:"data"`
}

// 文件上传操作的请求参数封装
type UploadFileRequest struct {
	Op          string `json:"op"`
	FileContent string `json:"filecontent"`
	Sha         string `json:"sha"`
	BizAttr     string `json:"biz_attr"`
}

// 文件上传操作的返回结果封装
type UploadFileResponse struct {
	Response
	Data struct {
		AccessUrl    string `json:"access_url"`
		Url          string `json:"url"`
		ResourcePath string `json:"resource_path"`
	} `json:"data"`
}

// 初次文件分片上传操作的请求参数封装
type PrepareToUploadSliceRequest struct {
	Op        string `json:"op"`
	FileSize  int64  `json:"filesize"`
	Sha       string `json:"sha"`
	BizAttr   string `json:"biz_attr"`
	Session   string `json:"session"`
	SliceSize int64  `json:"slice_size"`
}

// 后续文件分片上传操作的请求参数封装
type UploadSliceRequest struct {
	Op          string `json:"op"`
	FileContent string `json:"filecontent"`
	Sha         string `json:"sha"`
	Session     string `json:"session"`
	SliceSize   int64  `json:"slice_size"`
	Offset      int64  `json:"offset"`
}

// 文件分片上传操作的返回结果封装
type UploadSliceResponse struct {
	Response
	Data struct {
		Session      string `json:"session"`
		Offset       int64  `json:"offset"`
		SliceSize    int64  `json:"slice_size"`
		AccessUrl    string `json:"access_url"`
		Url          string `json:"url"`
		ResourcePath string `json:"resource_path"`
	} `json:"data"`
}

// 文件属性更新操作的请求参数封装
type UpdateFileRequest struct {
	Request
	BizAttr string `json:"biz_attr"`
}

// 文件属性更新操作的返回结果封装
type UpdateFileResponse struct {
	Response
}

// 文件查询操作的请求参数封装
type StatFileRequest struct {
	Request
}

// 文件查询操作的返回结果封装
type StatFileResponse struct {
	Response
	Data struct {
		Name      string `json:"name"`
		BizAttr   string `json:"biz_attr"`
		FileSize  string `json:"filesize"`
		Sha       string `json:"sha"`
		Ctime     string `json:"ctime"`
		Mtime     string `json:"Mtime"`
		AccessUrl string `json:"access_url"`
	} `json:"data"`
}

// 文件删除操作的请求参数封装
type DeleteFileRequest struct {
	Request
}

// 文件删除操作的返回结果封装
type DeleteFileResponse struct {
	Response
}

func (c *Client) generateFileId(bucket, path string) string {
	resource := fmt.Sprintf("/%s/%s%s", c.AppId, bucket, path)
	return resource
}

func (c *Client) generateResourceUrl(bucket, path string) string {
	resource := fmt.Sprintf("%s%s/%s%s", ENDPOINT, c.AppId, bucket, path)
	return resource
}

func (c *Client) validateFolderPath(path string) string {
	validFolderPath := path

	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		validFolderPath = "/" + path
	}
	if !strings.HasSuffix(path, "/") {
		validFolderPath += "/"
	}

	return validFolderPath
}

func (c *Client) validateFilePath(path string) string {
	validFilePath := path

	if path == "" {
		return "/"
	}
	if !strings.HasPrefix(path, "/") {
		validFilePath = "/" + path
	}

	return validFilePath
}

// 创建目录
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//     mode:    目录创建模式, 当目录路径冲突时是否覆盖原节点, OverWrite 为覆盖, Keep 为不覆盖
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.CreateFolder("cosdemo", "/hello", "hello", cos.OverWrite)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nCtime:", res.Data.Ctime,
//         "\nResource Path:", res.Data.ResourcePath)
//
func (c *Client) CreateFolder(bucket, path, bizAttr string) (*CreateFolderResponse, error) {
	var cosRequest CreateFolderRequest
	var cosResponse CreateFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "create"
	cosRequest.BizAttr = bizAttr

	httpRequest := http.Request{
		Method:      "POST",
		Uri:         resource,
		Timeout:     c.Timeout,
		ContentType: "application/json",
		Body:        cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 更新目录属性
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.UpdateFolder("cosdemo", "/hello", "hello-new-attr")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) UpdateFolder(bucket, path, bizAttr string) (*UpdateFolderResponse, error) {
	var cosRequest UpdateFolderRequest
	var cosResponse UpdateFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path))
	fileId := c.generateFileId(bucket, encodedPath)
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, "", strNow, strRand, fileId)
	sign := signer.SignOnce(c.SecretKey)

	cosRequest.Op = "update"
	cosRequest.BizAttr = bizAttr

	httpRequest := http.Request{
		Method:      "POST",
		Uri:         resource,
		Timeout:     c.Timeout,
		ContentType: "application/json",
		Body:        cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 查询目录属性
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//    client := cos.NewClient(appId, secretId, secretKey)
//
//    res, err := client.StatFolder("cosdemo", "/hello")
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    fmt.Println("Code:", res.Code,
//        "\nMessage:", res.Message,
//        "\nName:", res.Data.Name,
//        "\nBizAttr:", res.Data.BizAttr,
//        "\nCtime:", res.Data.Ctime,
//        "\nMtime:", res.Data.Mtime,
//        "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) StatFolder(bucket, path string) (*StatFolderResponse, error) {
	var cosRequest StatFolderRequest
	var cosResponse StatFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "stat"

	httpRequest := http.Request{
		Method:      "GET",
		Uri:         resource,
		Timeout:     c.Timeout,
		QueryString: cosRequest.Request,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 删除目录
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.DeleteFolder("cosdemo", "/hello")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) DeleteFolder(bucket, path string) (*DeleteFolderResponse, error) {
	var cosRequest DeleteFolderRequest
	var cosResponse DeleteFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path))
	fileId := c.generateFileId(bucket, encodedPath)
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, "", strNow, strRand, fileId)
	sign := signer.SignOnce(c.SecretKey)

	cosRequest.Op = "delete"

	httpRequest := http.Request{
		Method:      "POST",
		Uri:         resource,
		Timeout:     c.Timeout,
		ContentType: "application/json",
		Body:        cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 列举目录和文件
//     bucket:  Bucket 名称
//     path:    目录路径
//     context: Context 用于翻页, 需要往前/往后翻页需透传回来
//     pattern: 目录列举模式, 可选值为 Both(列举目录和文件), DirectoryOnly(仅列举目录), FileOnly(仅列举文件)
//     num:     本次拉取的目录和文件总数
//     order:   目录列举排序规则, 可选值为 Asc(正序), Desc(反序)
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.ListFolder("cosdemo", "/hello", "", cos.Both, 100, cos.Asc)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nContext:", res.Data.Context,
//         "\nHasMore:", res.Data.HasMore,
//         "\nDirCount:", res.Data.DirCount,
//         "\nFileCount:", res.Data.FileCount,
//     )
//
//     fmt.Println("*************************************")
//     for _, info := range res.Data.Infos {
//         fmt.Println("Name:", info.Name,
//             "\nBizAttr:", info.BizAttr,
//             "\nFileSize:", info.FileSize,
//             "\nFileLen:", info.FileLen,
//             "\nSha:", info.Sha,
//             "\nCtime:", info.Ctime,
//             "\nMtime:", info.Mtime,
//             "\nAccess URL:", info.AccessUrl,
//         )
//         fmt.Println("*************************************")
//     }
//
func (c *Client) ListFolder(bucket, path, context string, pattern ListPattern, num int, order ListOrder) (*ListFolderResponse, error) {
	var cosRequest ListFolderRequest
	var cosResponse ListFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "list"
	cosRequest.Pattern = pattern.String()
	cosRequest.Context = context
	cosRequest.Order = order.Int()
	cosRequest.Num = num

	httpRequest := http.Request{
		Method:      "GET",
		Uri:         resource,
		Timeout:     c.Timeout,
		QueryString: cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 前缀搜索目录和文件
//     bucket:  Bucket 名称
//     path:    目录路径
//     prefix:  搜索前缀
//     context: Context 用于翻页, 需要往前/往后翻页需透传回来
//     pattern: 目录列举模式, 可选值为 Both(列举目录和文件), DirectoryOnly(仅列举目录), FileOnly(仅列举文件)
//     num:     本次拉取的目录和文件总数
//     order:   目录列举排序规则, 可选值为 Asc(正序), Desc(反序)
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.PrefixSearch("cosdemo", "/hello", "A", "", cos.Both, 100, cos.Asc)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nContext:", res.Data.Context,
//         "\nHasMore:", res.Data.HasMore,
//         "\nDirCount:", res.Data.DirCount,
//         "\nFileCount:", res.Data.FileCount,
//     )
//
//     fmt.Println("*************************************")
//     for _, info := range res.Data.Infos {
//         fmt.Println("Name:", info.Name,
//             "\nBizAttr:", info.BizAttr,
//             "\nFileSize:", info.FileSize,
//             "\nFileLen:", info.FileLen,
//             "\nSha:", info.Sha,
//             "\nCtime:", info.Ctime,
//             "\nMtime:", info.Mtime,
//             "\nAccess URL:", info.AccessUrl,
//         )
//         fmt.Println("*************************************")
//     }
//
func (c *Client) PrefixSearch(bucket, path, prefix, context string, pattern ListPattern, num int, order ListOrder) (*ListFolderResponse, error) {
	var cosRequest ListFolderRequest
	var cosResponse ListFolderResponse

	encodedPath := utils.UrlEncode(c.validateFolderPath(path) + prefix)
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "list"
	cosRequest.Pattern = pattern.String()
	cosRequest.Context = context
	cosRequest.Order = order.Int()
	cosRequest.Num = num

	httpRequest := http.Request{
		Method:      "GET",
		Uri:         resource,
		Timeout:     c.Timeout,
		QueryString: cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 上传文件
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.UploadFile("cosdemo", "/hello/hello.txt", "/users/new.txt", "file attr")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadFile(bucket, dstPath, srcPath, bizAttr string) (*UploadFileResponse, error) {
	fileContent, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return nil, err
	}

	return c.UploadChunk(bucket, dstPath, fileContent, bizAttr)
}

// 上传内存块至云端
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     chunk:   本地内存块
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.UploadChunk("cosdemo", "/hello/hello.txt", []bytes("Hello"), "test hello")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadChunk(bucket, dstPath string, chunk []byte, bizAttr string) (*UploadFileResponse, error) {
	var cosRequest UploadFileRequest
	var cosResponse UploadFileResponse

	sha, err := utils.HashBufferWithSha1(chunk)
	if err != nil {
		return nil, err
	}

	encodedPath := utils.UrlEncode(c.validateFilePath(dstPath))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "upload"
	cosRequest.FileContent = string(chunk)
	cosRequest.Sha = sha
	cosRequest.BizAttr = bizAttr

	httpRequest := http.Request{
		Method:    "POST",
		Uri:       resource,
		Timeout:   c.Timeout,
		Body:      cosRequest,
		Multipart: true,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 大文件分片上传
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//     session: 唯一标识此文件传输过程的id, 由后台下发, 调用方透传
//     sliceSize: 分片大小, 用户可以根据网络状况自行设置
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.UploadSlice("cosdemo", "/hello/hello.bin", "/users/bigfile.bin", "file attr", "", 512 * 1024)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadSlice(bucket, dstPath, srcPath, bizAttr, session string, sliceSize int64) (*UploadSliceResponse, error) {
	preparing, err := c.PrepareToUploadSlice(bucket, dstPath, srcPath, bizAttr, session, sliceSize)
	if err != nil {
		return nil, err
	}
	if preparing.Code != 0 {
		return nil, fmt.Errorf("%s", preparing.Message)
	}

	if len(preparing.Data.Url) != 0 { // 秒传命中
		return preparing, nil
	}

	var offset int64
	if preparing.Data.SliceSize != 0 {
		sliceSize = preparing.Data.SliceSize
	}
	if preparing.Data.Offset != 0 {
		offset = preparing.Data.Offset
	}
	if preparing.Data.Session != "" {
		session = preparing.Data.Session
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		return nil, err
	}
	defer srcFile.Close()
	fileInfo, err := srcFile.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	srcFile.Seek(preparing.Data.Offset, 0)

	for fileSize > offset {
		var rawBytes []byte
		if fileSize-offset > sliceSize {
			rawBytes = make([]byte, sliceSize)
			nbytes, err := io.ReadFull(srcFile, rawBytes)
			if int64(nbytes) != sliceSize || err != nil {
				return nil, err
			}
		} else {
			leftSize := fileSize - offset
			rawBytes = make([]byte, leftSize)
			nbytes, err := io.ReadFull(srcFile, rawBytes)
			if int64(nbytes) != leftSize || err != nil {
				return nil, err
			}
		}

		continuing, err := c.ContinueUploadingSliceData(bucket, dstPath, rawBytes, session, offset)
		if err != nil {
			return nil, err
		}

		if len(continuing.Data.Url) > 0 {
			return continuing, nil
		} else {
			offset += sliceSize
		}
	}

	return preparing, nil
}

// 大文件分片上传(首次上传)
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//     session: 唯一标识此文件传输过程的id, 由后台下发, 调用方透传
//     sliceSize: 分片大小, 用户可以根据网络状况自行设置
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.PrepareToUploadSlice("cosdemo", "/hello/hello.bin", "/Users/bigfile.bin", "file attr", "", 512 * 1024)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nSession:", res.Session,
//         "\nOffset:", res.Offset,
//         "\nSliceSize:", res.SliceSize,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) PrepareToUploadSlice(bucket, dstPath, srcPath, bizAttr, session string, sliceSize int64) (*UploadSliceResponse, error) {
	var cosRequest PrepareToUploadSliceRequest
	var cosResponse UploadSliceResponse

	sha, size, err := utils.HashFileWithSha1(srcPath)
	if err != nil {
		return nil, err
	}

	encodedPath := utils.UrlEncode(c.validateFilePath(dstPath))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "upload_slice"
	cosRequest.FileSize = size
	cosRequest.Sha = sha
	cosRequest.BizAttr = bizAttr
	cosRequest.Session = session
	cosRequest.SliceSize = sliceSize

	httpRequest := http.Request{
		Method:    "POST",
		Uri:       resource,
		Timeout:   c.Timeout,
		Body:      cosRequest,
		Multipart: true,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 大文件分片上传(后续上传)
//     bucket:    Bucket 名称
//     dstpath:   目的文件路径, 须设置为腾讯云端文件路径
//     dataSlice: 本地文件路径
//     session:   唯一标识此文件传输过程的id, 由后台下发, 调用方透传
//     offset:    本次分片位移
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.ContinueUploadingSliceData("cosdemo", "/hello/bigfile.bin", []byte("data.bin"), "c0bd94d0-3956-4664-b99f-2658eef8e5f5+CpcFEtEHAA==", 1045248)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nSession:", res.Session,
//         "\nOffset:", res.Offset,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) ContinueUploadingSliceData(bucket, dstPath string, dataSlice []byte, session string, offset int64) (*UploadSliceResponse, error) {
	var cosRequest UploadSliceRequest
	var cosResponse UploadSliceResponse

	sha, err := utils.HashBufferWithSha1(dataSlice)
	if err != nil {
		return nil, err
	}

	encodedPath := utils.UrlEncode(c.validateFilePath(dstPath))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "upload_slice"
	cosRequest.FileContent = string(dataSlice)
	cosRequest.Sha = sha
	cosRequest.Session = session
	cosRequest.Offset = offset

	httpRequest := http.Request{
		Method:    "POST",
		Uri:       resource,
		Timeout:   c.Timeout,
		Body:      cosRequest,
		Multipart: true,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 更新文件属性
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.UpdateFile("cosdemo", "/hello", "hello-new-attr")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) UpdateFile(bucket, path, bizAttr string) (*UpdateFileResponse, error) {
	var cosRequest UpdateFileRequest
	var cosResponse UpdateFileResponse

	encodedPath := utils.UrlEncode(c.validateFilePath(path))
	fileId := c.generateFileId(bucket, encodedPath)
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, "", strNow, strRand, fileId)
	sign := signer.SignOnce(c.SecretKey)

	cosRequest.Op = "update"
	cosRequest.BizAttr = bizAttr

	httpRequest := http.Request{
		Method:      "POST",
		Uri:         resource,
		Timeout:     c.Timeout,
		ContentType: "application/json",
		Body:        cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 查询文件属性
//     bucket:  Bucket 名称
//     path:    文件路径
//
// 示例:
//
//    client := cos.NewClient(appId, secretId, secretKey)
//
//    res, err := client.StatFile("cosdemo", "/hello/new.txt")
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    fmt.Println("Code:", res.Code,
//        "\nMessage:", res.Message,
//        "\nName:", res.Data.Name,
//        "\nBizAttr:", res.Data.BizAttr,
//        "\nFileSize:", res.Data.FileSize,
//        "\nSha:", res.Data.Sha,
//        "\nCtime:", res.Data.Ctime,
//        "\nMtime:", res.Data.Mtime,
//        "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) StatFile(bucket, path string) (*StatFileResponse, error) {
	var cosRequest StatFileRequest
	var cosResponse StatFileResponse

	encodedPath := utils.UrlEncode(c.validateFilePath(path))
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strExpires := fmt.Sprintf("%d", now+EXPIRES)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, strExpires, strNow, strRand, "")
	sign := signer.Sign(c.SecretKey)

	cosRequest.Op = "stat"

	httpRequest := http.Request{
		Method:      "GET",
		Uri:         resource,
		Timeout:     c.Timeout,
		QueryString: cosRequest.Request,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}

// 删除文件
//     bucket:  Bucket 名称
//     path:    文件路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     res, err := client.DeleteFile("cosdemo", "/hello/new.txt")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) DeleteFile(bucket, path string) (*DeleteFileResponse, error) {
	var cosRequest DeleteFileRequest
	var cosResponse DeleteFileResponse

	encodedPath := utils.UrlEncode(c.validateFilePath(path))
	fileId := c.generateFileId(bucket, encodedPath)
	resource := c.generateResourceUrl(bucket, encodedPath)

	now := time.Now().Unix()
	strNow := fmt.Sprintf("%d", now)
	strRand := fmt.Sprintf("%d", rand.Int31())
	signer := auth.NewSignature(c.AppId, bucket, c.SecretId, "", strNow, strRand, fileId)
	sign := signer.SignOnce(c.SecretKey)

	cosRequest.Op = "delete"

	httpRequest := http.Request{
		Method:      "POST",
		Uri:         resource,
		Timeout:     c.Timeout,
		ContentType: "application/json",
		Body:        cosRequest,
	}
	httpRequest.AddHeader("Authorization", sign)
	httpResponse, err := httpRequest.Do()
	if err != nil {
		return nil, err
	}

	err = httpResponse.Body.Unmarshal(&cosResponse)
	if err != nil {
		return nil, err
	}

	return &cosResponse, nil
}
