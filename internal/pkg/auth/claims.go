package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func newClaims(userID int) *Claims {
	lifeTime := time.Now().Add(accessTokenLifeTime)

	claims := Claims{
		UserData: UserData{ID: userID},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(lifeTime),
		},
	}

	return &claims
}

func (c *Claims) valid() error {
	isTimeExpired := time.Now().After(c.ExpiresAt.Time)
	if isTimeExpired {
		return errExpiredToken
	}

	return nil
}
