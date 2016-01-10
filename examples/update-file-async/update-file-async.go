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

	resChan := client.UpdateFileAsync("cosdemo", "/hello/hello.txt", "hello-attr")
	resAsync := <-resChan

	if resAsync.Error != nil {
		fmt.Println(resAsync.Error)
		return
	}

	res := resAsync.Response

	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
}
