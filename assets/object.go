package assets

import (
	"fmt"

	"github.com/obada-foundation/sdkgo/hash"
)

// DataObject temporary name might chnage in future
type DataObject struct {
	VersionID               string
	URL                     string
	HashEncryptedDataObject string
	HashUnencryptedObject   string
	HashUnencryptedMetadata hash.Hash
	HashEncryptedMetadata   hash.Hash
}

// Hash hashes concatenated DataObject properties
func (do DataObject) Hash() (hash.Hash, error) {
	dataObjectStr := fmt.Sprintf(
		"%s%s%s%s%s%s",
		do.VersionID,
		do.URL,
		do.HashEncryptedDataObject,
		do.HashUnencryptedObject,
		do.HashUnencryptedMetadata.GetHash(),
		do.HashEncryptedMetadata.GetHash(),
	)

	return hash.NewHash([]byte(dataObjectStr), nil)
}
