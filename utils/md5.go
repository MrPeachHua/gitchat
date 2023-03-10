package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	temStr := h.Sum(nil)
	return hex.EncodeToString(temStr)
}
