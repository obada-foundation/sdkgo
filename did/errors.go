package did

import (
	"errors"
)

var (
	ErrNotSupportedDIDMethod = errors.New("given DID method is not supported")
)
