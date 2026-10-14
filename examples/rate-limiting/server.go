package ratelimiting

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "pong")
}

func gatedHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "gated pong")
}

// NewMux returns a mux with an ungated /ping endpoint and a rate-limited
// /gated endpoint for abuse-prevention demos.
func NewMux(limiter Limiter) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.Handle("/gated", WithRateLimit(limiter, http.HandlerFunc(gatedHandler)))

	return mux
}
