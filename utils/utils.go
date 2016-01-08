package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
)

const FileChunkSize = 8 * 1024

func isAlphaNum(c byte) bool {
	if !('a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9') {
		return false
	}

	return true
}

func toHex(i byte) byte {
	charTable := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F'}
	return charTable[i]
}

func UrlEncode(s string) string {
	encoded := ""
	for i := 0; i < len(s); i += 1 {
		char := s[i]
		if isAlphaNum(byte(char)) ||
			char == '-' ||
			char == '_' ||
			char == '.' ||
			char == '~' ||
			char == '/' {
			encoded += string(char)
		} else {
			encoded += string('%')
			encoded += string(toHex(byte(char >> 4)))
			encoded += string(toHex(byte(char % 16)))
		}
	}

	return encoded
}

func GetFileSize(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}

	return fileInfo.Size(), nil
}

func HashFileWithSha1(path string) (string, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", -1, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", -1, err
	}

	sha1 := sha1.New()
	fileSize := fileInfo.Size()
	blocks := uint64(math.Ceil(float64(fileSize) / float64(FileChunkSize)))
	for i := uint64(0); i < blocks; i += 1 {
		blockSize := int(math.Min(FileChunkSize, float64(fileSize-int64(i*FileChunkSize))))
		buf := make([]byte, blockSize)
		file.Read(buf)
		io.WriteString(sha1, string(buf))
	}

	return hex.EncodeToString(sha1.Sum(nil)), fileSize, nil
}

func HashBufferWithSha1(buf []byte) (string, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, string(buf))
	return hex.EncodeToString(sha1.Sum(nil)), nil
}

func StructToMap(s interface{}) (map[string]string, error) {
	m := make(map[string]string)

	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return m, fmt.Errorf("Only accepts structs, got %T", val)
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		var v string
		tag := typ.Field(i).Tag.Get("json")

		f := val.Field(i)
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', -1, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', -1, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		if v != "" {
			m[tag] = v
		}
	}

	return m, nil
}
