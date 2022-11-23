package jwt

import (
	"encoding/json"
	"net/http"

	"github.com/ilfey/go-back/internal/app/store/models"
)

func parseUserFromBody(r *http.Request) (u *models.User, err error) {
	if err = json.NewDecoder(r.Body).Decode(&u); err != nil {
		return
	}
	return
}
