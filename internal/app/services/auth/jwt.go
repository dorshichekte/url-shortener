package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"url-shortener/internal/app/constants"
	utilsString "url-shortener/internal/app/utils/string"
)

func (c *Claims) Valid() error {
	if time.Now().After(c.ExpiresAt.Time) {
		return constants.ErrTokenHasExpired
	}

	if c.UserID == "" {
		return constants.ErrEmptyUserID
	}

	return nil
}

func (c *Claims) Parse(accessToken string) error {
	token, err := jwt.ParseWithClaims(accessToken, c, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constants.ErrUnexpectedSigningMethod
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return constants.ErrTokenNotValid
	}

	return nil
}

func (c *Claims) CreateJWTToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func NewClaims() *Claims {
	userID := utilsString.CreateRandom()
	lifeTime := time.Now().Add(time.Hour * 8)

	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(lifeTime),
		},
		UserID: userID,
	}

	return claims
}
