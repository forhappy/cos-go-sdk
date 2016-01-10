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

	fmt.Println("Uploading...")

	client.UploadSliceWithCallback("cosdemo",
		"/hello/goasguen-cernvm-2015.pptx",
		"/Users/hpfu/goasguen-cernvm-2015.pptx",
		"goasguen-cernvm-2015.pptx",
		"",
		1024*512,
		func(res *cos.UploadSliceResponse, err error) {
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
}
