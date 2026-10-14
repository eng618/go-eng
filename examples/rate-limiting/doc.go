/*
Package ratelimiting is a super simple stdlib API gated by rate limiting for abuse prevention.

It teaches two strategies in sequence: a naive fixed-window limiter, then a
burst-friendly token-bucket limiter. Both sit in front of the API as
net/http middleware that returns 429 with Retry-After when the limit is hit.

Run the server locally with:

	go run ./examples/rate-limiting/cmd/rate-limiting

Available endpoints:

 1. GET /ping — ungated health check, always returns pong.
 2. GET /gated — rate-limited demo endpoint, returns 429 when over budget.
*/
package ratelimiting
