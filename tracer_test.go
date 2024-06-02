package tracer

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("New returned nil")
		return
	}

	tracer.Trace("Hello from trace package.")
	if buf.String() != "Hello from trace package.\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}

}
