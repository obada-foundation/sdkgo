package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"github.com/obada-foundation/sdkgo/tests"
	"strconv"
	"strings"
	"testing"
)

func TestNewIntProperty(t *testing.T) {
	logger, _ := tests.CreateSdkTestLogger()

	testCases := []struct {
		arg int64
	}{
		{1},
		{2},
		{100},
		{9999999999},
		{-1},
		{0},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %d", tc.arg)

		p, _ := NewIntProperty("creating int property", tc.arg, logger, false)

		if p.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, p.GetValue())
		}

		intHash, _ := hash.NewHash([]byte(strconv.FormatInt(tc.arg, 10)), logger, false)

		if p.GetHash() != intHash {
			t.Fatalf("Expecting to get %v but got %v", intHash, p.GetHash())
		}
	}
}

func TestNewIntPropertyWithNoLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	if _, err := NewIntProperty("Adding new int prop", 10, logger, false); err != nil {
		t.Fatalf(err.Error())
	}

	logBuffStr := buff.String()

	if logBuffStr != "" {
		t.Fatalf("Expected to not receive any logs but received %q", logBuffStr)
	}
}

func TestNewIntPropertyWithLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	description := "Adding new int prop"

	if _, err := NewIntProperty(description, 10, logger, true); err != nil {
		t.Fatalf(err.Error())
	}

	buffStr := buff.String()

	substr := fmt.Sprintf("\n <|%s|> => NewIntProperty(%v)", description, 10)

	if !strings.Contains(buffStr, substr) {
		t.Errorf("Expected that log contain substring \"%s\" but couldn't find it in \"%s\"", substr, buffStr)
	}
}
