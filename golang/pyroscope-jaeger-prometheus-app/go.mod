module example.com

go 1.18

replace example.com/internal/observability => ./internal/observability

require (
	github.com/pyroscope-io/client v0.3.0
	github.com/pyroscope-io/otel-profiling-go v0.4.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.33.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.8.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.8.0
	go.opentelemetry.io/otel/sdk v1.8.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
)
