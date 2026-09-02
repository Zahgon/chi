package middleware

import (
	"context"
	"net/http"
	"net/netip"
)

var clientIPCtxKey = &contextKey{"clientIP"}

const xForwardedForHeader = "X-Forwarded-For"

func ClientIPFromHeader(trustedHeader string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ClientIPFromXFF(trustedIPPrefixes ...string) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ClientIPFromXFFTrustedProxies(numTrustedProxies int) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func ClientIPFromRemoteAddr(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func GetClientIP(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GetClientIPAddr(ctx context.Context) netip.Addr {
	_ = "STUB: not implemented"
	return *new(netip.Addr)
}

func walkXFF(headers []string, visit func(entry string) bool) { _ = "STUB: not implemented"; return }

func inAnyPrefix(ip netip.Addr, prefixes []netip.Prefix) bool {
	_ = "STUB: not implemented"
	return false
}

func parseHeaderAddr(s string) (netip.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), false
}
