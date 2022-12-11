package jwt

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
)

var ErrInvalidAccessToken = errors.New("invalid access token")

func ParseToken(accessToken string, key []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", err
	}

	// parse claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", ErrInvalidAccessToken
}
