package properties

import (
	"github.com/obada-foundation/sdkgo/hash"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestNewIntProperty(t *testing.T) {
	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	testCases := []struct {
		arg int64
	}{
		{1},
		{2},
		{100},
		{9999999999},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %d", tc.arg)

		p, _ := NewIntProperty("creating int property", tc.arg, logger, false)

		if p.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, p.GetValue())
		}

		intHash, _ := hash.NewHash(strconv.FormatInt(tc.arg, 10), logger, false)

		if p.GetHash() != intHash {
			t.Fatalf("Expecting to get %v but got %v", intHash, p.GetHash())
		}
	}
}
