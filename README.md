[![Build Status](https://drone.io/github.com/forhappy/cos-go-sdk/status.png)](https://drone.io/github.com/forhappy/cos-go-sdk/latest)
[![Coverage Status](https://coveralls.io/repos/forhappy/cos-go-sdk/badge.svg?branch=master&service=github)](https://coveralls.io/github/forhappy/cos-go-sdk?branch=master)
[![GoDoc](https://godoc.org/github.com/forhappy/cos-go-sdk?status.png)](https://godoc.org/github.com/forhappy/cos-go-sdk)

##COS-Go-SDK 简介
- 腾讯云对象存储服务（Cloud Object Service）是基于腾讯多年海量服务经验，对外提供的可靠、安全、易用的海量存储服务。
- 腾讯云对象存储服务提供多样化接入方式，稳定安全，无线扩容以及遍布全国的加速节点为您提供高质量的上传与下载。
- COS Go SDK 基于[腾讯云对象存储服务 COS](http://www.qcloud.com/product/cos.html) 官方 Restful API 构建，完全兼容腾讯云对象存储服务接口, 用户可使用 COS-Go-SDK 实现数据的上传下载功能。

##环境
- COS-Go-SDK 推荐使用 Go 1.2 及以上 Go 语言版本。
- Windows，Linux，Mac OS X

##安装
```bash
go get github.com/forhappy/cos-go-sdk
```

##快速入门

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
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message,
            "\nCtime:", res.Data.Ctime,
            "\nResource Path:", res.Data.ResourcePath)
        wg.Done()
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
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message,
            "\nUrl:", res.Data.Url,
            "\nResourcePath:", res.Data.ResourcePath,
            "\nAccess Url:", res.Data.AccessUrl)
        wg.Done()
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
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Code:", res.Code,
            "\nMessage:", res.Message)
        wg.Done()
    })
wg.Wait()
```
