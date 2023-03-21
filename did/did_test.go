package did_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/obada-foundation/sdkgo/did"
	"github.com/obada-foundation/sdkgo/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDIDFromString(t *testing.T) {
	// Test not supported methods are passed
	tcs := []string{
		"",
		"did:bnb:1f4B9d871fed2dEcb2670A80237F7253DB5766De",
	}

	for _, DID := range tcs {
		_, err := did.DIDFromString(DID, nil)
		require.ErrorIs(t, err, did.ErrNotSupportedDIDMethod)
	}

}

type testCase struct {
	newDID     did.NewDID
	er         error
	did        string
	fullUSN    string
	usn        string
	hash       string
	withLogger bool
}

func TestDID(t *testing.T) {
	tcs := []testCase{
		{
			newDID: did.NewDID{
				SerialNumber: "SN123456X",
				Manufacturer: "SONY",
				PartNumber:   "PN123456S",
				Logger:       nil,
			},
			er:         nil,
			did:        "did:obada:64925be84b586363670c1f7e5ada86a37904e590d1f6570d834436331dd3eb88",
			fullUSN:    "25rc8AxGbLSrbZGXYAdKJoGXZUrn3XUZ2cM8SkZUS1AYy2f8meQ3X8HKvUzHX6sFGo2JM5jpc5ywEJLCrcip4SBh",
			usn:        "25rc8AxGbLSr",
			withLogger: false,
		},
		{
			newDID: did.NewDID{
				SerialNumber: "SN123456X",
				Manufacturer: "SONY",
				PartNumber:   "PN123456S",
				Logger:       nil,
			},
			er:         nil,
			did:        "did:obada:64925be84b586363670c1f7e5ada86a37904e590d1f6570d834436331dd3eb88",
			fullUSN:    "25rc8AxGbLSrbZGXYAdKJoGXZUrn3XUZ2cM8SkZUS1AYy2f8meQ3X8HKvUzHX6sFGo2JM5jpc5ywEJLCrcip4SBh",
			usn:        "25rc8AxGbLSr",
			hash:       "64925be84b586363670c1f7e5ada86a37904e590d1f6570d834436331dd3eb88",
			withLogger: true,
		},
	}

	for _, tc := range tcs {
		logPrefix := "DID Test :: "
		logger, loggerStr := testutil.TestLogger(logPrefix)

		if tc.withLogger {
			tc.newDID.Logger = logger
		}

		DID, err := did.MakeDID(tc.newDID)
		if tc.er == nil {
			require.NoError(t, err)
		}

		DIDFromStr, err := did.DIDFromString(DID.String(), tc.newDID.Logger)
		require.NoError(t, err)

		assert.Equal(t, tc.did, DID.String())
		assert.Equal(t, tc.fullUSN, DID.GetFullUSN())
		assert.Equal(t, tc.usn, DID.GetUSN())
		assert.Equal(t, DIDFromStr.GetHash(), DID.GetHash())

		if tc.withLogger {
			logs := strings.Split(loggerStr.String(), "\n")
			assert.Equal(
				t,
				logPrefix+fmt.Sprintf("MakeDID(%q, %q, %q)", tc.newDID.SerialNumber, tc.newDID.Manufacturer, tc.newDID.PartNumber),
				logs[0],
			)

			assert.Equal(t, fmt.Sprintf(" <|Making serialNumber|> => NewStringProperty(%q)", tc.newDID.SerialNumber), logs[2])
			assert.Equal(t, fmt.Sprintf(" <|Making manufacturer|> => NewStringProperty(%q)", tc.newDID.Manufacturer), logs[6])
			assert.Equal(t, fmt.Sprintf(" <|Making partNumber|> => NewStringProperty(%q)", tc.newDID.PartNumber), logs[10])
			assert.Equal(t, fmt.Sprintf(logPrefix+"Hash: %s", tc.hash), logs[17])
			assert.Equal(t, fmt.Sprintf(logPrefix+"DID: %s", tc.did), logs[18])
			assert.Equal(t, fmt.Sprintf(logPrefix+"USN: %s", tc.usn), logs[19])
			assert.Equal(t, fmt.Sprintf(logPrefix+"Full USN: %s", tc.fullUSN), logs[20])
		}
	}
}
