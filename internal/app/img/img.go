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
	x      int
	y      int
	bg     string
	fg     string
	tan    float64
	border int
}

type handler struct{}

func New() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.png", h.handlePNG())
}

func (h *handler) createImage(p imageParams) (*image.RGBA, error) {
	bg, err := h.parseHexColor(p.bg)
	if err != nil {
		return nil, err
	}
	fg, err := h.parseHexColor(p.fg)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, p.x, p.y))

	for yi := img.Bounds().Min.Y; yi < img.Bounds().Max.Y; yi++ {
		for xi := img.Bounds().Min.X; xi < img.Bounds().Max.X; xi++ {
			dx := int(float64(xi) * p.tan)
			// cx := p.x / 2
			empty := true // TODO (cx-cx/5 > dx || cx+cx/5 < dx)
			hb := p.border / 2

			//set background
			img.SetRGBA(xi, yi, bg)
			// create border
			if xi > p.x-p.border || xi < p.border || yi > p.y-p.border || yi < p.border {
				img.Set(xi, yi, fg)
			}
			// create x
			if sum := dx + yi; dx-hb < yi && dx+hb > yi && empty || sum > p.y-hb && sum < p.y+hb && empty {
				img.SetRGBA(xi, yi, fg)
			}
			// TODO set Text
		}
	}

	return img, nil
}

func (h *handler) parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 255
	switch len(s) {
	case 8: // ff008833
		_, err = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 6: // ffbbaa
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4: // ff05
		_, err = fmt.Sscanf(s, "%1x%1x%1x%1x", &c.R, &c.G, &c.B, &c.A)
		c.R *= 17
		c.G *= 17
		c.B *= 17
		c.A *= 17
	case 3: // 880
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("color parsing error. check the spelling of the color (fff, 0000, ffffff, 00000000)")
	}
	return
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
			w.Write([]byte("error: width value not parsed. you can specify a value in the range [1-2000]"))
			return
		}

		y, err := strconv.Atoi(vars["y"])
		if err != nil || y == 0 || y > 2000 {
			w.WriteHeader(412)
			w.Write([]byte("error: height value not parsed. you can specify a value in the range [1-2000]"))
			return
		}

		var border int
		if borderStr, err := h.parseQuery(queries, "border"); err != nil {
			border = 5
		} else {
			border, err = strconv.Atoi(borderStr)
			if err != nil || border <= 0 || border > 50 {
				w.WriteHeader(400)
				w.Write([]byte("error: border value must be a number in the range [1-50]"))
				return
			}
		}

		bg, err := h.parseQuery(queries, "bg")
		if err != nil {
			bg = "fff"
		}

		fg, err := h.parseQuery(queries, "fg")
		if err != nil {
			fg = "000"
		}

		img, err := h.createImage(imageParams{
			x:      x,
			y:      y,
			bg:     bg,
			fg:     fg,
			tan:    float64(y) / float64(x),
			border: border,
		})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		png.Encode(w, img)
	}
}
