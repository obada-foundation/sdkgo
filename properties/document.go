package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"log"
	"strconv"
)

// Document is mapping struct to Obit documents
type Document struct {
	name     StringProperty
	link     StringProperty
	checksum StringProperty
	hash     hash.Hash
}

// NewDocument creates a new Obit document
func NewDocument(name, link, docChecksum string, logger *log.Logger, debug bool) (Document, error) {
	var d Document

	if debug {
		logger.Printf("\n |New document| => NewDocument(%q, %q, %q)", name, link, docChecksum)
	}

	n, err := NewStringProperty("New document name", name, logger)

	if err != nil {
		return d, err
	}

	hl, err := NewStringProperty("New document hash link", link, logger)

	if err != nil {
		return d, err
	}

	dc, err := NewStringProperty("New document checksum", docChecksum, logger)

	if err != nil {
		return d, err
	}

	nh := n.GetHash()
	hlh := hl.GetHash()
	dch := dc.GetHash()
	docDec := nh.GetDec() + hlh.GetDec() + dch.GetDec()

	if debug {
		logger.Printf("(%d + %d + %d) -> %d", nh.GetDec(), hlh.GetDec(), dch.GetDec(), docDec)
	}

	h, err := hash.NewHash([]byte(strconv.FormatUint(docDec, 10)), logger)

	if err != nil {
		return d, err
	}

	d.name = n
	d.link = hl
	d.checksum = dc
	d.hash = h

	return d, nil
}

// GetName returns a document name
func (d *Document) GetName() StringProperty {
	return d.name
}

// GetLink returns a document hash link
func (d *Document) GetLink() StringProperty {
	return d.link
}

// GetChecksum returns a document data hash link
func (d *Document) GetChecksum() StringProperty {
	return d.checksum
}

// GetHash returns a document hash
func (d *Document) GetHash() hash.Hash {
	return d.hash
}

// Documents slice of documents
type Documents struct {
	documents []Document
	logger    *log.Logger
	debug     bool
}

// NewDocumentsCollection creates the collection of documents
func NewDocumentsCollection(logger *log.Logger, debug bool) Documents {
	return Documents{
		documents: make([]Document, 0),
		debug:     debug,
		logger:    logger,
	}
}

// AddDocument adds new document into Obit documents list
func (ds *Documents) AddDocument(d Document) {
	ds.documents = append(ds.documents, d)
}

// GetHash returns a hash of documents collection
func (ds *Documents) GetHash() (hash.Hash, error) {
	var docDec uint64

	description := "Making documents hash"

	if ds.debug {
		ds.logger.Printf("\n <|%s|> => %v", description, ds.documents)
	}

	if len(ds.documents) == 1 {
		return ds.documents[0].GetHash(), nil
	}

	for _, doc := range ds.documents {
		docDec += doc.GetHash().GetDec()
	}

	h, err := hash.NewHash([]byte(strconv.FormatUint(docDec, 10)), ds.logger)
	if err != nil {
		return h, fmt.Errorf("cannot hash %q: %w", docDec, err)
	}

	return h, nil
}

// GetAll returns all Obit documents
func (ds *Documents) GetAll() []Document {
	return ds.documents
}
