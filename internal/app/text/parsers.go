package text

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func parseCount(r *http.Request) (size int, code int, err error) {
	vars := mux.Vars(r)

	size, _err := strconv.Atoi(vars["count"])
	if _err != nil || size < 1 || size > 100 {
		code = 412
		err = fmt.Errorf("error: count value not parsed. you can specify a value in the range [1-100]")
		return
	}
	return
}
