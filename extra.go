package cos

// 目录创建操作异步接口的返回结果封装
type CreateFolderAsyncResponse struct {
	Response *CreateFolderResponse
	Error    error
}

// 目录更新操作异步接口的返回结果封装
type UpdateFolderAsyncResponse struct {
	Response *UpdateFolderResponse
	Error    error
}

// 目录查询操作异步接口的返回结果封装
type StatFolderAsyncResponse struct {
	Response *StatFolderResponse
	Error    error
}

// 目录删除操作异步接口的返回结果封装
type DeleteFolderAsyncResponse struct {
	Response *DeleteFolderResponse
	Error    error
}

// 目录列举及搜索操作异步接口的返回结果封装
type ListFolderAsyncResponse struct {
	Response *ListFolderResponse
	Error    error
}

// 文件上传操作异步接口的返回结果封装
type UploadFileAsyncResponse struct {
	Response *UploadFileResponse
	Error    error
}

// 文件分片上传操作异步接口的返回结果封装
type UploadSliceAsyncResponse struct {
	Response *UploadSliceResponse
	Error    error
}

// 文件属性更新操作异步接口的返回结果封装
type UpdateFileAsyncResponse struct {
	Response *UpdateFileResponse
	Error    error
}

// 文件查询操作异步接口的返回结果封装
type StatFileAsyncResponse struct {
	Response *StatFileResponse
	Error    error
}

// 文件删除操作异步接口的返回结果封装
type DeleteFileAsyncResponse struct {
	Response *DeleteFileResponse
	Error    error
}

// 创建目录异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.CreateFolderAsync("cosdemo", "/hello", "hello")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nCtime:", res.Data.Ctime,
//         "\nResource Path:", res.Data.ResourcePath)
//
func (c *Client) CreateFolderAsync(bucket, path, bizAttr string) <-chan *CreateFolderAsyncResponse {
	ch := make(chan *CreateFolderAsyncResponse, 1)
	go func() {
		var asyncResponse CreateFolderAsyncResponse
		response, err := c.CreateFolder(bucket, path, bizAttr)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 更新目录属性异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.UpdateFolderAsync("cosdemo", "/hello", "hello-new-attr")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) UpdateFolderAsync(bucket, path, bizAttr string) <-chan *UpdateFolderAsyncResponse {
	ch := make(chan *UpdateFolderAsyncResponse, 1)
	go func() {
		var asyncResponse UpdateFolderAsyncResponse
		response, err := c.UpdateFolder(bucket, path, bizAttr)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 查询目录属性异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.StatFolderAsync("cosdemo", "/hello")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//        "\nMessage:", res.Message,
//        "\nName:", res.Data.Name,
//        "\nBizAttr:", res.Data.BizAttr,
//        "\nCtime:", res.Data.Ctime,
//        "\nMtime:", res.Data.Mtime,
//        "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) StatFolderAsync(bucket, path string) <-chan *StatFolderAsyncResponse {
	ch := make(chan *StatFolderAsyncResponse, 1)
	go func() {
		var asyncResponse StatFolderAsyncResponse
		response, err := c.StatFolder(bucket, path)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 删除目录异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.DeleteFolderAsync("cosdemo", "/hello")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) DeleteFolderAsync(bucket, path string) <-chan *DeleteFolderAsyncResponse {
	ch := make(chan *DeleteFolderAsyncResponse, 1)
	go func() {
		var asyncResponse DeleteFolderAsyncResponse
		response, err := c.DeleteFolder(bucket, path)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 列举目录和文件异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
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
//     resAsync := client.ListFolderAsync("cosdemo", "/hello", "", cos.Both, 100, cos.Asc)
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
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
func (c *Client) ListFolderAsync(bucket, path, context string, pattern ListPattern, num int, order ListOrder) <-chan *ListFolderAsyncResponse {
	ch := make(chan *ListFolderAsyncResponse, 1)
	go func() {
		var asyncResponse ListFolderAsyncResponse
		response, err := c.ListFolder(bucket, path, context, pattern, num, order)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 前缀搜索目录和文件异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
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
//     resAsync := client.PrefixSearchAsync("cosdemo", "/hello", "A", "", cos.Both, 100, cos.Asc)
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
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
func (c *Client) PrefixSearchAsync(bucket, path, prefix, context string, pattern ListPattern, num int, order ListOrder) <-chan *ListFolderAsyncResponse {
	ch := make(chan *ListFolderAsyncResponse, 1)
	go func() {
		var asyncResponse ListFolderAsyncResponse
		response, err := c.PrefixSearch(bucket, path, prefix, context, pattern, num, order)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 上传文件异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.UploadFileAsync("cosdemo", "/hello/hello.txt", "/users/new.txt", "file attr")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadFileAsync(bucket, dstPath, srcPath, bizAttr string) <-chan *UploadFileAsyncResponse {
	ch := make(chan *UploadFileAsyncResponse, 1)
	go func() {
		var asyncResponse UploadFileAsyncResponse
		response, err := c.UploadFile(bucket, dstPath, srcPath, bizAttr)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 上传内存块至云端异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     chunk:   本地内存块
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.UploadChunkAsync("cosdemo", "/hello/hello.txt", []bytes("Hello"), "test hello")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadChunkAsync(bucket, dstPath string, chunk []byte, bizAttr string) <-chan *UploadFileAsyncResponse {
	ch := make(chan *UploadFileAsyncResponse, 1)
	go func() {
		var asyncResponse UploadFileAsyncResponse
		response, err := c.UploadChunk(bucket, dstPath, chunk, bizAttr)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 大文件分片上传异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
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
//     resAsync := client.UploadSliceAsync("cosdemo", "/hello/hello.bin", "/users/bigfile.bin", "file attr", "", 512 * 1024)
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nUrl:", res.Data.Url,
//         "\nResourcePath:", res.Data.ResourcePath,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) UploadSliceAsync(bucket, dstPath, srcPath, bizAttr, session string, sliceSize int64) <-chan *UploadSliceAsyncResponse {
	ch := make(chan *UploadSliceAsyncResponse, 1)
	go func() {
		var asyncResponse UploadSliceAsyncResponse
		response, err := c.UploadSlice(bucket, dstPath, srcPath, bizAttr, session, sliceSize)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 更新文件属性异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resAsync := client.UpdateFileAsync("cosdemo", "/hello", "hello-new-attr")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) UpdateFileAsync(bucket, path, bizAttr string) <-chan *UpdateFileAsyncResponse {
	ch := make(chan *UpdateFileAsyncResponse, 1)
	go func() {
		var asyncResponse UpdateFileAsyncResponse
		response, err := c.UpdateFile(bucket, path, bizAttr)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 查询文件属性异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    文件路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resChan := client.StatFileAsync("cosdemo", "/hello/new.txt")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message,
//         "\nName:", res.Data.Name,
//         "\nBizAttr:", res.Data.BizAttr,
//         "\nFileSize:", res.Data.FileSize,
//         "\nSha:", res.Data.Sha,
//         "\nCtime:", res.Data.Ctime,
//         "\nMtime:", res.Data.Mtime,
//         "\nAccess Url:", res.Data.AccessUrl)
//
func (c *Client) StatFileAsync(bucket, path string) <-chan *StatFileAsyncResponse {
	ch := make(chan *StatFileAsyncResponse, 1)
	go func() {
		var asyncResponse StatFileAsyncResponse
		response, err := c.StatFile(bucket, path)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 删除文件异步接口, 调用该函数可立即返回一管道, 调用方可在后续代码逻辑中读取该管道.
//     bucket:  Bucket 名称
//     path:    文件路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//
//     resChan := client.DeleteFile("cosdemo", "/hello/new.txt")
//
//     // Do your other work here
//
//     resAsync := <- resChan
//
//     if resAsync.Error != nil {
//         fmt.Println(resAsync.Error)
//         return
//     }
//
//     res := resAsync.Response
//     fmt.Println("Code:", res.Code,
//         "\nMessage:", res.Message)
//
func (c *Client) DeleteFileAsync(bucket, path string) <-chan *DeleteFileAsyncResponse {
	ch := make(chan *DeleteFileAsyncResponse, 1)
	go func() {
		var asyncResponse DeleteFileAsyncResponse
		response, err := c.DeleteFile(bucket, path)
		asyncResponse.Response = response
		asyncResponse.Error = err
		ch <- &asyncResponse
	}()
	return ch
}

// 基于回调函数的目录创建接口, 调用该函数可立即返回, 目录创建完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.CreateFolderWithCallback("cosdemo", "/hello123", "hello",
//         func(res *cos.CreateFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message,
//                 "\nCtime:", res.Data.Ctime,
//                 "\nResource Path:", res.Data.ResourcePath)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) CreateFolderWithCallback(bucket, path, bizAttr string, callback func(*CreateFolderResponse, error)) {
	go func() {
		response, err := c.CreateFolder(bucket, path, bizAttr)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的目录属性更新接口, 调用该函数可立即返回, 目录属性更新完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.UpdateFolderWithCallback("cosdemo", "/hello123", "hello-new-attr",
//         func(res *cos.UpdateFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) UpdateFolderWithCallback(bucket, path, bizAttr string, callback func(*UpdateFolderResponse, error)) {
	go func() {
		response, err := c.UpdateFolder(bucket, path, bizAttr)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的目录属性查询接口, 调用该函数可立即返回, 目录属性查询完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//     client.StatFolderWithCallback("cosdemo", "/hello",
//         func(res *cos.StatFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message,
//                 "\nName:", res.Data.Name,
//                 "\nBizAttr:", res.Data.BizAttr,
//                 "\nCtime:", res.Data.Ctime,
//                 "\nMtime:", res.Data.Mtime,
//                 "\nAccess Url:", res.Data.AccessUrl)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) StatFolderWithCallback(bucket, path string, callback func(*StatFolderResponse, error)) {
	go func() {
		response, err := c.StatFolder(bucket, path)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的目录删除接口, 调用该函数可立即返回, 目录删除完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.DeleteFolderWithCallback("cosdemo", "/hello123",
//         func(res *cos.DeleteFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) DeleteFolderWithCallback(bucket, path string, callback func(*DeleteFolderResponse, error)) {
	go func() {
		response, err := c.DeleteFolder(bucket, path)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的前缀搜索目录和文件接口, 调用该函数可立即返回, 搜索完成后调用回调函数.
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
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.ListFolderWithCallback("cosdemo", "/hello", "", cos.Both, 100, cos.Asc,
//         func(res *cos.ListFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message,
//                 "\nContext:", res.Data.Context,
//                 "\nHasMore:", res.Data.HasMore,
//                 "\nDirCount:", res.Data.DirCount,
//                 "\nFileCount:", res.Data.FileCount,
//             )
//
//             fmt.Println("*************************************")
//             for _, info := range res.Data.Infos {
//                 fmt.Println("Name:", info.Name,
//                     "\nBizAttr:", info.BizAttr,
//                     "\nFileSize:", info.FileSize,
//                     "\nFileLen:", info.FileLen,
//                     "\nSha:", info.Sha,
//                     "\nCtime:", info.Ctime,
//                     "\nMtime:", info.Mtime,
//                     "\nAccess URL:", info.AccessUrl,
//                 )
//                 fmt.Println("*************************************")
//             }
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) ListFolderWithCallback(bucket, path, context string, pattern ListPattern, num int, order ListOrder, callback func(*ListFolderResponse, error)) {
	go func() {
		response, err := c.ListFolder(bucket, path, context, pattern, num, order)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的搜索目录和文件接口, 调用该函数可立即返回, 搜索完成后调用回调函数.
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
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.PrefixSearchWithCallback("cosdemo", "/hello", "", cos.Both, 100, cos.Asc,
//         func(res *cos.ListFolderResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message,
//                 "\nContext:", res.Data.Context,
//                 "\nHasMore:", res.Data.HasMore,
//                 "\nDirCount:", res.Data.DirCount,
//                 "\nFileCount:", res.Data.FileCount,
//             )
//
//             fmt.Println("*************************************")
//             for _, info := range res.Data.Infos {
//                 fmt.Println("Name:", info.Name,
//                     "\nBizAttr:", info.BizAttr,
//                     "\nFileSize:", info.FileSize,
//                     "\nFileLen:", info.FileLen,
//                     "\nSha:", info.Sha,
//                     "\nCtime:", info.Ctime,
//                     "\nMtime:", info.Mtime,
//                     "\nAccess URL:", info.AccessUrl,
//                 )
//                 fmt.Println("*************************************")
//             }
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) PrefixSearchWithCallback(bucket, path, prefix, context string, pattern ListPattern, num int, order ListOrder, callback func(*ListFolderResponse, error)) {
	go func() {
		response, err := c.PrefixSearch(bucket, path, prefix, context, pattern, num, order)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的文件上传接口, 调用该函数可立即返回, 文件上传完成后调用回调函数.
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//
// client := cos.NewClient(appId, secretId, secretKey)
// var wg = sync.WaitGroup{}
//
// wg.Add(1)
//
// fmt.Println("Uploading...")
//
// client.UploadFileWithCallback("cosdemo",
//     "/hello/goasguen-cernvm-2015.pptx",
//     "/Users/goasguen-cernvm-2015.pptx",
//     "goasguen-cernvm-2015.pptx",
//     func(res *cos.UploadFileResponse, err error) {
//         if err != nil {
//             fmt.Println(err)
//             return
//         }
//
//         fmt.Println("Code:", res.Code,
//             "\nMessage:", res.Message,
//             "\nUrl:", res.Data.Url,
//             "\nResourcePath:", res.Data.ResourcePath,
//             "\nAccess Url:", res.Data.AccessUrl)
//
//         wg.Done()
//     })
//
// wg.Wait()
//
// fmt.Println("Uploaded...")
//
func (c *Client) UploadFileWithCallback(bucket, dstPath, srcPath, bizAttr string, callback func(*UploadFileResponse, error)) {
	go func() {
		response, err := c.UploadFile(bucket, dstPath, srcPath, bizAttr)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的内存块上传接口, 调用该函数可立即返回, 文件上传完成后调用回调函数.
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     chunk:   本地内存块
//     bizAttr: 文件属性, 由业务端维护
//
// 示例:
//
// client := cos.NewClient(appId, secretId, secretKey)
// var wg = sync.WaitGroup{}
//
// wg.Add(1)
//
// fmt.Println("Uploading...")
//
// client.UploadChunkWithCallback("cosdemo",
//     "/hello/goasguen-cernvm-2015.pptx",
//     []byte("file...bin"),
//     "goasguen-cernvm-2015.pptx",
//     func(res *cos.UploadChunkResponse, err error) {
//         if err != nil {
//             fmt.Println(err)
//             return
//         }
//
//         fmt.Println("Code:", res.Code,
//             "\nMessage:", res.Message,
//             "\nUrl:", res.Data.Url,
//             "\nResourcePath:", res.Data.ResourcePath,
//             "\nAccess Url:", res.Data.AccessUrl)
//
//         wg.Done()
//     })
//
// wg.Wait()
//
// fmt.Println("Uploaded...")
//
func (c *Client) UploadChunkWithCallback(bucket, dstPath string, chunk []byte, bizAttr string, callback func(*UploadFileResponse, error)) {
	go func() {
		response, err := c.UploadChunk(bucket, dstPath, chunk, bizAttr)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的大文件分片上传接口, 调用该函数可立即返回, 文件分片上传完成后调用回调函数.
//     bucket:  Bucket 名称
//     dstpath: 目的文件路径, 须设置为腾讯云端文件路径
//     srcPath: 本地文件路径
//     bizAttr: 文件属性, 由业务端维护
//     session: 唯一标识此文件传输过程的id, 由后台下发, 调用方透传
//     sliceSize: 分片大小, 用户可以根据网络状况自行设置
//
// 示例:
//
// client := cos.NewClient(appId, secretId, secretKey)
// var wg = sync.WaitGroup{}
//
// wg.Add(1)
//
// fmt.Println("Uploading...")
//
// client.UploadSliceWithCallback("cosdemo",
//     "/hello/goasguen-cernvm-2015.pptx",
//     "/Users/goasguen-cernvm-2015.pptx",
//     "goasguen-cernvm-2015.pptx",
//     "",
//     1024*512,
//     func(res *cos.UploadSliceResponse, err error) {
//         if err != nil {
//             fmt.Println(err)
//             return
//         }
//
//         fmt.Println("Code:", res.Code,
//             "\nMessage:", res.Message,
//             "\nUrl:", res.Data.Url,
//             "\nResourcePath:", res.Data.ResourcePath,
//             "\nAccess Url:", res.Data.AccessUrl)
//
//         wg.Done()
//     })
//
// wg.Wait()
//
// fmt.Println("Uploaded...")
//
func (c *Client) UploadSliceWithCallback(bucket, dstPath, srcPath, bizAttr, session string, sliceSize int64, callback func(*UploadSliceResponse, error)) {
	go func() {
		response, err := c.UploadSlice(bucket, dstPath, srcPath, bizAttr, session, sliceSize)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的文件属性更新接口, 调用该函数可立即返回, 文件属性更新完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//     bizAttr: 目录属性, 由业务端维护
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.UpdateFileWithCallback("cosdemo", "/hello123/new.txt", "hello-new-attr",
//         func(res *cos.UpdateFileResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) UpdateFileWithCallback(bucket, path, bizAttr string, callback func(*UpdateFileResponse, error)) {
	go func() {
		response, err := c.UpdateFile(bucket, path, bizAttr)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的文件属性查询接口, 调用该函数可立即返回, 文件属性查询完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//     client.StatFileWithCallback("cosdemo", "/hello/new.txt",
//         func(res *cos.StatFileResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message,
//                 "\nName:", res.Data.Name,
//                 "\nBizAttr:", res.Data.BizAttr,
//                 "\nCtime:", res.Data.Ctime,
//                 "\nMtime:", res.Data.Mtime,
//                 "\nAccess Url:", res.Data.AccessUrl)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) StatFileWithCallback(bucket, path string, callback func(*StatFileResponse, error)) {
	go func() {
		response, err := c.StatFile(bucket, path)
		if callback != nil {
			callback(response, err)
		}
	}()
}

// 基于回调函数的文件删除接口, 调用该函数可立即返回, 文件删除完成后调用回调函数.
//     bucket:  Bucket 名称
//     path:    目录路径
//
// 示例:
//
//     client := cos.NewClient(appId, secretId, secretKey)
//     var wg = sync.WaitGroup{}
//
//     wg.Add(1)
//
//     client.DeleteFileWithCallback("cosdemo", "/hello123/new.txt",
//         func(res *cos.DeleteFileResponse, err error) {
//             if err != nil {
//                 fmt.Println(err)
//                 return
//             }
//
//             fmt.Println("Code:", res.Code,
//                 "\nMessage:", res.Message)
//
//             wg.Done()
//         })
//
//     wg.Wait()
//
func (c *Client) DeleteFileWithCallback(bucket, path string, callback func(*DeleteFileResponse, error)) {
	go func() {
		response, err := c.DeleteFile(bucket, path)
		if callback != nil {
			callback(response, err)
		}
	}()
}
