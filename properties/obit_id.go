package properties

import (
	"fmt"
	"github.com/obada-protocol/sdk-go/base58"
	"github.com/obada-protocol/sdk-go/hash"
)

type ObitId struct {
	usn  string
	did  string
	hash hash.Hash
}

func NewObitId(snh SerialNumberHash, m Manufacturer, pn PartNumber) (ObitId, error) {
	var id ObitId

	snhHash := snh.GetHash()
	mh := m.GetHash()
	pnh := pn.GetHash()

	h, err := hash.NewHash(fmt.Sprintf("%x", snhHash.GetDec() + mh.GetDec() + pnh.GetDec()))

	if err != nil {
		return id, fmt.Errorf("cannot create obit id: %w", err)
	}

	hashStr := h.GetHash()

	id.hash = h
	id.did  = fmt.Sprintf("did:obada:%s", hashStr)
	id.usn  = base58.Encode([]byte(hashStr))[:8]

	return id, nil
}

func (id *ObitId) GetDid() string {
	return id.did
}

func (id *ObitId) GetUsn() string {
	return id.usn
}
