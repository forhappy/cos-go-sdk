[![Build Status](https://drone.io/github.com/forhappy/cos-go-sdk/status.png)](https://drone.io/github.com/forhappy/cos-go-sdk/latest)
[![Coverage Status](https://coveralls.io/repos/forhappy/cos-go-sdk/badge.svg?branch=master&service=github)](https://coveralls.io/github/forhappy/cos-go-sdk?branch=master)
[![GoDoc](https://godoc.org/github.com/forhappy/cos-go-sdk?status.svg)](https://godoc.org/github.com/forhappy/cos-go-sdk)
##COS-Go-SDK 简介
- 腾讯云对象存储服务（Cloud Object Service）是基于腾讯多年海量服务经验，对外提供的可靠、安全、易用的海量存储服务。
- 腾讯云对象存储服务提供多样化接入方式，稳定安全，无线扩容以及遍布全国的加速节点为您提供高质量的上传与下载。
- COS Go SDK 基于[腾讯云对象存储服务 COS](http://www.qcloud.com/product/cos.html) 官方 Restful API 构建，完全兼容腾讯云对象存储服务接口, 用户可使用 COS-Go-SDK 实现数据的上传下载功能。
- COS Go SDK 提供了三套接口，1. 普通接口，2. 异步接口，所谓异步接口，即调用此类接口时调用方会立即返回一管道，此后调用方可以读取改管道获取返回接口，3. 基于回调函数的接口，即调用此类接口时调用方需提供一个回调函数，调用时此类接口也会立即返回，当任务完成后会自动调用回调函数。

##环境
- COS-Go-SDK 推荐使用 Go 1.2 及以上 Go 语言版本。
- Windows，Linux，Mac OS X

##安装
```bash
go get github.com/forhappy/cos-go-sdk
```

##API 文档
[![GoDoc](https://godoc.org/github.com/forhappy/cos-go-sdk?status.svg)](https://godoc.org/github.com/forhappy/cos-go-sdk)

##快速入门

###文件查询完整示例
```go
package main
import (
	"fmt"
	"github.com/forhappy/cos-go-sdk"
)
func main() {
	appId := "YOUR-APP-ID"
	secretId := "YOUR-SECRET-ID"
	secretKey := "YOUR-SECRET-KEY"
	client := cos.NewClient(appId, secretId, secretKey)
	res, err := client.StatFile("cosdemo", "/hello/hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nName:", res.Data.Name,
		"\nBizAttr:", res.Data.BizAttr,
		"\nFileSize:", res.Data.FileSize,
		"\nFileLen:", res.Data.FileLen,
		"\nSha:", res.Data.Sha,
		"\nCtime:", res.Data.Ctime,
		"\nMtime:", res.Data.Mtime,
		"\nAccess Url:", res.Data.AccessUrl)
}
```

##更多示例

### 创建目录
- 普通接口
```go
client := cos.NewClient(appId, secretId, secretKey)
res, err := client.CreateFolder("cosdemo", "/hello", "hello",)
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message,
    "\nCtime:", res.Data.Ctime,
    "\nResource Path:", res.Data.ResourcePath)
```

- 异步接口
```go
client := cos.NewClient(appId, secretId, secretKey)
resChan := client.CreateFolderAsync("cosdemo", "/hello", "hello",)
// Do your other work here
resAsync := <- resChan
if resAsync.Error != nil {
    fmt.Println(resAsync.Error)
    return
}
res := resAsync.Response
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message,
    "\nCtime:", res.Data.Ctime,
    "\nResource Path:", res.Data.ResourcePath)
```

- 回调接口
```go
client := cos.NewClient(appId, secretId, secretKey)
var wg = sync.WaitGroup{}
wg.Add(1)
client.CreateFolderWithCallback("cosdemo", "/hello123", "hello",
    func(res *cos.CreateFolderResponse, err error) {
        defer wg.Done()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message,
            "\nCtime:", res.Data.Ctime,
            "\nResource Path:", res.Data.ResourcePath)
    })
wg.Wait()
```

###文件上传

- 普通接口
```go
client := cos.NewClient(appId, secretId, secretKey)
res, err := client.UploadFile("cosdemo", "/hello/hello.txt", "/users/new.txt", "file attr")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message,
    "\nUrl:", res.Data.Url,
    "\nResourcePath:", res.Data.ResourcePath,
    "\nAccess Url:", res.Data.AccessUrl)
```

- 异步接口
```go
client := cos.NewClient(appId, secretId, secretKey)
resChan := client.UploadFileAsync("cosdemo", "/hello/hello.txt", "/users/new.txt", "file attr")
// Do your other work here
resAsync := <- resChan
if resAsync.Error != nil {
    fmt.Println(resAsync.Error)
    return
}
res := resAsync.Response
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message,
    "\nUrl:", res.Data.Url,
    "\nResourcePath:", res.Data.ResourcePath,
    "\nAccess Url:", res.Data.AccessUrl)
```

- 回调接口
```go
client := cos.NewClient(appId, secretId, secretKey)
var wg = sync.WaitGroup{}
wg.Add(1)
fmt.Println("Uploading...")
client.UploadFileWithCallback("cosdemo",
    "/hello/goasguen-cernvm-2015.pptx",
    "/Users/goasguen-cernvm-2015.pptx",
    "goasguen-cernvm-2015.pptx",
    func(res *cos.UploadFileResponse, err error) {
        defer wg.Done()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message,
            "\nUrl:", res.Data.Url,
            "\nResourcePath:", res.Data.ResourcePath,
            "\nAccess Url:", res.Data.AccessUrl)
    })
wg.Wait()
fmt.Println("Uploaded...")
```

###删除文件

- 普通接口
```go
client := cos.NewClient(appId, secretId, secretKey)
res, err := client.DeleteFile("cosdemo", "/hello/hello.txt")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message)
```

- 异步接口
```go
client := cos.NewClient(appId, secretId, secretKey)
resChan := client.DeleteFileAsync("cosdemo", "/hello/hello.txt")
// Do your other work here
resAsync := <- resChan
if resAsync.Error != nil {
    fmt.Println(resAsync.Error)
    return
}
res := resAsync.Response
fmt.Println("Code:", res.Code,
    "\nMessage:", res.Message)
```

- 回调接口
```go
client := cos.NewClient(appId, secretId, secretKey)
var wg = sync.WaitGroup{}
wg.Add(1)
client.DeleteFileWithCallback("cosdemo", "/hello123/hello.txt",
    func(res *cos.DeleteFileResponse, err error) {
        defer wg.Done()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message)
    })
wg.Wait()
```

##完整示例

更多示例请查看 examples 目录

##项目文档

更多文档请查看 docs 目录
