package opentelemetry

import (
	metric "go.opentelemetry.io/otel/metric"
	trace "go.opentelemetry.io/otel/trace"
)

var WordCounterAPITracer trace.Tracer
var ApiCounter metric.Int64Counter

func RegisterTraces() {
	WordCounterAPITracer = TracerProvider.Tracer("subuTrace")
}

func RegisterMetrics() {
	ApiCounter, _ = MeterProvider.Meter("subuCounter").Int64Counter(
		"subu.counter",
		metric.WithDescription("Number of API calls."),
		metric.WithUnit("{call}"),
	)
}
