package img

import (
	"image"
	"image/color"
	"image/png"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ilfey/go-back/internal/app/handlers"
)

type imageParams struct {
	x          int
	y          int
	tan        float64
	border     int
	background color.RGBA
	foreground color.RGBA
}

type handler struct{}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.png", h.handlePNG())
}

func (h *handler) createImage(p imageParams) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, p.x, p.y))

	for yi := img.Bounds().Min.Y; yi < img.Bounds().Max.Y; yi++ {
		for xi := img.Bounds().Min.X; xi < img.Bounds().Max.X; xi++ {
			//set background
			img.SetRGBA(xi, yi, p.background)
			// create border
			if xi > p.x-p.border || xi < p.border || yi > p.y-p.border || yi < p.border {
				img.Set(xi, yi, p.foreground)
			}
			// create x
			if int(float64(xi)*p.tan) == yi || int(float64(xi)*p.tan)+yi == p.y-1 {
				img.SetRGBA(xi, yi, p.foreground)
			}
		}
	}

	return img
}

func (h *handler) handlePNG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		x, err := strconv.Atoi(vars["x"])
		if err != nil || x == 0 {
			w.WriteHeader(500)
			w.Write([]byte("error: width value not parsed"))
		}

		y, err := strconv.Atoi(vars["y"])
		if err != nil || y == 0 {
			w.WriteHeader(500)
			w.Write([]byte("error: height value not parsed"))
		}

		// TODO create queries
		img := h.createImage(imageParams{
			x:          x,
			y:          y,
			tan:        float64(y) / float64(x),
			border:     5,
			background: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
			foreground: color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			},
		})

		w.WriteHeader(200)
		png.Encode(w, img)
	}
}
