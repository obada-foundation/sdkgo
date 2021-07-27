package tests

import (
	"bytes"
	"log"
)

// CreateSdkTestLogger Creates a logger that is used in tests
func CreateSdkTestLogger() (*log.Logger, *bytes.Buffer) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, "TESTING SDK :: ", 0)

	return logger, &logStr
}
