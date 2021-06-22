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

	log := log.New(os.Stdout, "TESTING SDK :: ", 0)

	serialNumberHash := hex.EncodeToString(h.Sum(nil))

	snh, err := NewStringProperty(serialNumberHash, log, false)
	m, err := NewStringProperty("manufacturer", log, false)
	pn, err := NewStringProperty("part number", log, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	obitId, err := NewObitIDProperty(snh, m, pn, log, false)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedUsn := "2y5zjyCj"

	if obitId.GetUsn() != expectedUsn {
		t.Fatalf("Expected to to get usn %q but received %q", expectedUsn, obitId.GetUsn())
	}

	fmt.Println(obitId)
}
