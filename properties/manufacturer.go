package properties

import (
	"errors"
	"fmt"
	"github.com/obada-protocol/sdk-go/hash"
)

type Manufacturer struct {
	value string
	hash  hash.Hash
}

func NewManufacturer(manufacturer string) (Manufacturer, error) {
	var m Manufacturer

	if len(manufacturer) == 0 {
		return m, errors.New("manufacturer is required and cannot be empty")
	}

	h, err := hash.NewHash(manufacturer)

	if err != nil {
		return m, fmt.Errorf("cannot hash manufacturer %q: %w", manufacturer, err)
	}

	m.hash = h
	m.value = manufacturer

	return m, nil
}

func (m *Manufacturer) GetValue() string {
	return m.value
}

func (m *Manufacturer) GetHash() hash.Hash {
	return m.hash
}
