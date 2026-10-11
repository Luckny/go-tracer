// Package tracer is a debug tracing library
package tracer

import (
	"fmt"
	"io"
)

// Tracer Describes an object capable of tracing events
type Tracer interface {
	Trace(...any)
	Detail(...any)
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...any) {
	fmt.Fprintln(t.out, a...)
}

func (t *tracer) Detail(a ...any) {
	fmt.Fprint(t.out, "\t")
	fmt.Fprintln(t.out, a...)
}

// New Creates a tracer that writes to a io.Writer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (nt *nilTracer) Trace(a ...any) {}

func (nt *nilTracer) Detail(a ...any) {}

// Off Creates a tracer that ignore trace calls
func Off() Tracer {
	return &nilTracer{}
}
