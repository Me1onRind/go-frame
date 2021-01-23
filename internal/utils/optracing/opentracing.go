package optracing

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/metadata"
	//"go.opentelemetry.io/otel/propagation"
)

type OpentracingCarrier struct {
	data map[string]string
}

func NewOpentracingCarrierFromGrpcContext(ctx context.Context) *OpentracingCarrier {
	md, _ := metadata.FromContext(ctx)
	fmt.Println(md)
	g := &OpentracingCarrier{
		data: md,
	}
	if g.data == nil {
		g.data = map[string]string{}
	}
	return g
}

func NewOpentracingCarrier() *OpentracingCarrier {
	g := &OpentracingCarrier{
		data: map[string]string{},
	}
	return g
}

func (g *OpentracingCarrier) Get(key string) string {
	fmt.Println("??????????")
	return g.data[key]
}

func (g *OpentracingCarrier) Set(key, value string) {
	fmt.Println("??????????")
	g.data[key] = value
}
