package properties

import (
	"github.com/obada-foundation/sdkgo/hash"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Doc struct {
	Name     string
	HashLink string
}

// HashLinkProperty represent a hash of the hash link source
type HashLinkProperty struct {
	hashLink string
	hash     hash.Hash
}

// NewHashLinkProperty creates a new property and compute a hash of file under hash link
func NewHashLinkProperty(description, hashLink string, logger *log.Logger, debug bool) (HashLinkProperty, error) {
	var hlp HashLinkProperty

	if debug {
		logger.Printf("\n <|%s|> => NewStringProperty(%v)", description, hashLink)
	}

	resp, err := http.Get(hashLink)
	if err != nil {
		return hlp, err
	}
	defer resp.Body.Close()

	fileBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return hlp, err
	}

	h, err := hash.NewHash(fileBytes, logger, debug)

	hlp.hashLink = hashLink
	hlp.hash = h

	return hlp, nil
}

func (hlp HashLinkProperty) GetHash() hash.Hash {
	return hlp.hash
}

func (hlp HashLinkProperty) GetHashLink() string {
	return hlp.hashLink
}

// Document Obit document
type Document struct {
	name     StringProperty
	hashLink HashLinkProperty
	hash     hash.Hash
}

// NewDocument creates a new Obit document
func NewDocument(description, name, hashLink string, logger *log.Logger, debug bool) (Document, error) {
	var d Document

	if debug {
		logger.Printf("\n |%s| => NewDocument(%q, %q)", description, name, hashLink)
	}

	n, err := NewStringProperty("New document name", name, logger, debug)

	if err != nil {
		return d, err
	}

	hl, err := NewHashLinkProperty("New document hash link", hashLink, logger, debug)

	if err != nil {
		return d, err
	}

	nh := n.GetHash()
	hlh := hl.GetHash()
	docDec := nh.GetDec() + hlh.GetDec()

	if debug {
		logger.Printf("(%d + %d) -> %d", nh.GetDec(), hlh.GetDec(), docDec)
	}

	h, err := hash.NewHash([]byte(strconv.FormatUint(docDec, 10)), logger, debug)

	if err != nil {
		return d, err
	}

	d.name = n
	d.hashLink = hl
	d.hash = h

	return d, nil
}

// GetName returns a document name
func (d *Document) GetName() StringProperty {
	return d.name
}

// GetHashLink returns a document hash link
func (d *Document) GetHashLink() HashLinkProperty {
	return d.hashLink
}

// GetHash returns a document hash
func (d *Document) GetHash() hash.Hash {
	return d.hash
}

// DocumentCollection
type DocumentCollection struct {
	documents []Document
	hash      hash.Hash
}

// NewDocumentCollection creates the collection of documents
func NewDocumentCollection(description string, docs []Doc, logger *log.Logger, debug bool) (DocumentCollection, error) {
	var dc DocumentCollection

	if debug {
		logger.Printf("\n <|%s|> => NewDocumentCollection(%v)", description, docs)
	}

	return dc, nil
}

func (dc DocumentCollection) GetHash() hash.Hash {
	return dc.hash
}
