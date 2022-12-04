package img

import (
	"image/gif"
	"image/jpeg"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/endpoints/handlers"
	"github.com/ilfey/go-back/internal/pkg/resp"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger   *logrus.Entry
	fontPath string
}

func New(logger *logrus.Logger) handlers.Handler {

	log := logger.WithFields(logrus.Fields{
		"endpoint": "img",
	})

	fp := loadFont("arial.ttf")
	if len(fp) == 0 {
		log.Warn("No font")
	}

	return &handler{
		logger:   log,
		fontPath: fp,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/img/{x:[0-9]+}x{y:[0-9]+}.{type:png|jpg|gif}", h.handleImage()) // queries: border=[1-50], bg=[color], fg=[color]
}

func (h *handler) handleImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse imageParams
		params, code, err := parseImageParams(r)
		if err != nil {
			res := resp.NewErrorResponse(code, err.Error())
			res.Write(w)
			return
		}

		// create image
		ctx, err := createImage(params, h.fontPath)
		if err != nil {
			res := resp.NewErrorResponse(http.StatusInternalServerError, err.Error())
			res.Write(w)
			return
		}

		switch params.imageType {
		case "png":
			// send png
			w.WriteHeader(http.StatusOK)
			ctx.EncodePNG(w)

		case "jpg":
			// send jpg
			w.WriteHeader(http.StatusOK)
			jpeg.Encode(w, ctx.Image(), &jpeg.Options{
				Quality: 30,
			})

		case "gif":
			// send gif
			w.WriteHeader(http.StatusOK)
			gif.Encode(w, ctx.Image(), &gif.Options{
				NumColors: 256,
			})
		}
	}
}
