package encryption_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/obada-foundation/sdkgo/encryption"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryption(t *testing.T) {
	privKey := secp256k1.GenPrivKey()

	plainData := []byte("Test encryption")

	encData, err := encryption.Encrypt(privKey.PubKey(), plainData)
	require.NoError(t, err)

	decData, err := encryption.Decrypt(privKey, encData)
	require.NoError(t, err)

	assert.Equal(t, plainData, decData)
}
