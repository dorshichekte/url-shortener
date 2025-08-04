// Пакет auth реализует JWT-аутентификацию.
package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// New создает новый экземпляр сервиса аутентификации.
func New(accessSecret string) Auth {
	return &auth{accessSecret: accessSecret}
}

// Generate создает новую пару JWT-токенов для пользователя.
func (auth *auth) Generate(userID string) (TokenPair, error) {
	claims := newClaims(userID)
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(auth.accessSecret))
	if err != nil {
		return TokenPair{}, err
	}

	authData := TokenPair{
		AccessToken: accessToken,
	}

	return authData, nil
}

// ParseAccessToken парсит и валидирует access-токен.
func (auth *auth) ParseAccessToken(accessToken string) (UserData, error) {
	parseAccessTokenFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUnexpectedSigningMethod
		}

		return []byte(auth.accessSecret), nil
	}

	claims := &Claims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, parseAccessTokenFunc)
	if err != nil {
		return UserData{}, err
	}

	isValidErr := claims.valid()
	if isValidErr != nil {
		return UserData{}, isValidErr
	}

	return claims.UserData, nil
}
