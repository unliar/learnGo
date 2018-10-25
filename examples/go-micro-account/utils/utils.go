package utils

import (
	"crypto/md5"
	"fmt"
)

// CreateMD5 是用来创建md5值的
func CreateMD5(s, salt string) string {
	result := s + salt
	return fmt.Sprintf("%x", md5.Sum([]byte(result)))
}
