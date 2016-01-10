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

	res, err := client.UpdateFolder("cosdemo", "/hello/", "hello-attr")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Code:", res.Code, "\nMessage:", res.Message)
}
