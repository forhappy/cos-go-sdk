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

	resChan := client.PrefixSearchAsync("cosdemo", "/hello", "g", "", cos.Both, 100, cos.Asc)

	resAsync := <-resChan

	if resAsync.Error != nil {
		fmt.Println(resAsync.Error)
		return
	}

	res := resAsync.Response
	fmt.Println("Code:", res.Code,
		"\nMessage:", res.Message,
		"\nContext:", res.Data.Context,
		"\nHasMore:", res.Data.HasMore,
		"\nDirCount:", res.Data.DirCount,
		"\nFileCount:", res.Data.FileCount,
	)

	fmt.Println("*************************************")
	for _, info := range res.Data.Infos {
		fmt.Println("Name:", info.Name,
			"\nBizAttr:", info.BizAttr,
			"\nFileSize:", info.FileSize,
			"\nFileLen:", info.FileLen,
			"\nSha:", info.Sha,
			"\nCtime:", info.Ctime,
			"\nMtime:", info.Mtime,
			"\nAccess URL:", info.AccessUrl,
		)
		fmt.Println("*************************************")
	}

}
