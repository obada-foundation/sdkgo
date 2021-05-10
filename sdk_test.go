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

	_, err = sdk.NewObit("s", "s", "s", "did:obada:owner:123456", "did:obada:obd:1234")

	if err != nil {

	}
}

func TestSdk_NewObitId(t *testing.T) {
	sdk, err := NewSdk(nil, true)

	if err != nil {
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	_, err = sdk.NewObit("serialnimhash", "sony", "pn12345", "did:obada:owner:123456", "did:obada:obd:1234")

	if err != nil {
		t.Fatalf("Cannot create ObitId. %s", err)
	}
}
