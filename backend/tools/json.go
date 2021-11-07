package tools

import (
	"encoding/json"
	"io"
)

// ToJSON takes in an interface and converts it to JSON object
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON converts JSON object to interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(r)
}
