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

	snh, err := NewStringProperty(serialNumberHash)
	m, err := NewStringProperty("manufacturer")
	pn, err := NewStringProperty("part number")

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
