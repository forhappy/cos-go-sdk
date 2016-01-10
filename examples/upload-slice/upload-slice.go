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

	res, err := client.UploadSlice("cosdemo", "/hello/goasguen-cernvm-2015.pptx", "/Users/hpfu/goasguen-cernvm-2015.pptx", "goasguen-cernvm-2015.pptx", "", 1024*512)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nUrl:", res.Data.Url,
		"\nResourcePath:", res.Data.ResourcePath,
		"\nAccess Url:", res.Data.AccessUrl)
}
