package properties

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/base58"
	"github.com/obada-foundation/sdk-go/hash"
	"log"
)

type ObitID struct {
	usn  string
	did  string
	hash hash.Hash
}

// NewObitIDProperty creates new ObitID from given arguments
func NewObitIDProperty(serialNumberHash StringProperty, manufacturer StringProperty, partNumber StringProperty, log *log.Logger, debug bool) (ObitID, error) {
	var id ObitID

	snhHash := serialNumberHash.GetHash()
	mh := manufacturer.GetHash()
	pnh := partNumber.GetHash()

	sum := serialNumberHash.GetHash().GetDec() + manufacturer.GetHash().GetDec() + partNumber.GetHash().GetDec()

	if debug {
		log.Printf(
			"NewObitIDProperty(%v, %v, %v) -> (%d + %d + %d) -> %d",
			serialNumberHash,
			manufacturer,
			partNumber,
			snhHash.GetDec(),
			mh.GetDec(),
			pnh.GetDec(),
			sum,
		)
	}

	h, err := hash.NewHash(fmt.Sprintf("%x", sum), log, debug)

	if err != nil {
		return id, fmt.Errorf("cannot create obit id: %w", err)
	}

	hashStr := h.GetHash()

	id.hash = h
	id.did = fmt.Sprintf("did:obada:%s", hashStr)
	id.usn = base58.Encode([]byte(hashStr))[:8]

	if debug {
		log.Printf("Hash: %s", h.GetHash())
		log.Printf("Did: %s", id.did)
		log.Printf("Usn: %s", id.usn)
	}

	return id, nil
}

// GetHash returns ObitID hash
func (id *ObitID) GetHash() hash.Hash {
	return id.hash
}

// GetDid returns obit DID
func (id *ObitID) GetDid() string {
	return id.did
}

// GetUsn returns the universal serial number
func (id *ObitID) GetUsn() string {
	return id.usn
}
