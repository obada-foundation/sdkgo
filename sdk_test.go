package sdk_go

import (
	"testing"
)

func TestNewSdk(t *testing.T) {
	_, err := NewSdk(nil, false)

	if err != nil {
		t.Fatal("Cannot initialize OBADA SDK")
	}
}

func TestSdk_NewObit(t *testing.T) {
	sdk, err := NewSdk(nil, false)

	if err != nil {
		t.Fatal("Cannot initialize OBADA SDK")
	}

	var dto ObitDto

	dto.serialNumberHash = "s"
	dto.manufacturer = "s"
	dto.partNumber = "s"
	dto.ownerDid = "did:obada:owner:123456"
	dto.obdDid = "did:obada:obd:1234"

	_, err = sdk.NewObit(dto)

	if err != nil {

	}
}

func TestSdk_NewObitId(t *testing.T) {
	sdk, err := NewSdk(nil, true)

	if err != nil {
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	dto := ObitIdDto{
		serialNumberHash: "serialnimhash",
		manufacturer:     "sony",
		partNumber:       "pn12345",
	}

	_, err = sdk.NewObitId(dto)

	if err != nil {
		t.Fatalf("Cannot create ObitId. %s", err)
	}
}
