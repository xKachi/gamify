package data

import (
	"fmt"
	"strconv"
)

type ReleaseYear int32

func (r ReleaseYear) MarshalJSON() ([]byte, error) {
	// convert to string
	jsonValue := fmt.Sprintf("year %d", r)

	// place qoutation marks
	qoutedJSONValue := strconv.Quote(jsonValue)

	return []byte(qoutedJSONValue), nil
}
