package properties

import (
	"github.com/obada-foundation/sdk-go/hash"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestNewIntProperty(t *testing.T) {
	log := log.New(os.Stdout, "TESTING SDK :: ", 0)

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

		p, _ := NewIntProperty(tc.arg, log, false)

		if p.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, p.GetValue())
		}

		hash, _ := hash.NewHash(strconv.FormatInt(tc.arg, 10), log, false)

		if p.GetHash() != hash {
			t.Fatalf("Expecting to get %v but got %v", hash, p.GetHash())
		}
	}
}
