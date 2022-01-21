package trace

import "io"

// Tracer型はTraceというメソッドを1つだけ持つインターフェース
// ...interfaceという引数は任意の型の引数を任意の数だけ持つ
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
