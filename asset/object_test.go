package asset_test

import (
	"testing"

	"github.com/obada-foundation/sdkgo/asset"
	"github.com/obada-foundation/sdkgo/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	asset asset.DataArrayVersions
	err   string
}

func TestRootHash(t *testing.T) {

	tcs := []testCase{
		{
			asset: asset.DataArrayVersions{},
			err:   "cannot compute root hash because of no snapshots",
		},
		{
			asset: asset.DataArrayVersions{
				1: {},
			},
			err: "cannot compute root hash because of empty data array",
		},
		{
			asset: asset.DataArrayVersions{
				1:   {},
				2:   {},
				123: {},
			},
			err: "snapshot versions are not incremental: [1 2 123]",
		},
		{
			asset: asset.DataArrayVersions{
				1: {
					Objects: []asset.Object{
						{},
					},
				},
			},
			err: "empty metadata",
		},
		{
			asset: asset.DataArrayVersions{
				1: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.PhysicalAssetIdentifiers),
							},
						},
						{
							Metadata: map[string]string{
								"type": string(asset.MainImage),
							},
						},
					},
				},
			},
			err: "missing dataobject hash",
		},
		{
			asset: asset.DataArrayVersions{
				1: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.PhysicalAssetIdentifiers),
							},
							HashUnencryptedObject: "QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
							URL:                   "https://ipfs.io/ipfs/QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
						},
					},
				},
			},
			err: "",
		},
		{
			asset: asset.DataArrayVersions{
				1: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.PhysicalAssetIdentifiers),
							},
							HashUnencryptedObject: "QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
							URL:                   "https://ipfs.io/ipfs/QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
						},
					},
				},
				2: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.PhysicalAssetIdentifiers),
							},
							HashUnencryptedObject: "QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
							URL:                   "https://ipfs.io/ipfs/QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
						},
					},
				},
			},
			err: "data in a version 2 are the same as in previous version",
		},
		{
			asset: asset.DataArrayVersions{
				1: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.MainImage),
							},
							HashUnencryptedObject: "QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
							URL:                   "https://ipfs.io/ipfs/QmQqzMTavQgT4f4T5v6PWBp7XNKtoPmC9jvn12WPT3gkSE",
						},
					},
				},
				2: {
					Objects: []asset.Object{
						{
							Metadata: map[string]string{
								"type": string(asset.MainImage),
							},
							HashUnencryptedObject: "QmUyARmq5RUJk5zt7KUeaMLYB8SQbKHp3Gdqy5WSxRtPNa",
							URL:                   "https://ipfs.pixura.io/ipfs/QmUyARmq5RUJk5zt7KUeaMLYB8SQbKHp3Gdqy5WSxRtPNa/SeaofRoses.jpg",
						},
					},
				},
			},
			err: "",
		},
	}

	logger, logOutput := testutil.TestLogger("ROOT HASH :: ")
	defer t.Logf("Compute log: \n%v", logOutput)

	for _, tc := range tcs {
		_, err := asset.RootHash(tc.asset, logger)
		if tc.err != "" {
			assert.Equal(t, tc.err, err.Error())
			continue
		}

		require.NoError(t, err)
	}
}
