package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

type Signer interface {
	Sign(secretKey string) (string, error)
	SignOnce(secretKey string) (string, error)
}

type Signature struct {
	AppId       string
	Bucket      string
	SecretId    string
	ExpiredTime string
	CurrentTime string
	Rand        string
	FileId      string
}

func NewSignature(appId, bucket, secretId, expiredTime, currentTime, rand, fileId string) *Signature {
	return &Signature{
		AppId:       appId,
		Bucket:      bucket,
		SecretId:    secretId,
		ExpiredTime: expiredTime,
		CurrentTime: currentTime,
		Rand:        rand,
		FileId:      fileId,
	}
}

func (s *Signature) Sign(secretKey string) string {
	stringToSign := fmt.Sprintf("a=%s&k=%s&e=%s&t=%s&r=%s&f=%s&b=%s",
		s.AppId,
		s.SecretId,
		s.ExpiredTime,
		s.CurrentTime,
		s.Rand,
		"",
		s.Bucket,
	)

	hmacSha1 := hmac.New(sha1.New, []byte(secretKey))
	hmacSha1.Write([]byte(stringToSign))
	bytesSign := hmacSha1.Sum(nil)
	bytesSign = append(bytesSign, []byte(stringToSign)...)
	signature := base64.StdEncoding.EncodeToString(bytesSign)
	return signature
}

func (s *Signature) SignOnce(secretKey string) string {
	stringToSign := fmt.Sprintf("a=%s&k=%s&e=%s&t=%s&r=%s&f=%s&b=%s",
		s.AppId,
		s.SecretId,
		"0",
		s.CurrentTime,
		s.Rand,
		s.FileId,
		s.Bucket,
	)

	hmacSha1 := hmac.New(sha1.New, []byte(secretKey))
	hmacSha1.Write([]byte(stringToSign))
	bytesSign := hmacSha1.Sum(nil)
	bytesSign = append(bytesSign, []byte(stringToSign)...)
	signature := base64.StdEncoding.EncodeToString(bytesSign)
	return signature
}

func (s *Signature) String() string {
	str := fmt.Sprintf("a=%s&b=%s&k=%s&e=%s&t=%s&r=%s&f=%s",
		s.AppId,
		s.SecretId,
		s.ExpiredTime,
		s.CurrentTime,
		s.Rand,
		s.FileId,
		s.Bucket,
	)

	return str
}
