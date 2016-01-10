package main

import (
	"fmt"
	"github.com/forhappy/cos-go-sdk"
)

func main() {
	appId := "10016247"
	secretId := "AKIDj0mWjQXxi3B65jCZS8BcWXYbGOKRuZPx"
	secretKey := "ytvcnVSIC22qs24HFRdS6beGAoJfEZmA"

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
