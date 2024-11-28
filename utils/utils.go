package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

// The key should be stored in an environment variable or configuration file
var secretKey = []byte("your-secret-key")

// Parse a JWT token and return its claims (including user_id).
func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Parse the claims from the token.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
