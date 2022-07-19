package utils

import (
	"encoding/json"
	"io"
)

func ToJSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func FromJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
