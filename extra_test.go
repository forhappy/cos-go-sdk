package cos

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestCreateAndDeleteFolderAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	client.SetTimeout(time.Second * 5)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestUpdateAndStatFolderAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	time.Sleep(time.Second)
	resUpdateChan := client.UpdateFolderAsync(BUCKET, folderName, "new-attr")
	resUpdateAsync := <-resUpdateChan
	err = resUpdateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpdate := resUpdateAsync.Response
	if resUpdate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpdate.Code, resUpdate.Message)
	}

	resStatChan := client.StatFolderAsync(BUCKET, folderName)
	resStatAsync := <-resStatChan
	err = resStatAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resStat := resStatAsync.Response
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resStat.Code, resStat.Message)
	}
	if resStat.Data.BizAttr != "new-attr" {
		t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-attr", resStat.Data.BizAttr)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestListFolderAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resListChan := client.ListFolderAsync(BUCKET, folderName, "", Both, 100, Asc)
	resListAsync := <-resListChan
	err = resListAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resList := resListAsync.Response
	if resList.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resList.Code, resList.Message)
	}

	resSearchChan := client.PrefixSearchAsync(BUCKET, folderName, "testing", "", Both, 100, Asc)
	resSearchAsync := <-resSearchChan
	err = resSearchAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resSearch := resSearchAsync.Response
	if resSearch.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resSearch.Code, resSearch.Message)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}

}

func TestUploadFileAndChunkAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resUploadChan := client.UploadFileAsync(BUCKET, folderName+"/smallfile.bin", "data/smallfile.bin", "Golang testcase for cos sdk UploadFile.")
	resUploadAsync := <-resUploadChan
	err = resUploadAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpload := resUploadAsync.Response
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resUploadChunkChan := client.UploadChunkAsync(BUCKET, folderName+"/smallchunk.bin", []byte("data/smallchunk.bin"), "Golang testcase for cos sdk UploadFile.")
	resUploadChunkAsync := <-resUploadChunkChan
	err = resUploadChunkAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpload = resUploadChunkAsync.Response
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resDeleteChunkChan := client.DeleteFileAsync(BUCKET, folderName+"/smallchunk.bin")
	resDeleteChunkAsync := <-resDeleteChunkChan
	err = resDeleteChunkAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDeleteChunk := resDeleteChunkAsync.Response
	if resDeleteChunk.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteChunk.Code, resDeleteChunk.Message)
	}

	resDeleteFileChan := client.DeleteFileAsync(BUCKET, folderName+"/smallfile.bin")
	resDeleteFileAsync := <-resDeleteFileChan
	err = resDeleteFileAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDeleteFile := resDeleteFileAsync.Response
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}

}

func TestUploadSliceAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	resUploadChan := client.UploadSliceAsync(BUCKET, folderName+"/bigfile.bin", "data/bigfile.bin", "Golang testcase for cos sdk UploadSlice.", "", 512*1024)
	resUploadAsync := <-resUploadChan
	err = resUploadAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpload := resUploadAsync.Response
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	resDeleteFileChan := client.DeleteFileAsync(BUCKET, folderName+"/bigfile.bin")
	resDeleteFileAsync := <-resDeleteFileChan
	err = resDeleteFileAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDeleteFile := resDeleteFileAsync.Response
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestUpdateAndStatFileAsync(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)

	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	resCreateChan := client.CreateFolderAsync(BUCKET, folderName, "attr")
	resCreateAsync := <-resCreateChan
	err := resCreateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resCreate := resCreateAsync.Response
	if resCreate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resCreate.Code, folderName, resCreate.Message)
	}

	filename := folderName + "/smallfile.bin"
	resUploadChan := client.UploadFileAsync(BUCKET, filename, "data/smallfile.bin", "Golang testcase for cos sdk UploadFile.")
	resUploadAsync := <-resUploadChan
	err = resUploadAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpload := resUploadAsync.Response
	if resUpload.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpload.Code, resUpload.Message)
	}

	time.Sleep(time.Second)
	resUpdateChan := client.UpdateFileAsync(BUCKET, filename, "new-file-attr")
	resUpdateAsync := <-resUpdateChan
	err = resUpdateAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resUpdate := resUpdateAsync.Response
	if resUpdate.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resUpdate.Code, resUpdate.Message)
	}

	resStatChan := client.StatFileAsync(BUCKET, filename)
	resStatAsync := <-resStatChan
	err = resStatAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resStat := resStatAsync.Response
	if resStat.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, resStat.Code, filename, resStat.Message)
	}
	if resStat.Data.BizAttr != "new-file-attr" {
		t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-file-attr", resStat.Data.BizAttr)
	}

	resDeleteFileChan := client.DeleteFileAsync(BUCKET, filename)
	resDeleteFileAsync := <-resDeleteFileChan
	err = resDeleteFileAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDeleteFile := resDeleteFileAsync.Response
	if resDeleteFile.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDeleteFile.Code, resDeleteFile.Message)
	}

	resDeleteChan := client.DeleteFolderAsync(BUCKET, folderName)
	resDeleteAsync := <-resDeleteChan
	err = resDeleteAsync.Error
	if err != nil {
		t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
	}
	resDelete := resDeleteAsync.Response
	if resDelete.Code != 0 {
		t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, resDelete.Code, resDelete.Message)
	}
}

func TestCreateAndDeleteFolderWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}

func TestUpdateAndStatFolderWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	time.Sleep(time.Second)
	client.UpdateFolderWithCallback(BUCKET, folderName, "new-attr", func(res *UpdateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.StatFolderWithCallback(BUCKET, folderName, func(res *StatFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		if res.Data.BizAttr != "new-attr" {
			t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-attr", res.Data.BizAttr)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}

func TestListFolderWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.ListFolderWithCallback(BUCKET, folderName, "", Both, 100, Asc, func(res *ListFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.PrefixSearchWithCallback(BUCKET, folderName, "testing", "", Both, 100, Asc, func(res *ListFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}

func TestUploadChunkWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	filename := folderName + "/chunkfile.bin"
	wg.Add(1)
	client.UploadChunkWithCallback(BUCKET, filename, []byte("data/chunkfile.bin"), "Golang testcase for cos sdk UploadChunk.", func(res *UploadFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFileWithCallback(BUCKET, filename, func(res *DeleteFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}

func TestUploadSliceWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	filename := folderName + "/bigfile.bin"
	wg.Add(1)
	client.UploadSliceWithCallback(BUCKET, filename, "data/bigfile.bin", "Golang testcase for cos sdk UploadSlice.", "", 512*1024, func(res *UploadSliceResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFileWithCallback(BUCKET, filename, func(res *DeleteFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}

func TestUpdateAndStatFileWithCallback(t *testing.T) {
	client := NewClient(APPID, SECRETID, SECRETKEY)
	var wg = sync.WaitGroup{}

	wg.Add(1)
	folderName := "/testing" + strconv.Itoa(rand.Intn(1000000000))
	client.CreateFolderWithCallback(BUCKET, folderName, "attr", func(res *CreateFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s:%s]", err, res.Message)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s:%s]", 0, res.Code, folderName, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	filename := folderName + "/smallfile.bin"
	wg.Add(1)
	client.UploadFileWithCallback(BUCKET, filename, "data/smallfile.bin", "Golang testcase for cos sdk UploadFile.", func(res *UploadFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	time.Sleep(time.Second)
	client.UpdateFileWithCallback(BUCKET, filename, "new-attr", func(res *UpdateFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.StatFileWithCallback(BUCKET, filename, func(res *StatFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		if res.Data.BizAttr != "new-attr" {
			t.Errorf("Return bizAttr should match [EXPECTED:%s]:[ACTUAL:%s]", "new-attr", res.Data.BizAttr)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFileWithCallback(BUCKET, filename, func(res *DeleteFileResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()

	wg.Add(1)
	client.DeleteFolderWithCallback(BUCKET, folderName, func(res *DeleteFolderResponse, err error) {
		if err != nil {
			t.Errorf("Error should match [EXPECTED:nil]:[ACTUAL:%s]", err)
		}
		if res.Code != 0 {
			t.Errorf("Return code should match [EXPECTED:%d]:[ACTUAL:%d:%s]", 0, res.Code, res.Message)
		}
		wg.Done()
	})
	wg.Wait()
}
