package trace

import "io"

// Tracer型はTraceというメソッドを1つだけ持つインターフェース
// ...interfaceという引数は任意の型の引数を任意の数だけ持つ
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return nil
}
