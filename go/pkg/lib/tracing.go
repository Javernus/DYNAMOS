package lib

import (
	"context"
	"fmt"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

func InitTracer(serviceName string) (*ocagent.Exporter, error) {
	ocagentHost := os.Getenv("OC_AGENT_HOST")
	if ocagentHost == "" {
		return nil, fmt.Errorf("env OC_AGENT_HOST not declared")
	}

	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(5*time.Second),
		ocagent.WithAddress(ocagentHost),
		ocagent.WithServiceName(serviceName),
	)

	trace.RegisterExporter(oce)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return oce, err
}

func StartRemoteParentSpan(ctx context.Context, name string, parentTrace []byte) (context.Context, *trace.Span, error) {

	// Deserialize the span context
	spanContext, ok := propagation.FromBinary(parentTrace)
	if !ok {
		return nil, nil, fmt.Errorf("invalid span context")
	}

	ctx, span := trace.StartSpanWithRemoteParent(ctx, name, spanContext)
	return ctx, span, nil
}
