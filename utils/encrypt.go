package utils

import (
	"encoding/base64"
)

func SimpDecrypt(token string) string {
	decoded, _ := base64.StdEncoding.DecodeString(token)
	return string(decoded)
}
