package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/tests"
	"reflect"
	"strings"
	"testing"
)

func TestNewSliceStrProperty(t *testing.T) {
	description := "Creating SliceStrProperty"

	testCases := []struct {
		sliceStr []string
		logMsg   []string
		hash     string
	}{
		{
			sliceStr: []string{},
			logMsg: []string{
				fmt.Sprintf("\n <|%s|> => NewSliceStrProperty(%v)", description, []string{}),
			},
			hash: "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
		},
		{
			sliceStr: []string{"1", "2", "3"},
			logMsg: []string{
				fmt.Sprintf("\n <|%s|> => NewSliceStrProperty(%v)", description, []string{"1", "2", "3"}),
				"(1803989619 + 3564330554 + 1309098117) => 6677418290",
			},
			hash: "0224aa11c8ec1956cb7feb1528a1556ee4b351cf48a7f754a6434142a2a77a9b",
		},
	}

	for _, tc := range testCases {
		logger, buff := tests.CreateSdkTestLogger()

		ssp, err := NewSliceStrProperty(description, tc.sliceStr, logger, true)

		if err != nil {
			t.Fatalf(err.Error())
		}

		if !reflect.DeepEqual(tc.sliceStr, ssp.GetValue()) {
			t.Fatalf("Expected to to get usn %q but received %q", tc.sliceStr, ssp.GetValue())
		}

		h := ssp.GetHash()

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

func TestNewSliceStrPropertyWithNoLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	var sliceStr []string

	if _, err := NewSliceStrProperty("Creating SliceStrProperty", sliceStr, logger, false); err != nil {
		t.Fatalf(err.Error())
	}

	buffStr := buff.String()

	if buffStr != "" {
		t.Fatalf("Expected to not receive any logs but received %q", buffStr)
	}
}
