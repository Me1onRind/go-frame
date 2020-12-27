package logger

type Tracer interface {
	RequestID() string
	TraceID() string
	SpandID() string
}

type SimpleTracer struct {
	ReqID string
	TraID string
	SpaID string
}

func (t *SimpleTracer) RequestID() string {
	return t.ReqID
}
