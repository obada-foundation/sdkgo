package properties

import (
	"errors"
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
)

type PartNumber struct {
	value string
	hash  hash.Hash
}

func NewPartNumber(partNumber string) (PartNumber, error) {
	var pn PartNumber

	if len(partNumber) == 0 {
		return pn, errors.New("part number is required and cannot be empty")
	}

	h, err := hash.NewHash(partNumber)

	if err != nil {
		return pn, fmt.Errorf("cannot hash manufacturer %q: %w", partNumber, err)
	}

	pn.hash = h
	pn.value = partNumber

	return pn, nil
}

func (pn *PartNumber) GetValue() string {
	return pn.value
}

func (pn *PartNumber) GetHash() hash.Hash {
	return pn.hash
}
