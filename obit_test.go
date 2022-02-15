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
		trustAnchorToken string
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
					TrustAnchorToken: "",
				},
				parentChecksum: nil,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer:     "Sony",
				partNumber:       "PN123456",
				trustAnchorToken: "",
				checksum:         "2cd29b69ff050bbd186b5cd27d6435731e66d6cd8c77d5922e472625c61538f8",
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
					TrustAnchorToken: "",
				},
				parentChecksum: &h,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer:     "Sony",
				partNumber:       "PN123456",
				trustAnchorToken: "",
				checksum:         "565ade1e624fdccb0cdbd7e88d9cd2655533fecdd764378757638eaae3f1ece6",
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

		if tc.want.trustAnchorToken != obit.GetTrustAnchorToken().GetValue() {
			t.Errorf("obit.GetTrustAnchorToken() = %q, want %q", obit.GetTrustAnchorToken().GetValue(), tc.want.trustAnchorToken)
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
