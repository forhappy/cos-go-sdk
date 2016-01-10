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

	client.DeleteFileWithCallback("cosdemo", "/hello/", func(res *cos.DeleteFileResponse, err error) {
		defer wg.Done()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
	})

	wg.Wait()
}
