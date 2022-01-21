package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnil")
	} else {
		tracer.Trace("こんにちは")
		if buf.String() != "こんにちは" {
			t.Errorf("%s is not expected", buf.String())
		}
	}

	t.Error("まだテストしていない")
}
