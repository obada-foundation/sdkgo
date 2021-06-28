package properties

import (
	"github.com/obada-foundation/sdkgo/hash"
	"log"
	"os"
	"testing"
)

func TestNewStringProperty(t *testing.T) {

	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	testCases := []struct {
		arg string
	}{
		{"serial number"},
		{"ML6843FO"},
		{"did:obada:obd:1234"},
		{"2y5zjyCj"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		m, _ := NewStringProperty(tc.arg, logger, false)

		if m.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, m.GetValue())
		}

		strHash, _ := hash.NewHash(tc.arg, logger, false)

		if m.GetHash() != strHash {
			t.Fatalf("Expecting to get %v but got %v", strHash, m.GetHash())
		}
	}
}
