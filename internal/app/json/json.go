package json

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/ilfey/go-back/internal/app/handlers"
)

type handler struct{}

func New() handlers.Handler {
	rand.Seed(time.Now().UnixNano())
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	// TODO implement routes
	// router.Handle("/json/object", h.handleObject())
	// router.Handle("/json/array/object", h.handleArrayOfObject())
	// router.Handle("/json/array/string", h.handleArrayOfString())
	// router.Handle("/json/array/number", h.handleArrayOfNumber())
	router.Handle("/json/array/boolean", h.handleArrayOfBoolean()) // queries: len=[1-100]
}

func (h *handler) handleArrayOfBoolean() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		var length int

		// parse length
		if lenStr, err := parseQuery(queries, "len"); err != nil {
			length = 1
		} else {
			length, err = strconv.Atoi(lenStr)
			if err != nil || length < 1 || length > 100 {
				w.WriteHeader(400)
				w.Write([]byte("error: len value must be a number in the range [1-100]"))
				return
			}
		}

		// create and fill array
		arr := make([]bool, length)
		for i := 0; i < length; i++ {
			arr[i] = rand.Intn(2) == 1
		}

		// serialize array
		j, _ := json.Marshal(arr)

		w.WriteHeader(200)
		w.Write(j)
	}
}
