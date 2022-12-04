package ping

import (
	"net/http"
	"strings"

	"github.com/ilfey/go-back/internal/app/endpoints/jwt"
)

func checkAuthorization(r *http.Request, key []byte) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return false
	}

	if headerParts[0] != "Bearer" {
		return false
	}

	_, err := jwt.ParseToken(headerParts[1], key)
	if err != nil {
		if err == jwt.ErrInvalidAccessToken {
			return false
		}

		return false
	}
	return true
}