package sdkgo

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/obada-foundation/sdkgo/properties"
	"github.com/obada-foundation/sdkgo/tests"
	"log"
	"testing"
)

func TestNewSdk(t *testing.T) {
	if _, err := NewSdk(nil, false); err != nil {
		t.Fatal("Cannot initialize OBADA SDK")
	}
}

func TestSdk_NewObit(t *testing.T) {
	logger, _ := tests.CreateSdkTestLogger()

	sdk, err := NewSdk(logger, true)

	if err != nil {
		t.Fatal("Cannot initialize OBADA SDK")
	}

	var dto ObitDto

	dto.SerialNumberHash = "s"
	dto.Manufacturer = "s"
	dto.PartNumber = "s"
	dto.OwnerDid = "did:obada:owner:123456"
	dto.ObdDid = "did:obada:obd:1234"
	dto.Documents = []properties.Doc{}

	_, err = sdk.NewObit(dto)

	if err != nil {
		t.Fatalf("Cannot create an Obit. %s", err)
	}
}

func TestSdk_NewObitID(t *testing.T) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, "TESTING SDK :: ", 0)

	sdk, err := NewSdk(logger, true)

	if err != nil {
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	dto := ObitIDDto{
		SerialNumberHash: "serialnumberhash",
		Manufacturer:     "sony",
		PartNumber:       "pn12345",
	}

	_, err = sdk.NewObitID(dto)

	if err != nil {
		fmt.Println(logStr.String())
		t.Fatalf("Cannot create ObitID. %s", err)
	}

	fmt.Println(logStr.String())
}

func TestSdk_ObitIDDtoValidation(t *testing.T) {
	sdk, err := NewSdk(nil, false)

	if err != nil {
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	var dto ObitIDDto

	_, err = sdk.NewObitID(dto)

	errs := err.(validator.ValidationErrors)

	for err := range errs {
		fmt.Println(errs[err])
	}
}
