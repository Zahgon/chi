package middleware

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

func NewWrapResponseWriter(w http.ResponseWriter, protoMajor int) WrapResponseWriter {
	_ = "STUB: not implemented"
	return *new(WrapResponseWriter)
}

type WrapResponseWriter interface {
	http.ResponseWriter

	Status() int

	BytesWritten() int

	Tee(io.Writer)

	Unwrap() http.ResponseWriter

	Discard()
}

type basicWriter struct {
	http.ResponseWriter
	tee         io.Writer
	code        int
	bytes       int
	wroteHeader bool
	discard     bool
}

func (b *basicWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (b *basicWriter) Write(buf []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *basicWriter) maybeWriteHeader() { _ = "STUB: not implemented"; return }

func (b *basicWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (b *basicWriter) BytesWritten() int { _ = "STUB: not implemented"; return 0 }

func (b *basicWriter) Tee(w io.Writer) { _ = "STUB: not implemented"; return }

func (b *basicWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func (b *basicWriter) Discard() { _ = "STUB: not implemented"; return }

type flushWriter struct {
	basicWriter
}

func (f *flushWriter) Flush() { _ = "STUB: not implemented"; return }

var _ http.Flusher = &flushWriter{}

type hijackWriter struct {
	basicWriter
}

func (f *hijackWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

var _ http.Hijacker = &hijackWriter{}

type flushHijackWriter struct {
	basicWriter
}

func (f *flushHijackWriter) Flush() { _ = "STUB: not implemented"; return }

func (f *flushHijackWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

var _ http.Flusher = &flushHijackWriter{}
var _ http.Hijacker = &flushHijackWriter{}

type httpFancyWriter struct {
	basicWriter
}

func (f *httpFancyWriter) Flush() { _ = "STUB: not implemented"; return }

func (f *httpFancyWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (f *http2FancyWriter) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *httpFancyWriter) ReadFrom(r io.Reader) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ http.Flusher = &httpFancyWriter{}
var _ http.Hijacker = &httpFancyWriter{}
var _ http.Pusher = &http2FancyWriter{}
var _ io.ReaderFrom = &httpFancyWriter{}

type http2FancyWriter struct {
	basicWriter
}

func (f *http2FancyWriter) Flush() { _ = "STUB: not implemented"; return }

var _ http.Flusher = &http2FancyWriter{}
