package testutil

import (
	"bytes"
	"log"
)

func TestLogger(prefix string) (*log.Logger, *bytes.Buffer) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, prefix, 0)

	return logger, &logStr
}
