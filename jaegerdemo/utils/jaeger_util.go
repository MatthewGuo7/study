package utils

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"log"
)

type TraceObjectStruct struct {
	tracer opentracing.Tracer
	closer io.Closer
}

func NewTraceObjectStruct(tracer opentracing.Tracer, closer io.Closer) *TraceObjectStruct {
	return &TraceObjectStruct{tracer: tracer, closer: closer}
}

var TraceObject *TraceObjectStruct

func InitJaeger() {
	cfg := config.Configuration{
		ServiceName:         "OrderService",
		Disabled:            false,
		RPCMetrics:          false,
		Tags:                nil,
		Sampler:             &config.SamplerConfig{
			Type:                     "const",
			Param:                    1,
			SamplingServerURL:        "",
			SamplingRefreshInterval:  0,
			MaxOperations:            0,
			OperationNameLateBinding: false,
			Options:                  nil,
		},
		Reporter:            &config.ReporterConfig{
			QueueSize:                  0,
			BufferFlushInterval:        0,
			LogSpans:                   false,
			LocalAgentHostPort:         "127.0.0.1:6831" ,
			DisableAttemptReconnecting: false,
			AttemptReconnectInterval:   0,
			CollectorEndpoint:          "",
			User:                       "",
			Password:                   "",
			HTTPHeaders:                nil,
		},
		Headers:             nil,
		BaggageRestrictions: nil,
		Throttler:           nil,
	}
	tracer, closer, err := cfg.NewTracer()
	if nil != err {
		log.Fatal(err)
	}

	TraceObject = NewTraceObjectStruct(tracer, closer)
	opentracing.SetGlobalTracer(tracer)
}
