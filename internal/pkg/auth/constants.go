package auth

import "time"

const (
	accessTokenLifeTime = time.Hour * 12
)

const (
	errMessageExpiredToken            = "Token has expired"
	errMessageUnexpectedSigningMethod = "Unexpected signing method"
)
