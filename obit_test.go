package sdkgo

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/obada-foundation/sdkgo/hash"
	"github.com/obada-foundation/sdkgo/properties"
)

func TestObit(t *testing.T) {
	type doc struct {
		link     string
		name     string
		checksum string
	}

	type want struct {
		did      string
		checksum string
	}

	type args struct {
		serialNumber   string
		manufacturer   string
		partNumber     string
		docs           []doc
		parentChecksum *hash.Hash
	}

	type testCase struct {
		args args
		want want
	}

	h, err := hash.NewHash([]byte("some data"), nil, false)

	if err != nil {
		t.Fatal(err)
	}

	testCases := []testCase{
		{
			args: args{
				serialNumber: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer: "Sony",
				partNumber:   "PN123456",
				docs: []doc{
					{
						name:     "ubuntu iso",
						link:     "http://releases.ubuntu.com/22.04/ubuntu-22.04-desktop-amd64.iso",
						checksum: "b85286d9855f549ed9895763519f6a295a7698fb9c5c5345811b3eefadfb6f07",
					},
				},
				parentChecksum: nil,
			},
			want: want{
				did:      "did:obada:c7e24b72c739a5619d24dd8e87e1ea4a829a508166203661ea8eda7a8a0b5978",
				checksum: "07f1f7d0f85d39e8ed7a9b4832db52118fcdbb9ca63d37c5c2e360bd20402225",
			},
		},
		{
			args: args{
				serialNumber:   "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer:   "Sony",
				partNumber:     "PN123456",
				docs:           []doc{},
				parentChecksum: &h,
			},
			want: want{
				did:      "did:obada:c7e24b72c739a5619d24dd8e87e1ea4a829a508166203661ea8eda7a8a0b5978",
				checksum: "a8e83c684a166fcf80e88c2fc80206f35ed93fa752890164657145829e065734",
			},
		},
	}

	var str bytes.Buffer

	logger := log.New(&str, "TESTING SDK :: ", 0)

	sdk := NewSdk(logger, true)

	for _, tc := range testCases {
		DID, err := sdk.NewObitDID(tc.args.serialNumber, tc.args.manufacturer, tc.args.partNumber)
		if err != nil {
			t.Fatalf("Cannot create DID from the given data: (%s, %s, %s). Reason: %s", tc.args.serialNumber, tc.args.manufacturer, tc.args.partNumber, err.Error())
		}

		var docs []properties.Document

		for _, d := range tc.args.docs {
			doc, er := sdk.NewDocument(d.name, d.link, d.checksum)
			if er != nil {
				t.Fatalf("Cannot create Document from the given data: (%v). Reason: %s", d, er.Error())
			}

			docs = append(docs, doc)
		}

		documents := sdk.NewDocuments(docs)

		obit, err := sdk.NewObit(DID, documents)
		if err != nil {
			fmt.Println(str.String())
			t.Fatalf("Cannot create obit from given data: %v. Reason: %s", DID, err.Error())
		}

		checksum, err := obit.GetChecksum(tc.args.parentChecksum)

		if err != nil {
			fmt.Println(str.String())
			t.Fatalf(err.Error())
		}

		if tc.want.did != obit.obitDID.GetDid() {
			fmt.Println(str.String())
			t.Fatalf("DID doesn't match. Given: %s, want %s", obit.obitDID.GetDid(), tc.want.did)
		}

		if tc.want.checksum != checksum.GetHash() {
			fmt.Println(str.String())
			t.Fatalf("obit.Checksum(%v) = %q, want %q", tc.args.parentChecksum, checksum.GetHash(), tc.want.checksum)
		}

		fmt.Println(str.String())

	}
}
