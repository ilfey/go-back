package json

import (
	"fmt"
	"math"
	"strconv"
)

const (
	INT   = "int"
	FLOAT = "float"
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

func parseIntMin(q map[string][]string) (max, code int, err error) {
	if maxStr, _err := parseQuery(q, "min"); _err != nil {
		max = 0
	} else {
		max, _err = strconv.Atoi(maxStr)
		if _err != nil || max < math.MinInt32 || max > math.MaxInt32 {
			code = 400
			err = fmt.Errorf("error: min value must be a number in the range [%d-%d]", math.MinInt32, math.MaxInt32)
			return
		}
	}
	return
}

func parseIntMax(q map[string][]string) (max, code int, err error) {
	if maxStr, _err := parseQuery(q, "max"); _err != nil {
		max = 100
	} else {
		max, _err = strconv.Atoi(maxStr)
		if _err != nil || max < math.MinInt32 || max > math.MaxInt32 {
			code = 400
			err = fmt.Errorf("error: max value must be a number in the range [%d-%d]", math.MinInt32, math.MaxInt32)
			return
		}
	}
	return
}

func checkIntMinMax(min, max int) (code int, err error) {
	if min > max {
		code = 412
		err = fmt.Errorf("the min value cannot be greater than the max")
		return
	}
	return
}
