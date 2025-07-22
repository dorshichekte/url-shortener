package string

import (
	"encoding/hex"
	"math/rand"
)

var hashBuffer = make([]byte, 4)

func CreateRandom() string {
	rand.Read(hashBuffer)
	return hex.EncodeToString(hashBuffer)
}
