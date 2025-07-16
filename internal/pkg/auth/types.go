package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	ID int
}

type Claims struct {
	jwt.RegisteredClaims
	UserData UserData
}

type TokenPair struct {
	AccessToken string
}

type auth struct {
	accessSecret string
}

type Auth interface {
	Generate(userID int) (TokenPair, error)
	ParseAccessToken(accessToken string) (UserData, error)
}
