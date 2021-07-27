package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/tests"
	"strings"
	"testing"
)

func TestNewRecord(t *testing.T) {
	logger, _ := tests.CreateSdkTestLogger()

	r, err := NewRecord("Adding new record", "color", "red", logger, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	k := r.GetKey()

	if k.GetValue() != "color" {
		t.Fatalf("Expected to to get value %q but received %q", "color", k.GetValue())
	}

	v := r.GetValue()

	if v.GetValue() != "red" {
		t.Fatalf("Expected to to get value %q but received %q", "color", v.GetValue())
	}

	h := r.GetHash()

	if h.GetHash() != "7679a2b17275bfde1fe16a977fd29044404c1094010b451c2ab1d5c5b94beeb6" {
		t.Fatalf("Expected to to get hash %q but received %q", "7679a2b17275bfde1fe16a977fd29044404c1094010b451c2ab1d5c5b94beeb6", h.GetHash())
	}
}

func TestNewRecordWithNoLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	if _, err := NewRecord("Adding new record", "color", "red", logger, false); err != nil {
		t.Fatalf(err.Error())
	}

	logBuffStr := buff.String()

	if logBuffStr != "" {
		t.Fatalf("Expected to not receive any logs but received %q", logBuffStr)
	}
}

func TestNewRecordWithLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	if _, err := NewRecord("Adding new record", "color", "red", logger, true); err != nil {
		t.Fatalf(err.Error())
	}

	buffStr := buff.String()

	substr := "TESTING SDK :: \n |Adding new record| => NewRecord(\"color\", \"red\")\n"

	if !strings.Contains(buffStr, substr) {
		t.Errorf("Expected that log contain substring \"%s\" but couldn't find it in \"%s\"", substr, buffStr)
	}

	substr = "(1948798365 + 2985630289) -> 4934428654"

	if !strings.Contains(buffStr, substr) {
		t.Errorf("Expected that log contain substring \"%s\" but couldn't find it in \"%s\"", substr, buffStr)
	}
}

func TestNewKVCollectionWithNoLog(t *testing.T) {
	logger, buff := tests.CreateSdkTestLogger()

	var kvs []KV

	if _, err := NewKVCollection("Making KV collection property", kvs, logger, false); err != nil {
		t.Fatalf(err.Error())
	}

	buffStr := buff.String()

	if buffStr != "" {
		t.Fatalf("Expected to not receive any logs but received %q", buffStr)
	}
}

func TestNewKVCollection(t *testing.T) {
	description := "Adding new KV collection"

	testCases := []struct {
		arg  []string
		kvs  []KV
		len  int
		hash string
	}{
		{
			arg: []string{
				fmt.Sprintf("\n <|%s|> => NewKVCollection(%v)", description, []KV{}),
			},
			kvs:  []KV{},
			len:  0,
			hash: "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
		},
		{
			arg: []string{
				fmt.Sprintf("\n <|%s|> => NewKVCollection(%v)", description, []KV{{
					Key:   "color",
					Value: "red",
				}}),
				description + " :: creating key/value record",
			},
			kvs: []KV{{
				Key:   "color",
				Value: "red",
			}},
			len:  1,
			hash: "1a06f97fc1cc121d27e6e48f9820495b625db5b23397ec1b3a6e0d6ec15f12c6",
		},
	}

	for _, tc := range testCases {
		logger, buff := tests.CreateSdkTestLogger()

		t.Logf("Testing: %v", tc.kvs)

		kvCollection, err := NewKVCollection(description, tc.kvs, logger, true)

		if err != nil {
			t.Fatalf(err.Error())
		}

		buffStr := buff.String()

		for _, substr := range tc.arg {
			t.Logf("Testing substring: %q", substr)

			if !strings.Contains(buffStr, substr) {
				t.Errorf("Expected that log contain substring \"%s\" but couldn't find it in \"%s\"", substr, buffStr)
			}
		}

		if len(kvCollection.GetAll()) != tc.len {
			t.Errorf("Expected to have %d elements in slice, %d given", tc.len, len(kvCollection.GetAll()))
		}

		h := kvCollection.GetHash()

		if h.GetHash() != tc.hash {
			t.Errorf("Expected to have %q hash for collection, %q given", tc.hash, h.GetHash())
		}
	}
}
