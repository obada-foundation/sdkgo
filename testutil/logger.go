package testutil

import (
	"bytes"
	"log"
)

// TestLogger creates logger that writes to bytes buffer instead of stdout
func TestLogger(prefix string) (*log.Logger, *bytes.Buffer) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, prefix, 0)

	return logger, &logStr
}
