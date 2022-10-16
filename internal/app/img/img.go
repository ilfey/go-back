package img

import (
	"image/png"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
)

type handler struct{}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.png", h.handlePNG())
}

func (h *handler) handlePNG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse imageParams
		params, code, err := parseImageParams(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}
		// create image
		img, err := createImage(params)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		// send image
		w.WriteHeader(200)
		png.Encode(w, img)
	}
}
