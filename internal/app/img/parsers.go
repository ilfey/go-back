package img

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func parseQuery(q map[string][]string, p string) (val string, err error) {
	vals := q[p]
	if len(vals) == 0 {
		err = fmt.Errorf("no %s parameter", p)
		return
	}
	val = vals[0]
	return
}

func parseImageParams(r *http.Request) (params *imageParams, code int, err error) {
	queries := r.URL.Query()
	vars := mux.Vars(r)

	// parse x
	x, _err := strconv.Atoi(vars["x"])
	if _err != nil || x == 0 || x > 2000 {
		code = 412
		err = fmt.Errorf("error: width value not parsed. you can specify a value in the range [1-2000]")
		return
	}

	// parse y
	y, _err := strconv.Atoi(vars["y"])
	if _err != nil || y == 0 || y > 2000 {
		code = 412
		err = fmt.Errorf("error: width value not parsed. you can specify a value in the range [1-2000]")
		return
	}

	// parse border
	var border int

	if borderStr, _err := parseQuery(queries, "border"); _err != nil {
		border = 5
	} else {
		border, err = strconv.Atoi(borderStr)
		if err != nil || border <= 0 || border > 50 {
			code = 400
			err = fmt.Errorf("error: border value must be a number in the range [1-50]")
			return
		}
	}

	//parse background
	bg, _err := parseQuery(queries, "bg")
	if _err != nil {
		bg = "fff"
	}

	// parse foreground
	fg, _err := parseQuery(queries, "fg")
	if _err != nil {
		fg = "000"
	}

	params = &imageParams{
		x:      x,
		y:      y,
		bg:     bg,
		fg:     fg,
		tan:    float64(y) / float64(x),
		border: float64(border),
	}

	return
}
