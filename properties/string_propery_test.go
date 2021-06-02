package properties

import (
	"github.com/obada-foundation/sdk-go/hash"
	"testing"
)

func TestNewStringProperty(t *testing.T) {
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

		m, _ := NewStringProperty(tc.arg)

		if m.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, m.GetValue())
		}

		hash, _ := hash.NewHash(tc.arg)

		if m.GetHash() != hash {
			t.Fatalf("Expecting to get %v but got %v", hash, m.GetHash())
		}
	}
}
