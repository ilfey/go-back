package jwt

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/endpoints/handlers"
	"github.com/ilfey/go-back/internal/app/store/models"
	"github.com/ilfey/go-back/internal/app/store/sqlstore"
	"github.com/ilfey/go-back/internal/pkg/resp"
)

type handler struct {
	store    *sqlstore.Store
	key      []byte
	lifeSpan time.Duration
}

func New(store *sqlstore.Store, key []byte, lifeSpan int) handlers.Handler {
	return &handler{
		store:    store,
		key:      key,
		lifeSpan: time.Duration(lifeSpan),
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/jwt/register", h.handleRegister()).Methods(http.MethodPost)
	router.HandleFunc("/jwt/login", h.handleLogin()).Methods(http.MethodPost)
}

func (h *handler) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse user
		var user *models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			res := resp.NewErrorResponse(http.StatusBadRequest, "bad request")
			res.Write(w)
			return
		}

		// create user
		if err := h.store.User().Create(context.TODO(), user); err != nil {
			res := resp.NewErrorResponse(http.StatusBadRequest, "user not created")
			res.Write(w)
			return
		}

		res := resp.NewErrorResponse(http.StatusCreated, "user created")
		res.Write(w)
	}
}

func (h *handler) handleLogin() http.HandlerFunc {
	type response struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// parse user
		var user *models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			res := resp.NewErrorResponse(http.StatusBadRequest, "bad request")
			res.Write(w)
			return
		}

		var userExists *models.User
		if len(user.Email) == 0 {
			// get user by username
			var err error
			userExists, err = h.store.User().FindByUsername(context.Background(), user.Username)
			if err != nil {
				res := resp.NewErrorResponse(http.StatusUnauthorized, "user is not exists")
				res.Write(w)
				return
			}
		} else {
			// get user by email
			var err error
			userExists, err = h.store.User().FindByEmail(context.Background(), user.Email)
			if err != nil {
				res := resp.NewErrorResponse(http.StatusUnauthorized, "user is not exists")
				res.Write(w)
				return
			}
		}

		// compare passwords
		if userExists.ComparePassword(user.Password) {
			// create token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: jwt.At(time.Now().Add(time.Hour * h.lifeSpan)),
					IssuedAt:  jwt.At(time.Now()),
				},
				Username: user.Username,
			})

			// generate access token
			accessToken, err := token.SignedString(h.key)
			if err != nil {
				res := resp.NewErrorResponse(http.StatusInternalServerError, "failed to create token")
				res.Write(w)
				return
			}

			// convert to json
			j, err := json.Marshal(&response{
				Token: accessToken,
			})
			if err != nil {
				res := resp.NewErrorResponse(http.StatusInternalServerError, "error creating response")
				res.Write(w)
				return
			}

			w.WriteHeader(200)
			w.Write(j)
			return
		}

		// password is not valid
		res := resp.NewErrorResponse(http.StatusUnauthorized, "password is not valid")
		res.Write(w)
	}
}
