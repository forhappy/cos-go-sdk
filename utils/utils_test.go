package utils

import (
	"testing"
)

func TestUrlEncode(t *testing.T) {
	tests := map[string]string{
		"/cos/hello world":     "/cos/hello%20world",
		"/cos/hello=world":     "/cos/hello%3Dworld",
		"/appid/bucketname/测试": "/appid/bucketname/%E6%B5%8B%E8%AF%95",
	}

	for str, expected := range tests {
		actual := UrlEncode(str)
		if expected != actual {
			t.Errorf("Should match [EXPECTED:%s]:[ACTUAL:%s]", expected, actual)
		}
	}
}
