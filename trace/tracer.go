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

// Goにおけるインターフェースは明示的にimplementsと書く必要がない
// interfaceが要求するメソッドを実装しておくことで、interfaceを実装していることになる
type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{} // 新しい構造体を返す
}
