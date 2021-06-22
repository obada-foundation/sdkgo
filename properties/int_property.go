package properties

import (
	"github.com/obada-foundation/sdk-go/hash"
	"log"
	"strconv"
)

type IntProperty struct {
	value int64
	hash  hash.Hash
}

func NewIntProperty(int int64, log *log.Logger, debug bool) (IntProperty, error) {
	var ip IntProperty

	if debug {
		log.Printf("\nNewIntProperty(%v)", int)
	}

	h, err := hash.NewHash(strconv.FormatInt(int, 10), log, debug)

	if err != nil {
		return ip, err
	}

	ip.hash = h
	ip.value = int

	return ip, nil
}

func (sp IntProperty) GetValue() int64 {
	return sp.value
}

func (sp IntProperty) GetHash() hash.Hash {
	return sp.hash
}

