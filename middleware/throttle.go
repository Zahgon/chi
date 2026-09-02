package middleware

import (
	"net/http"
	"time"
)

const (
	errCapacityExceeded = "Server capacity exceeded."
	errTimedOut         = "Timed out while waiting for a pending request to complete."
	errContextCanceled  = "Context was canceled."
)

var (
	defaultBacklogTimeout = time.Second * 60
)

type ThrottleOpts struct {
	RetryAfterFn   func(ctxDone bool) time.Duration
	Limit          int
	BacklogLimit   int
	BacklogTimeout time.Duration
	StatusCode     int
}

func Throttle(limit int) func(http.Handler) http.Handler { _ = "STUB: not implemented"; return nil }

func ThrottleBacklog(limit, backlogLimit int, backlogTimeout time.Duration) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ThrottleWithOpts(opts ThrottleOpts) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type token struct{}

type throttler struct {
	tokens         chan token
	backlogTokens  chan token
	retryAfterFn   func(ctxDone bool) time.Duration
	backlogTimeout time.Duration
	statusCode     int
}

func (t throttler) setRetryAfterHeaderIfNeeded(w http.ResponseWriter, ctxDone bool) {
	_ = "STUB: not implemented"
	return
}
