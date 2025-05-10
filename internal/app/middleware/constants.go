package middleware

const (
	authCookieName = "Authorization"
	userIDKey      = contextKey("userID")
)

func UserIDKey() contextKey {
	return userIDKey
}
