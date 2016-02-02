package main

import (
	"fmt"
	"time"

	"github.com/forhappy/cos-go-sdk"
)

func main() {
	appId := "10016247"
	secretId := "AKIDj0mWjQXxi3B65jCZS8BcWXYbGOKRuZPx"
	secretKey := "ytvcnVSIC22qs24HFRdS6beGAoJfEZmA"

	client := cos.NewClient(appId, secretId, secretKey)
	client.SetTimeout(5000 * time.Millisecond)

	res, err := client.UploadFile("cosdemo", "/hello/spark.pdf", "/Users/hpfu/spark.pdf", "new testcases")
	if err != nil {
		if client.IsTimeout(err) {
			fmt.Println("Request timeout.")
		} else {
			fmt.Println(err)
		}

		return
	}

	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nUrl:", res.Data.Url,
		"\nResourcePath:", res.Data.ResourcePath,
		"\nAccess Url:", res.Data.AccessUrl)
}
