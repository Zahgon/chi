package middleware

import (
	"io"
	"os"
)

var (
	nBlack   = []byte{'\033', '[', '3', '0', 'm'}
	nRed     = []byte{'\033', '[', '3', '1', 'm'}
	nGreen   = []byte{'\033', '[', '3', '2', 'm'}
	nYellow  = []byte{'\033', '[', '3', '3', 'm'}
	nBlue    = []byte{'\033', '[', '3', '4', 'm'}
	nMagenta = []byte{'\033', '[', '3', '5', 'm'}
	nCyan    = []byte{'\033', '[', '3', '6', 'm'}
	nWhite   = []byte{'\033', '[', '3', '7', 'm'}

	bBlack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	bRed     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	bGreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	bYellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	bBlue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	bMagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	bCyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	bWhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}

	reset = []byte{'\033', '[', '0', 'm'}
)

var IsTTY bool

func init() {

	fi, err := os.Stdout.Stat()
	if err == nil {
		m := os.ModeDevice | os.ModeCharDevice
		IsTTY = fi.Mode()&m == m
	}
}

func cW(w io.Writer, useColor bool, color []byte, s string, args ...any) {
	_ = "STUB: not implemented"
	return
}
