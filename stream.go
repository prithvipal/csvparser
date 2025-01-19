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
func (e *Encoder) Encode(i interface{}) error {
	data := marshal(i)
	return e.w.WriteAll(data)
}

// NewDecoder returns a new decoder that reader from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: csv.NewReader(r)}
}

// An Decoder reads cvs values to an input stream.
type Decoder struct {
	r *csv.Reader
	//w io.Writer
}
