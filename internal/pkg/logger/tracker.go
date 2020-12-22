package logger

type Tracer interface {
	RequestID() string
}

type SimpleTracer struct {
	ReqID string
}

func (t *SimpleTracer) RequestID() string {
	return t.ReqID
}
