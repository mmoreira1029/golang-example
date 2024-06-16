package otelsdk

import (
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func SetupMeterProvider(name, version string) (metric.MeterProvider, error) {
	res, err := NewResource(name, version)
	if err != nil {
		return nil, err
	}

	meterProvider, err := NewMeterProvider(res)
	if err != nil {
		return nil, err
	}

	otel.SetMeterProvider(meterProvider)

	return otel.GetMeterProvider(), nil
}

func SetupTracerProvider(name, version string) (trace.TracerProvider, error) {
	res, err := NewResource(name, version)
	if err != nil {
		return nil, err
	}

	tracerProvider, err := NewTraceProvider(res)

	if err != nil {
		return nil, err
	}

	prop := NewPropagator()
	otel.SetTextMapPropagator(prop)

	otel.SetTracerProvider(tracerProvider)

	return otel.GetTracerProvider(), nil
}

func SetupMetrics(meterProvider metric.MeterProvider) (metric.Float64Counter, metric.Float64Histogram, error) {
	meter := meterProvider.Meter("http-metrics",
		metric.WithInstrumentationVersion("v0.0.0"),
	)

	requestCount, err := meter.Float64Counter(
		"request.count",
		metric.WithDescription("Number of HTTP requests"),
		metric.WithUnit("{call}"),
	)
	if err != nil {
		return nil, nil, err
	}

	requestDuration, err := meter.Float64Histogram(
		"request.duration",
		metric.WithDescription("Duration of request execution"),
		metric.WithUnit("sec"),
	)
	if err != nil {
		return nil, nil, err
	}

	return requestCount, requestDuration, nil
}

func NewResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.NewWithAttributes(semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(serviceVersion),
	), nil
}

func NewPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func NewMeterProvider(res *resource.Resource) (*sdkmetric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(metricExporter,
				sdkmetric.WithInterval(5*time.Second))),
	)
	return meterProvider, nil
}

func NewTraceProvider(res *resource.Resource) (*sdktrace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter,
			sdktrace.WithBatchTimeout(5*time.Second)),
		sdktrace.WithResource(res),
	)
	return traceProvider, nil
}
