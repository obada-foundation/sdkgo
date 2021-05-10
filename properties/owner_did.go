package properties

import (
	"errors"
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
)

type OwnerDid struct {
	value string
	hash  hash.Hash
}

func NewOwnerDid(ownerDid string) (OwnerDid, error) {
	var od OwnerDid

	if len(ownerDid) == 0 {
		return od, errors.New("owner did is required and cannot be empty")
	}

	h, err := hash.NewHash(ownerDid)

	if err != nil {
		return od, fmt.Errorf("cannot hash owner did %q: %w", ownerDid, err)
	}

	od.hash = h
	od.value = ownerDid

	return od, nil
}

func (pn *OwnerDid) GetValue() string {
	return pn.value
}

func (pn *OwnerDid) GetHash() hash.Hash {
	return pn.hash
}
