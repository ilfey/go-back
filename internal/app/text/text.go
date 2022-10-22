package text

import (
	"net/http"

	"github.com/XANi/loremipsum"
	"github.com/gorilla/mux"

	"github.com/ilfey/go-back/internal/app/handlers"
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
	router.HandleFunc("/text/word", h.handleWord())
	router.HandleFunc("/text/sentence", h.handleSentence())
	router.HandleFunc("/text/paragraph", h.handleParagraph())
	router.HandleFunc("/text/words/{count:[0-9]+}", h.handleWords())
	router.HandleFunc("/text/sentences/{count:[0-9]+}", h.handleSentences())
	router.HandleFunc("/text/paragraphs/{count:[0-9]+}", h.handleParagraphs())
}

func (h *handler) handleWord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		word := h.txtGen.Word()
		w.Write([]byte(word))
	}
}

func (h *handler) handleSentence() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		word := h.txtGen.Sentence()
		w.Write([]byte(word))
	}
}

func (h *handler) handleParagraph() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		word := h.txtGen.Paragraph()
		w.Write([]byte(word))
	}
}

func (h *handler) handleWords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, code, err := parseCount(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		words := h.txtGen.Words(count)

		w.WriteHeader(200)
		w.Write([]byte(words))
	}
}

func (h *handler) handleSentences() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, code, err := parseCount(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		sentences := h.txtGen.Sentences(count)

		w.WriteHeader(200)
		w.Write([]byte(sentences))
	}
}

func (h *handler) handleParagraphs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, code, err := parseCount(r)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		paragraphs := h.txtGen.Paragraphs(count)

		w.WriteHeader(200)
		w.Write([]byte(paragraphs))
	}
}
