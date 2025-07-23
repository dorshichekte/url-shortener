package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// UserData содержит информацию о пользователе.
type UserData struct {
	ID string
}

// Claims структура claims для JWT токена.
type Claims struct {
	jwt.RegisteredClaims
	UserData UserData
}

// TokenPair содержит пару токенов доступа.
type TokenPair struct {
	AccessToken string
}

type auth struct {
	accessSecret string
}

// Auth интерфейс сервиса аутентификации.
type Auth interface {
	Generate(userID string) (TokenPair, error)
	ParseAccessToken(accessToken string) (UserData, error)
}
