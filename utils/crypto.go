package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

/**
加密字符串到密码
*/
func EncryptStringToPassword(plainText string) (encrypted string, err error) {
	h := sha1.New()
	_, err = io.WriteString(h, plainText)
	hSum := h.Sum(nil)
	encrypted = hex.EncodeToString(hSum)
	return
}
