package properties

import (
"crypto/sha256"
"encoding/hex"
"github.com/obada-protocol/sdk-go/hash"
"testing"
)

func TestSerialNumberHash(t *testing.T) {
	testCases := []struct {
		arg string
	}{
		{"serial number"},
		{"ML6843FO"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		h := sha256.New()
		_, err := h.Write([]byte(tc.arg))

		if err != nil {
			t.Errorf("Cannot get example sha256 from given %q value. %q", tc.arg, err)
		}

		serialNumberHash := hex.EncodeToString(h.Sum(nil))

		snh, _ := NewSerialNumberHash(serialNumberHash)

		if snh.GetValue() != serialNumberHash {
			t.Fatalf("Expecting to get %q but got %q", serialNumberHash, snh.GetValue())
		}

		hash, _ := hash.NewHash(serialNumberHash)

		if snh.GetHash() != hash {
			t.Fatalf("Expecting to get %v but got %v", hash, snh.GetHash())
		}
	}
}

func TestNewSerialNumberHashErrors(t *testing.T) {
	testCases := []struct {
		arg       string
		wantError string
	}{
		{"", "serial number hash must be a valid SHA256 hash"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		_, err := NewSerialNumberHash(tc.arg)

		if err.Error() != tc.wantError {
			t.Errorf("NewSerialNumberHash(%q) must return error %q, but received %q", tc.arg, tc.wantError, err.Error())
		}
	}
}

