package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}
