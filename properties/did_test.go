package properties

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/obada-foundation/sdkgo/tests"
	"strings"
	"testing"
)

func sha256gen(str string) string {
	h := sha256.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func TestNewObitDIDProperty(t *testing.T) {
	testCases := []struct {
		serialNumberHash string
		manufacturer     string
		partNumber       string
		usn              string
		fullUsn          string
		logMsg           []string
		did              string
		hash             string
	}{
		{
			serialNumberHash: sha256gen("SN123456X"),
			manufacturer:     "SONY",
			partNumber:       "PN123456S",
			usn:              "31DGPLLfpckW",
			fullUsn:          "31DGPLLfpckWXfvMVVBKwEeN6Lc1X3BmaBaqyRtNQimTfLeU35eymxYyMmDStWo46Cr2UtGZoapTBsvjvTwoSxYT",
			did:              "did:obada:d7cf869423d12f623f5611e48d6f6665bbc4a270b6e09da2f4c32bcb1b949ecd",
			hash:             "d7cf869423d12f623f5611e48d6f6665bbc4a270b6e09da2f4c32bcb1b949ecd",
			logMsg: []string{
				"<|Making DID|> => NewDIDProperty",
				"Hash: d7cf869423d12f623f5611e48d6f6665bbc4a270b6e09da2f4c32bcb1b949ecd",
				"Did: did:obada:d7cf869423d12f623f5611e48d6f6665bbc4a270b6e09da2f4c32bcb1b949ecd",
				"Hash: d7cf869423d12f623f5611e48d6f6665bbc4a270b6e09da2f4c32bcb1b949ecd",
				"Full Usn: 31DGPLLfpckWXfvMVVBKwEeN6Lc1X3BmaBaqyRtNQimTfLeU35eymxYyMmDStWo46Cr2UtGZoapTBsvjvTwoSxYT",
			},
		},
	}

	for _, tc := range testCases {
		logger, buff := tests.CreateSdkTestLogger()

		snh, err := NewStringProperty("", tc.serialNumberHash, nil, false)

		if err != nil {
			t.Fatalf(err.Error())
		}

		m, err := NewStringProperty("", tc.manufacturer, nil, false)

		if err != nil {
			t.Fatalf(err.Error())
		}

		pn, err := NewStringProperty("", tc.partNumber, nil, false)

		if err != nil {
			t.Fatalf(err.Error())
		}

		DID, err := NewDIDProperty(snh, m, pn, logger, true)

		if err != nil {
			t.Fatalf(err.Error())
		}

		if DID.GetUsn() != tc.usn {
			t.Fatalf("Expected to to get usn %q but received %q", tc.usn, DID.GetUsn())
		}

		if DID.GetFullUsn() != tc.fullUsn {
			t.Fatalf("Expected to to get usn %q but received %q", tc.fullUsn, DID.GetFullUsn())
		}

		if DID.GetDid() != tc.did {
			t.Fatalf("Expected to to get did %q but received %q", tc.did, DID.GetDid())
		}

		h := DID.GetHash()

		if h.GetHash() != tc.hash {
			t.Fatalf("Expected to to get hash %q but received %q", tc.hash, h.GetHash())
		}

		buffStr := buff.String()

		for _, substr := range tc.logMsg {
			if !strings.Contains(buffStr, substr) {
				t.Errorf("Expected that log contain substring \"%s\" but couldn't find it in \"%s\"", substr, buffStr)
			}
		}
	}
}

func TestNewDIDPropertyWithNoLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	var snh, m, pn StringProperty

	if _, err := NewDIDProperty(snh, m, pn, logger, false); err != nil {
		t.Fatalf(err.Error())
	}

	buffStr := buff.String()

	if buffStr != "" {
		t.Fatalf("Expected to not receive any logs but received %q", buffStr)
	}
}
