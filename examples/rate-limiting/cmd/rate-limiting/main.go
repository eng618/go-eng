// Command rate-limiting runs the super simple API with abuse prevention.
//
// One endpoint is ungated (/ping), one is rate-limited (/gated).
package main

import (
	"net/http"
	"time"

	ratelimiting "github.com/eng618/go-eng/examples/rate-limiting"
)

func main() {
	// Naive option: ratelimiting.NewFixedWindow(5, time.Second).
	// Production-friendly option below: burst 5, sustained ~5/sec.
	limiter := ratelimiting.NewTokenBucket(5, 200*time.Millisecond)

	server := &http.Server{
		Addr:              "localhost:8081",
		Handler:           ratelimiting.NewMux(limiter),
		ReadHeaderTimeout: 5 * time.Second,
	}

	_ = server.ListenAndServe()
}
