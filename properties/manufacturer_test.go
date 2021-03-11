package properties

import (
	"github.com/obada-protocol/sdk-go/hash"
	"testing"
)

func TestNewManufacturer(t *testing.T) {
	testCases := []struct {
		arg string
	}{
		{"HP"},
		{"Apple"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		m, _ := NewManufacturer(tc.arg)

		if m.GetValue() != tc.arg {
			t.Fatalf("Expecting to get %q but got %q", tc.arg, m.GetValue())
		}

		hash, _ := hash.NewHash(tc.arg)

		if m.GetHash() != hash {
			t.Fatalf("Expecting to get %v but got %v", hash, m.GetHash())
		}
	}
}

func TestNewManufacturerErrors(t *testing.T) {
	testCases := []struct {
		arg       string
		wantError string
	}{
		{"", "manufacturer is required and cannot be empty"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		_, err := NewManufacturer(tc.arg)

		if err.Error() != tc.wantError {
			t.Errorf("NewManufacturer(%q) must return error %q, but received %q", tc.arg, tc.wantError, err.Error())
		}
	}
}
