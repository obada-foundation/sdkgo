package sdkgo

import (
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"github.com/obada-foundation/sdkgo/properties"
)

// NewObit creates new obit
func (sdk *Sdk) NewObit(did properties.DID, docs properties.Documents) (Obit, error) {
	var o Obit
	o.debug = sdk.debug
	o.logger = sdk.logger

	if sdk.debug {
		sdk.logger.Printf("<|%s|> => NewObit(%v, %v)", "<|Creating new Obit|>", did, docs)
	}

	o.obitDID = did
	o.documents = docs

	return o, nil
}

// GetDID returns obit DID
func (o *Obit) GetDID() properties.DID {
	return o.obitDID
}

// GetDocuments returns Obit documents
func (o *Obit) GetDocuments() properties.Documents {
	return o.documents
}

// GetChecksum returns Obit checksum
func (o *Obit) GetChecksum(parentChecksum *hash.Hash) (hash.Hash, error) {
	var checksum hash.Hash

	if o.debug {
		o.logger.Println("\n\n<|Obit checksum calculation|>")
	}

	documentsHash, err := o.documents.GetHash()
	if err != nil {
		return checksum, err
	}

	sum := o.obitDID.GetHash().GetDec() +
		documentsHash.GetDec()

	if o.debug {
		o.logger.Printf(
			"(%d + %d) -> %d -> Dec2Hex(%d) -> %s",
			o.obitDID.GetHash().GetDec(),
			documentsHash.GetDec(),
			sum,
			sum,
			fmt.Sprintf("%x", sum),
		)
	}

	if parentChecksum != nil {
		prhDec := parentChecksum.GetDec()

		if o.debug {
			o.logger.Printf("(%d + %d) -> %d", sum, prhDec, sum+prhDec)
		}

		sum += prhDec
	}

	checksum, err = hash.NewHash([]byte(fmt.Sprintf("%x", sum)), o.logger, o.debug)
	if err != nil {
		return checksum, err
	}

	if o.debug {
		o.logger.Printf("Checksum(%v) -> %q", o, checksum.GetHash())
	}

	return checksum, nil
}
