package middleware

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	LogEntryCtxKey = &contextKey{"LogEntry"}

	DefaultLogger func(next http.Handler) http.Handler
)

func Logger(next http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func RequestLogger(f LogFormatter) func(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

type LogFormatter interface {
	NewLogEntry(r *http.Request) LogEntry
}

type LogEntry interface {
	Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{})
	Panic(v interface{}, stack []byte)
}

func GetLogEntry(r *http.Request) LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func WithLogEntry(r *http.Request, entry LogEntry) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

type LoggerInterface interface {
	Print(v ...interface{})
}

type DefaultLogFormatter struct {
	Logger  LoggerInterface
	NoColor bool
}

func (l *DefaultLogFormatter) NewLogEntry(r *http.Request) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

type defaultLogEntry struct {
	*DefaultLogFormatter
	request  *http.Request
	buf      *bytes.Buffer
	useColor bool
}

func (l *defaultLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l *defaultLogEntry) Panic(v interface{}, stack []byte) { _ = "STUB: not implemented"; return }

func init() {
	color := true
	if runtime.GOOS == "windows" {
		color = false
	}
	DefaultLogger = RequestLogger(&DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: !color})
}
