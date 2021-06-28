package properties

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"log"
	"strconv"
)

// SliceStrProperty represent slice of strings property and their hash
type SliceStrProperty struct {
	value []string
	hash  hash.Hash
}

// NewSliceStrProperty creates a new obit property from slice strings
func NewSliceStrProperty(value []string, logger *log.Logger, debug bool) (SliceStrProperty, error) {
	var stp SliceStrProperty

	if debug {
		logger.Printf("\nNewSliceStrProperty(%v)", value)
	}

	var dec uint64

	for _, str := range value {
		h, err := hash.NewHash(str, logger, debug)

		if err != nil {
			return stp, fmt.Errorf("cannot hash %q: %w", value, err)
		}

		dec += h.GetDec()
	}

	h, err := hash.NewHash(strconv.FormatUint(dec, 10), logger, debug)

	if err != nil {
		return stp, fmt.Errorf("cannot hash %q: %w", value, err)
	}

	stp.hash = h
	stp.value = value

	return stp, nil
}

// GetValue returns a slice of strings
func (stp SliceStrProperty) GetValue() []string {
	return stp.value
}

// GetHash returns a slice of strings hash
func (stp SliceStrProperty) GetHash() hash.Hash {
	return stp.hash
}
