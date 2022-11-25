package text

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/XANi/loremipsum"
	"github.com/gorilla/mux"

	"github.com/ilfey/go-back/internal/app/endpoints/handlers"
)

type handler struct {
	txtGen *loremipsum.LoremIpsum
}

func New() handlers.Handler {
	return &handler{
		txtGen: loremipsum.New(),
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/text/word", h.handleWords())           // queries: amount=[0-100]
	router.HandleFunc("/text/sentence", h.handleSentences())   // queries: amount=[0-100]
	router.HandleFunc("/text/paragraph", h.handleParagraphs()) // queries: amount=[0-100]
}

func (h *handler) handleWords() http.HandlerFunc {
	type response struct {
		Words string `json:"words"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amounts := r.URL.Query()["amount"]

		var resp response

		// parse query: amount
		if len(amounts) != 0 {
			// parse int
			amount, err := strconv.Atoi(amounts[0])
			if err != nil || amount < 0 || amount > 100 {
				http.Error(w, "error: amount value not parsed. you can specify a value in the range [0-100]", http.StatusPreconditionFailed)
				return
			}

			// generate words
			resp = response{
				Words: h.txtGen.Words(amount),
			}
		} else {
			// generate word
			resp = response{
				Words: h.txtGen.Word(),
			}
		}

		// convert to json
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
	}
}

func (h *handler) handleSentences() http.HandlerFunc {
	type response struct {
		Sentences string `json:"sentences"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amounts := r.URL.Query()["amount"]

		var resp response

		// parse query: amount
		if len(amounts) != 0 {
			// parse int
			amount, err := strconv.Atoi(amounts[0])
			if err != nil || amount < 0 || amount > 100 {
				http.Error(w, "error: amount value not parsed. you can specify a value in the range [0-100]", http.StatusPreconditionFailed)
				return
			}

			// generate words
			resp = response{
				Sentences: h.txtGen.Sentences(amount),
			}
		} else {
			// generate word
			resp = response{
				Sentences: h.txtGen.Sentence(),
			}
		}

		// convert to json
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
	}
}

func (h *handler) handleParagraphs() http.HandlerFunc {
	type response struct {
		Paragraphs string `json:"paragraphs"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amounts := r.URL.Query()["amount"]

		var resp response

		// parse query: amount
		if len(amounts) != 0 {
			// parse int
			amount, err := strconv.Atoi(amounts[0])
			if err != nil || amount < 0 || amount > 100 {
				http.Error(w, "error: amount value not parsed. you can specify a value in the range [0-100]", http.StatusPreconditionFailed)
				return
			}

			// generate words
			resp = response{
				Paragraphs: h.txtGen.Paragraphs(amount),
			}
		} else {
			// generate word
			resp = response{
				Paragraphs: h.txtGen.Paragraph(),
			}
		}

		// convert to json
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write(b)
	}
}
