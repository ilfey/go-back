package text

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/XANi/loremipsum"
	"github.com/gorilla/mux"

	"github.com/ilfey/go-back/internal/app/endpoints/handlers"
	"github.com/ilfey/go-back/internal/pkg/resp"
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
	router.HandleFunc("/text/{type:word|sentence|paragraph}", h.handleText()) // queries: amount=[1-100]
}

func (h *handler) handleText() http.HandlerFunc {
	type response struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amounts := r.URL.Query()["amount"]

		var amount int
		// parse query: amount
		if len(amounts) != 0 {
			// parse int
			var err error
			amount, err = strconv.Atoi(amounts[0])
			if err != nil || amount < 1 || amount > 100 {
				res := resp.New(http.StatusPreconditionFailed, "amount value not parsed. you can specify a value in the range [1-100]")
				res.Write(w)
				return
			}
		} else {
			amount = 1
		}

		// generate words
		var res response

		switch mux.Vars(r)["type"] {
		case "word":
			// generate word
			res = response{
				Text: h.txtGen.Words(amount),
			}

		case "sentence":
			// generate sentence
			res = response{
				Text: h.txtGen.Sentences(amount),
			}

		case "paragraph":
			// generate paragraph
			res = response{
				Text: h.txtGen.Paragraphs(amount),
			}
		}

		// convert to json
		j, err := json.Marshal(res)
		if err != nil {
			res := resp.New(http.StatusInternalServerError, "error creating response")
			res.Write(w)
			return
		}

		w.WriteHeader(200)
		w.Write(j)
	}
}
