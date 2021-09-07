package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-food-api/internal/logger"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func InitTracing(serviceName string) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		logger.GlobalLogger.Fatal().Msgf("Jaeger tracer creation failed: %v", err)
	}
	opentracing.SetGlobalTracer(tracer)
}
