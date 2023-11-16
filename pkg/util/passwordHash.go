package util

import (
	"crypto/md5"
	"fmt"
)

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
}

func VerifyPassword(requestPassword, dbPassword string) bool {
	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword
}
