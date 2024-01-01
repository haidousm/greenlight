package data

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	re := regexp.MustCompile(`(?:(\d+) mins?"$)`)

	matches := re.FindStringSubmatch(string(jsonValue))
	if matches == nil || len(matches) > 2 {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}
