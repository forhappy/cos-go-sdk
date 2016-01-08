package cos

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const APPID = "10016247"
const SECRETID = "AKIDj0mWjQXxi3B65jCZS8BcWXYbGOKRuZPx"
const SECRETKEY = "ytvcnVSIC22qs24HFRdS6beGAoJfEZmA"
const BUCKET = "cosdemo"

func init() {
	rand.Seed(time.Now().Unix())
}

func TestCreateAndDeleteFolder(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	client.SetTimeout(time.Second * 5)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resDelete, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resCreate.Message)
	}
}

func TestUpdateAndStatFolder(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	time.Sleep(time.Second)
	resUpdate, err := client.UpdateFolder(BUCKET, folderName, "new-attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpdate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpdate.Code, resUpdate.Message)
	}

	resStat, err := client.StatFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resStat.Code, resStat.Message)
	}
	if resStat.Data.BizAttr != "new-attr" {
		t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-attr", resStat.Data.BizAttr)
	}

	resStat, err = client.StatFolder(BUCKET, "")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resStat.Code, resStat.Message)
	}

	resDelete, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestListFolder(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resList, err := client.ListFolder(BUCKET, folderName, "", Both, 100, Asc)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resList.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resList.Code, resList.Message)
	}

	resList, err = client.ListFolder(BUCKET, folderName, "", DirectoryOnly, 100, Asc)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resList.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resList.Code, resList.Message)
	}

	resList, err = client.ListFolder(BUCKET, folderName, "", FileOnly, 100, Desc)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resList.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resList.Code, resList.Message)
	}

	resSearch, err := client.PrefixSearch(BUCKET, folderName, "testing", "", Both, 100, Asc)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resSearch.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resSearch.Code, resSearch.Message)
	}

	resDelete, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestUploadFile(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resUpload, err := client.UploadFile(BUCKET, folderName+"/smallfile.bin", "data/smallfile.bin", "Golang testcase for cos sdk UploadFile.")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resUpload, err = client.UploadFile(BUCKET, folderName+"/smallfile-.bin", "data/nosuchfile.bin", "Golang testcase for cos sdk UploadFile.")
	if err == nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", "No such file.", err)
	}

	resDeleteFile, err := client.DeleteFile(BUCKET, folderName+"/smallfile.bin")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDeleteFolder, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFolder.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFolder.Code, resDeleteFolder.Message)
	}
}

func TestPrepareToUploadSlice(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resUpload, err := client.PrepareToUploadSlice(BUCKET, folderName+"/bigfile.bin", "data/nosuchbigfile.bin", "Golang testcase for cos sdk UploadSlice.", "aaabbbccc", 512*1024)
	if err == nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", "No such file", err)
	}
	if resUpload != nil {
		t.Errorf("Return should match [EXPECTED:nil]:[ACTUAL:%v]", 0, resUpload)
	}

	resUpload, err = client.PrepareToUploadSlice(BUCKET, folderName+"/nosuchfolder/bigfile.bin", "data/x/nosuchbigfile.bin", "Golang testcase for cos sdk UploadSlice.", "aaabbbccc", 512*1024)
	if err == nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", "No such file", err)
	}
	if resUpload != nil {
		t.Errorf("Return should match [EXPECTED:nil]:[ACTUAL:%v]", 0, resUpload)
	}

	resUpload, err = client.PrepareToUploadSlice(BUCKET, folderName+"/nosuchfolder/bigfile.bin", "data/nosuchbigfile.bin", "Golang testcase for cos sdk UploadSlice.", "aaabbbccc", 1024)
	if err == nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", "No such file", err)
	}
	if resUpload != nil {
		t.Errorf("Return should match [EXPECTED:nil]:[ACTUAL:%v]", 0, resUpload)
	}

	resDeleteFolder, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFolder.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFolder.Code, resDeleteFolder.Message)
	}
}
func TestUploadSlice(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resUpload, err := client.UploadSlice(BUCKET, folderName+"/bigfile.bin", "data/nosuchbigfile.bin", "Golang testcase for cos sdk UploadSlice.", "", 512*1024)
	if err == nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", "No such file", err)
	}

	resUpload, err = client.UploadSlice(BUCKET, folderName+"/bigfile.bin", "data/bigfile.bin", "Golang testcase for cos sdk UploadSlice.", "", 512*1024)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resUpload, err = client.UploadSlice(BUCKET, folderName+"/bigfile.bin", "data/bigfile.bin", "Golang testcase for cos sdk UploadSlice.", "", 512*1024)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resDeleteFile, err := client.DeleteFile(BUCKET, folderName+"/bigfile.bin")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDeleteFolder, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFolder.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFolder.Code, resDeleteFolder.Message)
	}
}

func TestUpdateAndStatFile(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreate, err := client.CreateFolder(BUCKET, folderName, "attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s:%s]", nil, err, resCreate.Message)
	}
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	filename := folderName + "/smallfile.bin"
	resUpload, err := client.UploadFile(BUCKET, filename, "data/smallfile.bin", "Golang testcase for cos sdk UploadFile.")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	time.Sleep(time.Second)
	resUpdate, err := client.UpdateFile(BUCKET, filename, "new-file-attr")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resUpdate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpdate.Code, resUpdate.Message)
	}

	resStat, err := client.StatFile(BUCKET, filename)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resStat.Code, filename, resStat.Message)
	}
	if resStat.Data.BizAttr != "new-file-attr" {
		t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-file-attr", resStat.Data.BizAttr)
	}

	resStat, err = client.StatFile(BUCKET, "")
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resStat.Code, resStat.Message)
	}

	resDeleteFile, err := client.DeleteFile(BUCKET, filename)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDelete, err := client.DeleteFolder(BUCKET, folderName)
	if err != nil {
		t.Errorf("Error should match [EXPECTED:%s]:[ACTUAL:%s]", nil, err)
	}
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}
