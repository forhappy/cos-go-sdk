package main

import (
	"fmt"
	"sync"

	"github.com/forhappy/cos-go-sdk"
)

func main() {
	appId := "10016247"
	secretId := "AKIDj0mWjQXxi3B65jCZS8BcWXYbGOKRuZPx"
	secretKey := "ytvcnVSIC22qs24HFRdS6beGAoJfEZmA"

	client := cos.NewClient(appId, secretId, secretKey)
	var wg = sync.WaitGroup{}

	wg.Add(1)

	client.StatFileWithCallback("cosdemo", "/hello/hello.txt", func(res *cos.StatFileResponse, err error) {
		defer wg.Done()

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

	})

	wg.Wait()
}
