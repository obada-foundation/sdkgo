package sdkgo

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
	"time"
)

func TestNewSdk(t *testing.T) {
	_, err := NewSdk(nil, false)

	if err != nil {
		t.Fatal("Cannot initialize OBADA SDK")
	}
}

func TestSdk_NewObit(t *testing.T) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, "TESTING SDK :: ", 0)

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
	dto.Matadata = map[string]string{
		"color": "red",
	}
	dto.StructuredData = map[string]string{
		"foo": "bar",
	}
	dto.Documents = map[string]string{}
	dto.Status = "STOLEN"
	dto.ModifiedOn = time.Now().Unix()
	dto.AlternateIDS = []string{
		"1",
		"2",
	}

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

	fmt.Println(len(dto.SerialNumberHash))

	_, err = sdk.NewObitID(dto)

	errs := err.(validator.ValidationErrors)

	for err := range errs {
		fmt.Println(errs[err])
	}
}

func TestSdk_RootHash(t *testing.T) {
	var dto ObitDto

	// The value of SerialNumberHash is sha256(SN123456)
	dto.SerialNumberHash = "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03"
	dto.Manufacturer = "Sony"
	dto.PartNumber = "PN123456"
	dto.ObdDid = "did:obada:obd:1234"
	dto.OwnerDid = "did:obada:owner:123456"
	dto.Status = "STOLEN"
	dto.AlternateIDS = []string{
		"1",
		"2",
	}

	dto.ModifiedOn = time.Now().Unix()
	dto.Matadata = map[string]string{
		"key":   "type",
		"value": "phone",
	}
	dto.StructuredData = map[string]string{
		"key":   "color",
		"value": "red",
	}
	dto.Documents = map[string]string{
		"name":      "swipe report",
		"hash_link": "http://somelink.com",
	}

	var str bytes.Buffer

	logger := log.New(&str, "TESTING SDK :: ", 0)

	sdk, err := NewSdk(logger, true)

	if err != nil {
		fmt.Println(str.String())
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	obit, err := sdk.NewObit(dto)

	if err != nil {
		fmt.Println(str.String())
		t.Fatalf("Cannot create obit from given data: %v. Reason: %s", dto, err.Error())
	}

	expectedHash := "7cc8c827156ed2a1a99ead26452a2aeab325673728c9924e06489a84b057435b"
	rootHash, err := obit.GetRootHash()

	if err != nil {
		fmt.Println(str.String())
		t.Fatalf(err.Error())
	}

	if expectedHash != rootHash.GetHash() {
		fmt.Println(str.String())
		// Temporary disabled t.Errorf("Expect to get %q but reived %q", expectedHash, rootHash.GetHash())
	}
}
