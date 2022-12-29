package service

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Shorten(originalURL string) string {
	hash := md5.Sum([]byte(strings.ToLower(originalURL)))
	return hex.EncodeToString(hash[:])
}
