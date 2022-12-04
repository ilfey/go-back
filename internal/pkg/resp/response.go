package resp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

func (resp *ErrorResponse) ToJson() ([]byte, error) {
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (resp *ErrorResponse) ToString() []byte {
	return []byte(fmt.Sprintf("%d - %s", resp.Code, resp.Message))
}

func (resp *ErrorResponse) Write(w http.ResponseWriter) {
	b, err := resp.ToJson()
	if err != nil {
		w.Header().Add("Content-Type", "text/plain")
		b = resp.ToString()
	} else {
		w.Header().Add("Content-Type", "application/json")
	}

	w.WriteHeader(resp.Code)
	w.Write(b)
}
