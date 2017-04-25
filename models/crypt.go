package models

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Md5(buf []byte) string {
	has := md5.New()
	has.Write(buf)
	return fmt.Sprintf("%x", has.Sum(nil))
}
func Sha1(buf []byte) string {
	has := sha1.New()
	has.Write(buf)
	return fmt.Sprintf("%x", has.Sum(nil))
}
