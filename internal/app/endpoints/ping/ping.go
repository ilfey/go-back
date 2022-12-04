package ping

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/endpoints/handlers"
	"github.com/ilfey/go-back/internal/pkg/resp"
)

type handler struct {
	Key []byte
}

func New(key []byte) handlers.Handler {
	return &handler{
		Key: key,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/ping", h.handlePing())
}

func (h *handler) handlePing() http.HandlerFunc {
	type response struct {
		IsAuthorized bool   `json:"is_authorized"`
		Code         int    `json:"code"`
		Message      string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		res := response{}

		// check authorization
		res.IsAuthorized = checkAuthorization(r, h.Key)
		if !res.IsAuthorized {
			w.WriteHeader(http.StatusUnauthorized)
			res.Code = http.StatusUnauthorized
			res.Message = "you are not authorized"
		} else {
			res.Code = http.StatusOK
			res.Message = "ok"
		}

		j, err := json.Marshal(&res)
		if err != nil {
			res := resp.New(http.StatusInternalServerError, "error creating response")
			res.Write(w)
			return
		}

		w.Write(j)
	}
}
