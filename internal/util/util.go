package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Password(pass *string) {
	hash := md5.Sum([]byte(*pass))
	*pass = hex.EncodeToString(hash[:])
}

func CheckFields(args ...interface{}) bool {
	for _, arg := range args {
		if arg == "" || arg == 0 {
			return false
		}
	}
	return true
}
