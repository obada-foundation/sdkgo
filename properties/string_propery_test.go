package properties

import (
	"github.com/obada-foundation/sdk-go/hash"
	"log"
	"os"
	"testing"
)

func TestNewStringProperty(t *testing.T) {

	log := log.New(os.Stdout, "TESTING SDK :: ", 0)

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

		m, _ := NewStringProperty(tc.arg, log, false)

		if m.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, m.GetValue())
		}

		hash, _ := hash.NewHash(tc.arg, log, false)

		if m.GetHash() != hash {
			t.Fatalf("Expecting to get %v but got %v", hash, m.GetHash())
		}
	}
}
