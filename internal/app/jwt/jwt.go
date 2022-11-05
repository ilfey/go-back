package jwt

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
)

type handler struct{}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/jwt/register", h.handleRegister())
}

func (h *handler) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("not implemented"))
	}
}
