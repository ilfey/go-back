package json

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/XANi/loremipsum"
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
	router.Handle("/json/array/string", h.handleArrayOfString())   // queries: len=[1-100]
	router.Handle("/json/array/number", h.handleArrayOfNumber())   // queries: len=[1-100], min=[math.MinInt32-math.MaxInt32]
	router.Handle("/json/array/boolean", h.handleArrayOfBoolean()) // queries: len=[1-100]
}

func (h *handler) handleArrayOfString() http.HandlerFunc {
	li := loremipsum.New()

	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		// parse length
		length, code, err := parseLegth(queries)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		arr := li.WordList(length)

		// swap arr
		rand.Shuffle(length, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		// serialize array
		j, _ := json.Marshal(arr)

		w.WriteHeader(200)
		w.Write(j)
	}
}

func (h *handler) handleArrayOfNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		// parse length
		length, code, err := parseLegth(queries)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		var j []byte

		// parse min
		min, code, err := parseIntMin(queries)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		// parse max
		max, code, err := parseIntMax(queries)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		// check min and max values
		code, err = checkIntMinMax(min, max)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		arr := make([]int, length)

		for index := 0; index < length; index++ {
			arr[index] = rand.Intn(max-min+1) + min
		}

		j, _ = json.Marshal(arr)

		w.WriteHeader(200)
		w.Write(j)
	}
}

func (h *handler) handleArrayOfBoolean() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		// parse length
		length, code, err := parseLegth(queries)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
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
