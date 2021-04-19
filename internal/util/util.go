package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Password(pass *string) {
	hash := md5.Sum([]byte(*pass))
	*pass = hex.EncodeToString(hash[:])
}

func AllTrue(args ...interface{}) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}
	return true
}
