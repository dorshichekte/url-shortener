package utils

import (
	"math/rand"
	"time"
)

const (
	defaultRandomStringLength = 6
	charset                   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func CreateRandomString() string {
	scr := rand.NewSource(time.Now().UnixNano())
	r := rand.New(scr)

	var result []byte
	for i := 0; i < defaultRandomStringLength; i++ {
		result = append(result, charset[r.Intn(len(charset))])
	}

	return string(result)
}
