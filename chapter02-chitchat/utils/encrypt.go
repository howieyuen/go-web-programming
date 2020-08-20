package utils

import (
	"crypto/sha1"
	"fmt"
)

func Encrypt(plainText string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
}
