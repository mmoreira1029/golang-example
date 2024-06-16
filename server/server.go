package server

import (
	"fmt"
	"golang-service-otel-example/otelsdk"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	tracer          trace.Tracer
	requestCount    metric.Float64Counter
	requestDuration metric.Float64Histogram
	port            string
}

func NewServer(name, version, port string) (Server, error) {
	meterProvider, err := otelsdk.SetupMeterProvider(name, version)
	if err != nil {
		return Server{}, err
	}

	tracerProvider, err := otelsdk.SetupTracerProvider(name, version)
	if err != nil {
		return Server{}, err
	}

	tracer := tracerProvider.Tracer("hello-trace")

	requestCount, requestDuration, err := otelsdk.SetupMetrics(meterProvider)
	if err != nil {
		return Server{}, err
	}

	return Server{
		tracer:          tracer,
		requestCount:    requestCount,
		requestDuration: requestDuration,
		port:            port,
	}, nil

}

func (s *Server) StartServer() error {
	http.HandleFunc("/hello", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			_, span := s.tracer.Start(r.Context(), "hello-trace")
			defer span.End()

			start := time.Now()

			fmt.Println("Hello!")

			duration := time.Since(start)

			//HTTP Request Duration in Seconds
			s.requestDuration.Record(r.Context(), duration.Seconds())

			//HTTP Request Count
			s.requestCount.Add(r.Context(), 1)

		}))

	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
}
