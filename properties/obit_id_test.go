package properties

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewObitId(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("serial_number"))

	serialNumberHash := hex.EncodeToString(h.Sum(nil))

	snh, err := NewSerialNumberHash(serialNumberHash)
	m, err := NewManufacturer("manufacturer")
	pn, err := NewPartNumber("part number")

	if err != nil {
		t.Fatalf(err.Error())
	}

	obitId, err := NewObitId(snh, m, pn)

	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedUsn := "2y5zjyCj"

	if obitId.GetUsn() != expectedUsn {
		t.Fatalf("Expected to to get usn %q but received %q", expectedUsn, obitId.GetUsn())
	}

	fmt.Println(obitId)
}
