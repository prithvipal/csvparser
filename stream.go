package csvparser

import "io"

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// An Encoder writes cvs values to an output stream.
type Encoder struct {
	w io.Writer
}
