package jwt

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
	"github.com/ilfey/go-back/internal/app/store/models"
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
}

func (h *handler) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u models.User

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "bad request", 400)
			return
		}

		if err := h.store.User().Create(context.TODO(), &u); err != nil {
			http.Error(w, "Bad Request", 400)
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("success"))
	}
}
