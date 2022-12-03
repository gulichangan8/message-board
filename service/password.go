package service

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// IsPasswordErr 检查密码字符是否违规
func IsPasswordErr(password string) bool {
	if strings.Contains(password, "\\") == true {
		return false
	} else {
		return true
	}
}

// IsShortPassword 检查密码长度
func IsShortPassword(password string) bool {
	var name = []rune(password)
	if len(name) > 6 {
		return true
	} else {
		return false
	}
}

func Hmac(key, date string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(date))
	return hex.EncodeToString(hash.Sum(nil))
}
