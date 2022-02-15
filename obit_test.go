package sdkgo

import (
	"bytes"
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"log"
	"testing"
)

func TestObit(t *testing.T) {
	type want struct {
		serialNumberHash string
		manufacturer     string
		partNumber       string
		obdDid           string
		ownerDid         string
		checksum         string
	}

	type args struct {
		dto            ObitDto
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
				dto: ObitDto{
					ObitIDDto: ObitIDDto{
						// The value of SerialNumberHash is sha256(SN123456)
						SerialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
						Manufacturer:     "Sony",
						PartNumber:       "PN123456",
					},
					ObdDid:   "did:obada:obd:1234",
					OwnerDid: "did:obada:owner:123456",
				},
				parentChecksum: nil,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer:     "Sony",
				partNumber:       "PN123456",
				obdDid:           "did:obada:obd:1234",
				ownerDid:         "did:obada:owner:123456",
				checksum:         "af9d7cf448738b105f9ba35230bb5fd6c0e8c295b2be78bd5f7c1278870eb416",
			},
		},
		{
			args: args{
				dto: ObitDto{
					ObitIDDto: ObitIDDto{
						// The value of SerialNumberHash is sha256(SN123456)
						SerialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
						Manufacturer:     "Sony",
						PartNumber:       "PN123456",
					},
					ObdDid:   "did:obada:obd:1234",
					OwnerDid: "did:obada:owner:123456",
				},
				parentChecksum: &h,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer:     "Sony",
				partNumber:       "PN123456",
				obdDid:           "did:obada:obd:1234",
				ownerDid:         "did:obada:owner:123456",
				checksum:         "20e48ddaf574116ca9ffd0d0c6538dae10c276c124d6e89408ac634a25235616",
			},
		},
	}

	var str bytes.Buffer

	logger := log.New(&str, "TESTING SDK :: ", 0)

	sdk, err := NewSdk(logger, true)

	if err != nil {
		fmt.Println(str.String())
		t.Fatalf("Cannot initialize OBADA SDK. %s", err)
	}

	for _, tc := range testCases {
		obit, err := sdk.NewObit(tc.args.dto)

		if err != nil {
			fmt.Println(str.String())
			t.Fatalf("Cannot create obit from given data: %v. Reason: %s", tc.args.dto, err.Error())
		}

		if tc.want.serialNumberHash != obit.GetSerialNumberHash().GetValue() {
			t.Errorf("obit.GetSerialNumberHash() = %q, want %q", obit.GetSerialNumberHash().GetValue(), tc.want.serialNumberHash)
		}

		if tc.want.manufacturer != obit.GetManufacturer().GetValue() {
			t.Errorf("obit.GetManufacturer() = %q, want %q", obit.GetManufacturer().GetValue(), tc.want.manufacturer)
		}

		if tc.want.partNumber != obit.GetPartNumber().GetValue() {
			t.Errorf("obit.GetPartNumber() = %q, want %q", obit.GetPartNumber().GetValue(), tc.want.partNumber)
		}

		if tc.want.partNumber != obit.GetPartNumber().GetValue() {
			t.Errorf("obit.GetPartNumber() = %q, want %q", obit.GetPartNumber().GetValue(), tc.want.partNumber)
		}

		if tc.want.obdDid != obit.GetObdDID().GetValue() {
			t.Errorf("obit.GetObdDID() = %q, want %q", obit.GetObdDID().GetValue(), tc.want.obdDid)
		}

		if tc.want.ownerDid != obit.GetOwnerDID().GetValue() {
			t.Errorf("obit.GetOwnerDID() = %q, want %q", obit.GetOwnerDID().GetValue(), tc.want.ownerDid)
		}

		checksum, err := obit.GetChecksum(tc.args.parentChecksum)

		if err != nil {
			fmt.Println(str.String())
			t.Fatalf(err.Error())
		}

		if tc.want.checksum != checksum.GetHash() {
			t.Errorf("obit.Checksum(%v) = %q, want %q", tc.args.parentChecksum, tc.want.checksum, checksum.GetHash())
		}
	}
}
