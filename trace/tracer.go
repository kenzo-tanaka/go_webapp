package trace

import (
	"fmt"
	"io"
)

// Tracer型はTraceというメソッドを1つだけ持つインターフェース
// ...interfaceという引数は任意の型の引数を任意の数だけ持つ
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
