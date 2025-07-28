package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("secret_key")

func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"UserID": userID,
		"role":   role,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}

func ParseJWT(tokenstring string) (uint, string, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil || !token.Valid {
		return 0, "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("userID missing in the token")
	}
	userIDFlaot, ok := claims["UserID"].(float64)
	if !ok {
		return 0, "", errors.New("userID missing in token")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return 0, "", errors.New("role missing in the token")
	}
	return uint(userIDFlaot), role, nil
}
