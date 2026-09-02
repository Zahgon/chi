package middleware

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"sync"
)

var defaultCompressibleContentTypes = []string{
	"text/html",
	"text/css",
	"text/plain",
	"text/javascript",
	"text/markdown",
	"text/csv",
	"text/vtt",
	"application/javascript",
	"application/x-javascript",
	"application/json",
	"application/atom+xml",
	"application/rss+xml",
	"application/xml",
	"text/xml",
	"image/svg+xml",
}

func Compress(level int, types ...string) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type Compressor struct {
	encoders map[string]EncoderFunc

	pooledEncoders map[string]*sync.Pool

	allowedTypes     map[string]struct{}
	allowedWildcards map[string]struct{}

	encodingPrecedence []string
	level              int
}

func NewCompressor(level int, types ...string) *Compressor { _ = "STUB: not implemented"; return nil }

func (c *Compressor) SetEncoder(encoding string, fn EncoderFunc) { _ = "STUB: not implemented"; return }

func (c *Compressor) Handler(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (c *Compressor) selectEncoder(h http.Header, w io.Writer) (io.Writer, string, func()) {
	_ = "STUB: not implemented"
	return *new(io.Writer), "", nil
}

func matchAcceptEncoding(accepted []string, encoding string) bool {
	_ = "STUB: not implemented"
	return false
}

type EncoderFunc func(w io.Writer, level int) io.Writer

type ioResetterWriter interface {
	io.Writer
	Reset(w io.Writer)
}

type compressResponseWriter struct {
	http.ResponseWriter

	w                io.Writer
	contentTypes     map[string]struct{}
	contentWildcards map[string]struct{}
	encoding         string
	wroteHeader      bool
	compressible     bool
}

func (cw *compressResponseWriter) isCompressible() bool { _ = "STUB: not implemented"; return false }

func (cw *compressResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (cw *compressResponseWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (cw *compressResponseWriter) writer() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

type compressFlusher interface {
	Flush() error
}

func (cw *compressResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func (cw *compressResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (cw *compressResponseWriter) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (cw *compressResponseWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (cw *compressResponseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func encoderGzip(w io.Writer, level int) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func encoderDeflate(w io.Writer, level int) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}
