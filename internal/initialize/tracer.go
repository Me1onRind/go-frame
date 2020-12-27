package initialize

import (
	"github.com/jaegertracing/jaeger-client-go/config"
	opentracing "github.com/opentracing/opentracing-go"
	"go-frame/global"
)

func SetupOpentracingTracer() error {
	var err error
	global.Tracer, err = newJaegerTracer("go-frame", "")
	if err != nil {
		return err
	}
	opentracing.SetGlobalTracer(global.Tracer)
	return nil
}

func newJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Disabled:    false,
		Tags:        []opentracing.Tag{},
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: agentHostPort,
		},
	}

	tracer, _, err := cfg.NewTracer()
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
