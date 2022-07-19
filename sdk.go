package sdkgo

import (
	"log"

	"github.com/obada-foundation/sdkgo/properties"
)

// Sdk OBADA SDK
type Sdk struct {
	logger *log.Logger
	debug  bool
}

// Obit represent asset data structure
type Obit struct {
	obitDID   properties.DID
	documents properties.Documents
	debug     bool
	logger    *log.Logger
}

// NewSdk creates a new OBADA SDK instance
func NewSdk(logger *log.Logger, debug bool) *Sdk {
	return &Sdk{
		logger: logger,
		debug:  debug,
	}
}

// NewDocument create SDK document property
func (sdk *Sdk) NewDocument(name, link, docChecksum string) (properties.Document, error) {
	return properties.NewDocument(name, link, docChecksum, sdk.logger, sdk.debug)
}

// NewDocuments create documents collection
func (sdk *Sdk) NewDocuments(docs []properties.Document) properties.Documents {
	documents := properties.NewDocumentsCollection(sdk.logger, sdk.debug)

	for _, doc := range docs {
		documents.AddDocument(doc)
	}

	return documents
}

// NewObitDID creates new obit DID
func (sdk *Sdk) NewObitDID(serialNumber, manufacturer, partNumber string) (properties.DID, error) {
	var obitID properties.DID

	if sdk.debug {
		sdk.logger.Printf("NewObitDID(%q, %q, %q)", serialNumber, manufacturer, partNumber)
	}

	snProp, err := properties.NewStringProperty(
		"Making serialNumber",
		serialNumber,
		sdk.logger,
		sdk.debug,
	)

	if err != nil {
		return obitID, err
	}

	mnProp, err := properties.NewStringProperty(
		"Making manufacturer",
		manufacturer,
		sdk.logger,
		sdk.debug,
	)

	if err != nil {
		return obitID, err
	}

	pnProp, err := properties.NewStringProperty(
		"Making partNumber",
		partNumber,
		sdk.logger,
		sdk.debug,
	)

	if err != nil {
		return obitID, err
	}

	obitID, err = properties.NewDIDProperty(snProp, mnProp, pnProp, sdk.logger, sdk.debug)

	if err != nil {
		return obitID, err
	}

	return obitID, nil
}
