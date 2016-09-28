package tools

import (
	"fmt"
	"crypto/sha256"
)

func Crypt_sha256(to_hash string) string {
	
	s := sha256.Sum256([]byte(to_hash))
	
	return fmt.Sprintf("%x", s)
}