package properties

import (
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
)

type ObdDid struct {
	value string
	hash  hash.Hash
}

func NewObdDid(ownerDid string) (ObdDid, error) {
	var obd ObdDid

	h, err := hash.NewHash(ownerDid)

	if err != nil {
		return obd, fmt.Errorf("cannot hash owner did %q: %w", ownerDid, err)
	}

	obd.hash = h
	obd.value = ownerDid

	return obd, nil
}

func (pn *ObdDid) GetValue() string {
	return pn.value
}

func (pn *ObdDid) GetHash() hash.Hash {
	return pn.hash
}
