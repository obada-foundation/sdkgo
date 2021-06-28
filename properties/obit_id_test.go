package properties

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNewObitIDProperty(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("serial_number"))

	logger := log.New(os.Stdout, "TESTING SDK :: ", 0)

	serialNumberHash := hex.EncodeToString(h.Sum(nil))

	snh, err := NewStringProperty("", serialNumberHash, logger, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	m, err := NewStringProperty("", "manufacturer", logger, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	pn, err := NewStringProperty("", "part number", logger, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	obitID, err := NewObitIDProperty(snh, m, pn, logger, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedUsn := "2y5zjyCj"

	if obitID.GetUsn() != expectedUsn {
		t.Fatalf("Expected to to get usn %q but received %q", expectedUsn, obitID.GetUsn())
	}

	fmt.Println(obitID)
}
