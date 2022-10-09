package img

import (
	"fmt"
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
			if int(float64(xi)*p.tan) == yi || int(float64(xi)*p.tan)+yi == p.y {
				img.SetRGBA(xi, yi, p.foreground)
			}
		}
	}

	return img
}

func (h *handler) parseHexColor(s string) color.RGBA {
	var c color.RGBA
	c.A = 255
	switch len(s) {
	case 9:
		fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 7:
		fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 5:
		fmt.Sscanf(s, "#%1x%1x%1x%1x", &c.R, &c.G, &c.B, &c.A)
		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A *= 17

	case 4:
		fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	}
	return c
}

func (h *handler) parseQuery(q map[string][]string, p string) (val string, err error) {
	vals := q[p]
	if len(vals) == 0 {
		err = fmt.Errorf("no %s parameter", p)
		return
	}
	val = vals[0]
	return
}

func (h *handler) handlePNG() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		vars := mux.Vars(r)

		x, err := strconv.Atoi(vars["x"])
		if err != nil || x == 0 || x > 2000 {
			w.WriteHeader(412)
			w.Write([]byte("error: width value not parsed. you can specify a value in the range (1-2000)"))
			return
		}

		y, err := strconv.Atoi(vars["y"])
		if err != nil || y == 0 || y > 2000 {
			w.WriteHeader(412)
			w.Write([]byte("error: height value not parsed. you can specify a value in the range (1-2000)"))
			return
		}

		var border int
		if borderStr, err := h.parseQuery(queries, "border"); err != nil {
			border = 5
		} else {
			border, err = strconv.Atoi(borderStr)
			if err != nil || border <= 0 || border > 50 {
				w.WriteHeader(400)
				w.Write([]byte("error: border value must be a number in the range (1, 50)"))
				return
			}
		}

		// TODO create queries
		img := h.createImage(imageParams{
			x:          x,
			y:          y,
			tan:        float64(y) / float64(x),
			border:     border,
			background: h.parseHexColor("#fff"),
			foreground: h.parseHexColor("#000"),
		})

		w.WriteHeader(200)
		png.Encode(w, img)
	}
}
