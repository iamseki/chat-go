package main

import (
	"fmt"
	"io"
)

// Tracer interface
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct {
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *nilTracer) Trace(a ...interface{}) {
}

// Off bla
func Off() Tracer {
	return &nilTracer{}
}

// New bla
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
