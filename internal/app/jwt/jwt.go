package jwt

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
	"github.com/ilfey/go-back/internal/app/store/sqlstore"
)

type handler struct {
	store *sqlstore.Store
	key   []byte
}

func New(store *sqlstore.Store, key []byte) handlers.Handler {
	return &handler{
		store: store,
		key:   key,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/jwt/register", h.handleRegister()).Methods(http.MethodPost)
	router.HandleFunc("/jwt/login", h.handleLogin()).Methods(http.MethodPost)
}

func (h *handler) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse user
		u, err := parseUserFromBody(r)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		// create user
		if err := h.store.User().Create(context.TODO(), u); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("success"))
	}
}

func (h *handler) handleLogin() http.HandlerFunc {
	type response struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// parse user
		u, err := parseUserFromBody(r)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// get user
		userExists, err := h.store.User().FindByUsername(context.Background(), u.Username)
		if err != nil {
			http.Error(w, "user is not exists", http.StatusUnauthorized)
			return
		}

		// compare passwords
		if userExists.ComparePassword(u.Password) {
			// create token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)), // TODO add life cicle
					IssuedAt:  jwt.At(time.Now()),
				},
				Username: u.Username,
			})

			accessToken, err := token.SignedString(h.key)
			if err != nil {
				http.Error(w, "failed to create token", http.StatusInternalServerError)
				return
			}

			// convert to json
			j, err := json.Marshal(&response{
				Token: accessToken,
			})
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(200)
			w.Write(j)
			return
		}

		http.Error(w, "password is not valid", http.StatusUnauthorized)
	}
}
