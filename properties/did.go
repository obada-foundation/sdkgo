package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/base58"
	"github.com/obada-foundation/sdkgo/hash"
	"log"
)

// DID represent obit decentralized identifier
type DID struct {
	usn     string
	fullUsn string
	did     string
	hash    hash.Hash
}

// NewDIDProperty creates new DID from given arguments
func NewDIDProperty(serialNumberHash, manufacturer, partNumber StringProperty, logger *log.Logger, debug bool) (DID, error) {
	var did DID

	snhHash := serialNumberHash.GetHash()
	mh := manufacturer.GetHash()
	pnh := partNumber.GetHash()

	sum := serialNumberHash.GetHash().GetDec() + manufacturer.GetHash().GetDec() + partNumber.GetHash().GetDec()

	if debug {
		logger.Printf(
			"\n<|%s|> => NewDIDProperty(%v, %v, %v) -> (%d + %d + %d) -> %d",
			"Making DID",
			serialNumberHash,
			manufacturer,
			partNumber,
			snhHash.GetDec(),
			mh.GetDec(),
			pnh.GetDec(),
			sum,
		)
	}

	h, err := hash.NewHash([]byte(fmt.Sprintf("%x", sum)), logger, debug)

	if err != nil {
		return did, fmt.Errorf("cannot create DID: %w", err)
	}

	hashStr := h.GetHash()

	did.hash = h
	did.did = "did:obada:" + hashStr

	fullUsn := base58.Encode([]byte(hashStr))

	did.usn = fullUsn[:12]
	did.fullUsn = fullUsn

	if debug {
		logger.Printf("Hash: %s", h.GetHash())
		logger.Printf("Did: %s", did.did)
		logger.Printf("Usn: %s", did.usn)
		logger.Printf("Full Usn: %s", did.fullUsn)
	}

	return did, nil
}

// GetHash returns DID hash
func (id *DID) GetHash() hash.Hash {
	return id.hash
}

// GetDid returns obit DID
func (id *DID) GetDid() string {
	return id.did
}

// GetUsn returns short universal serial number
func (id *DID) GetUsn() string {
	return id.usn
}

// GetFullUsn returns full universal serial number
func (id *DID) GetFullUsn() string {
	return id.fullUsn
}
