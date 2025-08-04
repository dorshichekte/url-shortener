package string

import (
	"crypto/rand"
	"encoding/hex"
)

// CreateRandom создает случайную строку.
func CreateRandom() string {
	buf := make([]byte, 4)
	if _, err := rand.Read(buf); err != nil {
		panic("failed to generate secure random string: " + err.Error())
	}
	return hex.EncodeToString(buf)
}
