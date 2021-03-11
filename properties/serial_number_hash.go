package properties

import (
	"errors"
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
)

type SerialNumberHash struct {
	value string
	hash  hash.Hash
}

func NewSerialNumberHash(serialNumberHash string) (SerialNumberHash, error) {
	var snh SerialNumberHash

	if len(serialNumberHash) == 0 {
		return snh, errors.New("serial number hash must be a valid SHA256 hash")
	}

	h, err := hash.NewHash(serialNumberHash)

	if err != nil {
		return snh, fmt.Errorf("cannot hash serial number hash %q: %w", serialNumberHash, err)
	}

	snh.hash = h
	snh.value = serialNumberHash

	return snh, nil
}

func (snh *SerialNumberHash) GetValue() string {
	return snh.value
}

func (snh *SerialNumberHash) GetHash() hash.Hash {
	return snh.hash
}
