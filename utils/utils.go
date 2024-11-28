package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint) (string, error) {
	// Set token expiration time
	// Set token expiration time
	accessTokenExpireMinutes := os.Getenv("ACCESS_TOKEN_EXPIRE_MINUTES")
	var num int
	num, _ = strconv.Atoi(accessTokenExpireMinutes)

	// Token validity period is ACCESS_TOKEN_EXPIRE_MINUTES minutes
	expirationTime := time.Now().Add(time.Duration(num) * time.Minute)

	// Creating claims (payload)
	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	// Use the key to sign and generate a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the signature method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Parsing the claims in the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
