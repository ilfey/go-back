package jwt

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
	"github.com/ilfey/go-back/internal/app/store/sqlstore"
)

type handler struct {
	store *sqlstore.Store
}

func New(store *sqlstore.Store) handlers.Handler {
	return &handler{
		store: store,
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
		if !userExists.ComparePassword(u.Password) {
			// convert to json
			j, err := json.Marshal(userExists)
			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}

			w.WriteHeader(200)
			w.Write(j)
			return
		}

		http.Error(w, "password is not valid", http.StatusUnauthorized)
	}
}
