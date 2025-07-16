package auth

import "time"

const (
	accessTokenLifeTime = time.Hour * 12
)

const (
	expiredToken            = "Token has expired"
	invalidToken            = "Token is invalid"
	initializationToken     = "Failed initialize token"
	emptyUserID             = "Empty user id"
	unexpectedSigningMethod = "Unexpected signing method"
)
