package img

import (
	"image/gif"
	"image/jpeg"
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
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.jpeg", h.handleJPG())
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.jpg", h.handleJPG())
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.gif", h.handleGIF())
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
		ctx, err := createImage(params)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		// send image
		w.WriteHeader(200)
		ctx.EncodePNG(w)
	}
}

func (h *handler) handleJPG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse imageParams
		params, code, err := parseImageParams(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}
		// create image
		ctx, err := createImage(params)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		// send image
		w.WriteHeader(200)
		jpeg.Encode(w, ctx.Image(), &jpeg.Options{
			Quality: 30,
		})

	}
}

func (h *handler) handleGIF() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse imageParams
		params, code, err := parseImageParams(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}
		// create image
		ctx, err := createImage(params)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		// send image
		w.WriteHeader(200)
		gif.Encode(w, ctx.Image(), &gif.Options{
			NumColors: 256,
		})
	}
}
