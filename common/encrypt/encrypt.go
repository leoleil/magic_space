package encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

// MD5加密
func GetMd5Key(key string) string {
	data := []byte(key)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// base64加密
func GetBase64Key(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(key))
}

// base64解密
func DecodeBase64Key(key string) string {
	decoded, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return string(decoded)
	}
	return ""
}
