package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidPlaytimeFormat = errors.New("invalid playtime format")

type Playtime int32

func (r Playtime) MarshalJSON() ([]byte, error) {
	// convert to string
	jsonValue := fmt.Sprintf("%d hours", r)

	// place qoutation marks
	qoutedJSONValue := strconv.Quote(jsonValue)

	return []byte(qoutedJSONValue), nil
}

// Implement a UnmarshalJSON() method on the Playtime type so that it satisfies the
// json.Unmarshaler interface. IMPORTANT: Because UnmarshalJSON() needs to modify the
// receiver (our Playtime type), we must use a pointer receiver for this to work
// correctly. Otherwise, we will only be modifying a copy (which is then discarded when
// this method returns).
func (r *Playtime) UnmarshalJSON(jsonValue []byte) error {
	// We expect that the incoming JSON value will be a string in the format
	// "<playtime> hours", and the first thing we need to do is remove the surrounding
	// double-quotes from this string. If we can't unquote it, then we return the
	// ErrInvalidRuntimeFormat error.
	unqoutedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidPlaytimeFormat
	}

	// Split the string to isolate the part containing the number.
	parts := strings.Split(unqoutedJsonValue, " ")

	// Sanity check the parts of the string to make sure it was in the expected format.
	// If it isn't, we return the ErrInvalidPlaytimeFormat error again.
	if len(parts) != 2 || parts[1] != "hours" {
		return ErrInvalidPlaytimeFormat
	}

	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidPlaytimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidPlaytimeFormat
	}
	// Convert the int32 to a Year type and assign this to the receiver. Note that we
	// use the * operator to deference the receiver (which is a pointer to a Year
	// type) in order to set the underlying value of the pointer.
	*r = Playtime(i)
	return nil
}
