package optracing

import (
	"fmt"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
	"go.elastic.co/apm/module/apmhttp"
)

func RequestIDToTraceparent(requestID string) string {
	b := make([]byte, 16)
	copy(b, []byte(requestID))
	return fmt.Sprintf("00-%s--0000000000000001-01", b)
}

func RequestIDFromSpan(sm opentracing.SpanContext) string {
	carrier := opentracing.TextMapCarrier{}
	_ = opentracing.GlobalTracer().Inject(sm, opentracing.TextMap, &carrier)
	if v, ok := carrier[apmhttp.W3CTraceparentHeader]; ok {
		return RequestIDFromW3CTraceparent(v)
	}
	return ""
}

func RequestIDFromW3CTraceparent(value string) string {
	arr := strings.Split(value, "-")
	if len(arr) >= 2 {
		return arr[1]
	}
	return ""
}
