package middleware

import (
	"io"
	"net/http"
	"os"
)

func Recoverer(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

var recovererErrorWriter io.Writer = os.Stderr

func PrintPrettyStack(rvr any) { _ = "STUB: not implemented"; return }

func printPrettyStack(rvr any, useColor bool) { _ = "STUB: not implemented"; return }

type prettyStack struct {
}

func (s prettyStack) parse(debugStack []byte, rvr any, useColor bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s prettyStack) decorateLine(line string, useColor bool, num int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s prettyStack) decorateFuncCallLine(line string, useColor bool, num int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s prettyStack) decorateSourceLine(line string, useColor bool, num int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
