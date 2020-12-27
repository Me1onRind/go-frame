package logger

type Tracer interface {
	PrefixFields() []*Field
}

type SimpleTracer struct {
	ReqID    string
	prefixKV []*Field
}

func NewSimpleTracer(fields ...*Field) *SimpleTracer {
	return &SimpleTracer{
		prefixKV: fields,
	}
}

func (t *SimpleTracer) PrefixFields() []*Field {
	return t.prefixKV
}
