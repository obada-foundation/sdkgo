package sdkgo

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/obada-foundation/sdkgo/properties"
	"github.com/obada-foundation/sdkgo/tests"
)

func TestSdk_NewObit(t *testing.T) {
	logger, _ := tests.CreateSdkTestLogger()

	sdk := NewSdk(logger, true)

	DID, err := sdk.NewObitDID("s", "s", "s")
	if err != nil {
		t.Fatalf("Cannot create an DID. %s", err)
	}

	docs := properties.NewDocumentsCollection(logger, true)

	if _, err := sdk.NewObit(DID, docs); err != nil {
		t.Fatalf("Cannot create an Obit. %s", err)
	}
}

func TestSdk_NewObitDID(t *testing.T) {
	var logStr bytes.Buffer

	logger := log.New(&logStr, "TESTING SDK :: ", 0)

	sdk := NewSdk(logger, true)

	if _, err := sdk.NewObitDID("serialnumberhash", "sony", "pn12345"); err != nil {
		fmt.Println(logStr.String())
		t.Fatalf("Cannot create ObitID. %s", err)
	}

	fmt.Println(logStr.String())
}
