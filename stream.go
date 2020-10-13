package csvparser

import (
	"encoding/csv"
	"io"
)

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: csv.NewWriter(w)}
}

// An Encoder writes cvs values to an output stream.
type Encoder struct {
	w *csv.Writer
	//w io.Writer
}

// Encode method writes cvs encoding
func (e Encoder) Encode(i interface{}) {
	data := Marshal(i)
	e.w.WriteAll(data)
}
