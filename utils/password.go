package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordEncrypt(rawPassword string) string {
	s := sha256.New()
	s.Write([]byte(rawPassword))
	return hex.EncodeToString(s.Sum(nil))
}

func PasswordVerify(rawPassword string, encryptedPassword string) bool {
	return PasswordEncrypt(rawPassword) == encryptedPassword
}
