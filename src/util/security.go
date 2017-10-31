package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	res := fmt.Sprintf("%x", has)
	return res
}
