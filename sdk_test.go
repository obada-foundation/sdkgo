package sdk_go

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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

	dto.SerialNumberHash = "s"
	dto.Manufacturer = "s"
	dto.PartNumber = "s"
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
		SerialNumberHash: "serialnumberhash",
		Manufacturer:     "sony",
		PartNumber:       "pn12345",
	}

	_, err = sdk.NewObitId(dto)

	if err != nil {
		t.Fatalf("Cannot create ObitId. %s", err)
	}
}

func TestSdk_NewObitIdDtoValidation(t *testing.T) {
	sdk, err := NewSdk(nil, true)

	if err != nil {
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	var dto ObitIdDto

	fmt.Println(len(dto.SerialNumberHash))

	_, err = sdk.NewObitId(dto)

	errs := err.(validator.ValidationErrors)

	for err := range errs {
		fmt.Println(errs[err])
	}
}
