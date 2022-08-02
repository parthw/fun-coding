package main

import (
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"example.com/internal/observability"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func index(w http.ResponseWriter, r *http.Request) {
	b := bytes.NewBufferString("<h1>environment vars:</h1>")
	for _, env := range os.Environ() {
		b.WriteString(env + "<br>")
	}
	_, _ = b.WriteTo(w)
}

func main() {
	c := observability.ReadConfig()
	c.AppName = "ride-sharing-app"

	// Configure profiler.
	p, err := observability.Profiler(c)
	if err != nil {
		log.Fatalf("failed to initialize profiler: %v\n", err)
	}
	defer func() {
		_ = p.Stop()
	}()

	// Configure tracing.
	tp, err := observability.TracerProvider(c)
	if err != nil {
		log.Fatalf("failed to initialize profiler: %v\n", err)
	}
	defer func() {
		_ = tp.Shutdown(context.Background())
	}()

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(index), "IndexHandler"))
	//	http.Handle("/bike", otelhttp.NewHandler(http.HandlerFunc(bikeRoute), "BikeHandler"))

	if err = http.ListenAndServe(":5000", nil); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
