package sdk_go

import "testing"

func TestNewObit(t *testing.T) {

	_, err := NewObit("some serial", "Sony", "PN1257F")

	if err != nil {
		t.Errorf("cannot create obit: %s", err)
	}
}
