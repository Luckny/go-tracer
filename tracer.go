package tracer

import (
	"fmt"
	"io"
)

// Describes an object capable of tracing events
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// Creates a tracer that writes to a io.Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (nt *nilTracer) Trace(a ...interface{}) {}

// Creates a tracer that ignore trace calls
func Off() Tracer {
	return &nilTracer{}
}
