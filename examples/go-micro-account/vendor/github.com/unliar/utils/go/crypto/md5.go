package crypto

import "fmt"
import "crypto/md5"

// GetMD5 return a md5 string of your input
func GetMD5(s ...string) string {
	var d string
	for _, item := range s {
		d = d + item
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(d)))
}
