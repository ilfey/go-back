package json

import (
	"fmt"
	"strconv"
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

func parseLegth(q map[string][]string) (length, code int, err error) {
	if lengthStr, _err := parseQuery(q, "len"); _err != nil {
		length = 1
	} else {
		length, _err = strconv.Atoi(lengthStr)
		if _err != nil || length < 1 || length > 100 {
			code = 400
			err = fmt.Errorf("error: len value must be a number in the range [1-100]")
			return
		}
	}
	return
}
