package hash

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/obada-foundation/sdkgo/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHash(t *testing.T) {
	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	testCases := []struct {
		arg         string
		wantSha256  string
		wantDecimal uint64
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", 3820012610},
		{"Some value", "cc68b65c670cea83c4b9e110822af132258d882b7bc79c3f3645bdec06131e71", 3429414492},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		h, _ := NewHash([]byte(tc.arg), logger)

		if h.GetHash() != tc.wantSha256 {
			t.Errorf("getHash() = %q, want %q", h.GetHash(), tc.wantSha256)
		}

		if h.GetDec() != tc.wantDecimal {
			t.Errorf("getDec() = %d, want %d", h.GetDec(), tc.wantDecimal)
		}
	}

}

func TestErrorsWhenTryingConvertNonHexToDec(t *testing.T) {
	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	testCases := []struct {
		arg       string
		wantError string
	}{
		{"not", "given string string \"not\" is not valid hex"},
		{"not valid", "given string string \"not valid\" is not valid hex"},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		got, err := hashToDec(tc.arg, logger)

		if got != 0 {
			t.Errorf("hashToDec(%q) = %d, want %d", tc.arg, got, 0)
		}

		if err.Error() != tc.wantError {
			t.Errorf("hashToDec(%q) must return error %q, but received %q", tc.arg, tc.wantError, err.Error())
		}
	}
}

func TestHashToDecimalConversion(t *testing.T) {
	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	testCases := []struct {
		arg  string
		want uint64
	}{
		{"0000", 0},
		{"ff", 255},
		{"100", 256},
		{"ffff", 65535},
		{"ffffff", 16777215},
		{"ffffffff", 4294967295},
		{"ffffffffaa", 4294967295},
	}

	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)

		got, err := hashToDec(tc.arg, logger)

		if err != nil {
			t.Fatalf("Cannot convert %s to decimal. %s", tc.arg, err.Error())
		}

		if got != tc.want {
			t.Fatalf("hashToDec(%q) = %d, want %d", tc.arg, got, tc.want)
		}
	}
}

type sumHashesTc struct {
	hashes     []string
	sum        uint64
	withLogger bool
	logger     *log.Logger
}

func TestSumHashes(t *testing.T) {
	testCases := []sumHashesTc{
		{
			hashes:     make([]string, 0),
			sum:        uint64(0),
			withLogger: false,
			logger:     nil,
		},
		{
			hashes:     make([]string, 0),
			sum:        uint64(0),
			withLogger: true,
			logger:     nil,
		},
		{
			hashes: []string{
				"7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed",
			},
			sum:        uint64(1989329837),
			withLogger: false,
			logger:     nil,
		},
	}

	for _, tc := range testCases {
		logPrefix := "Hashes sum :: "
		logger, loggerStr := testutil.TestLogger(logPrefix)

		if tc.withLogger {
			tc.logger = logger
		}

		hashes := make([]Hash, 0)

		for _, hashStr := range tc.hashes {
			h, err := FromString(hashStr, nil)
			require.NoError(t, err)

			hashes = append(hashes, h)
		}

		sum := SumHashes(tc.logger, hashes...)
		assert.Equal(t, tc.sum, sum)

		if tc.withLogger {
			logs := strings.Split(loggerStr.String(), "\n")

			assert.Equal(t, fmt.Sprint("Hashes sum :: "), logs[0])
			assert.Equal(t, fmt.Sprintf("<|Computing sum of hashes|> => SumHashes([]) -> ([]) -> %d", sum), logs[1])
		}
	}
}

func TestFromString(t *testing.T) {}
