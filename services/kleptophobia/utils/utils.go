package utils

import "crypto/md5"

func GetHash(s string) []byte {
	res := md5.Sum([]byte(s))
	return res[:]
}
