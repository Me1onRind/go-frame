package logger

type Tracer interface {
	PrefixFields() []*Field
}

type SimpleTracer struct {
	string
	prefixKV []*Field
}

func NewSimpleTracer(fields ...*Field) Tracer {
	return &SimpleTracer{
		prefixKV: fields,
	}
}

func (t *SimpleTracer) PrefixFields() []*Field {
	return t.prefixKV
}
