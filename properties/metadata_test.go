package properties_test

import (
	"log"
	"testing"

	"github.com/obada-foundation/sdkgo/properties"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	metadata   map[string]string
	json       string
	hash       string
	withLogger *log.Logger
	er         error
}

func TestNewMetadata(t *testing.T) {
	testCases := []testCase{
		{
			metadata:   make(map[string]string, 0),
			json:       `{}`,
			hash:       "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
			withLogger: nil,
			er:         nil,
		},
		{
			metadata: map[string]string{
				"": "",
			},
			withLogger: nil,
			er:         properties.ErrEmptyMetadataKey,
		},
		{
			metadata: map[string]string{
				"key": "value",
			},
			json:       `{"key":"value"}`,
			hash:       "eef78d5539cfc1997378d661d078cd784cc48175926823faf26a0366647b333a",
			withLogger: nil,
			er:         nil,
		},
	}

	for _, tc := range testCases {
		md, err := properties.NewMetadata(tc.metadata, nil)
		if tc.er != nil {
			require.ErrorIs(t, properties.ErrEmptyMetadataKey, err)
			continue
		}
		require.NoError(t, err)

		mdJSON, err := md.ToJSON()
		require.NoError(t, err)

		mdHash, err := md.Hash()
		require.NoError(t, err)

		assert.Equal(t, tc.json, string(mdJSON))
		assert.Equal(t, tc.hash, mdHash.GetHash())
	}

	md, err := properties.NewMetadata(make(map[string]string, 0), nil)
	require.NoError(t, err)

	err = md.AddRecord("keyX", "valueX")
	require.NoError(t, err)

	mdStr, err := md.String()
	require.NoError(t, err)

	mdHash, err := md.Hash()
	require.NoError(t, err)

	assert.Equal(t, `{"keyX":"valueX"}`, mdStr)
	assert.Equal(t, "6e169cde16d98a7eef2ecd92f0fa1a674d59ab15c89ea2570dc81ed70892488a", mdHash.GetHash())
}
