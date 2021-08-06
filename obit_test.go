package sdkgo

import (
	"bytes"
	"fmt"
	"github.com/obada-foundation/sdkgo/hash"
	"github.com/obada-foundation/sdkgo/properties"
	"log"
	"reflect"
	"testing"
)

func TestObit(t *testing.T) {
	type want struct {
		serialNumberHash string
		manufacturer string
		partNumber string
		obdDid string
		ownerDid string
		status string
		alternateIDS []string
		modifiedOn int64
		metadata map[string]string
		structuredData map[string]string
		checksum string
	}

	type args struct {
		dto ObitDto
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
					Status:   "STOLEN",
					AlternateIDS: []string{
						"1",
						"2",
					},
					ModifiedOn: 1628244835,
					Matadata: []properties.KV{
						{
							Key:   "type",
							Value: "phone",
						},
					},
					StructuredData: []properties.KV{
						{
							Key:   "color",
							Value: "red",
						},
					},
					Documents: []properties.Doc{},
				},
				parentChecksum: nil,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer: "Sony",
				partNumber: "PN123456",
				obdDid: "did:obada:obd:1234",
				ownerDid: "did:obada:owner:123456",
				status: "STOLEN",
				alternateIDS: []string{"1", "2"},
				modifiedOn: 1628244835,
				metadata: map[string]string{
					"type": "phone",
				},
				structuredData: map[string]string{
					"color": "red",
				},
				checksum: "fa953c53d8cb5221166ed99257e23c78ed1c6b1e64edae661ab85d03f2e1421b",
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
					Status:   "STOLEN",
					AlternateIDS: []string{
						"1",
						"2",
					},
					ModifiedOn: 1628244835,
					Matadata: []properties.KV{
						{
							Key:   "type",
							Value: "phone",
						},
					},
					StructuredData: []properties.KV{
						{
							Key:   "color",
							Value: "red",
						},
					},
					Documents: []properties.Doc{},
				},
				parentChecksum: &h,
			},
			want: want{
				serialNumberHash: "6dc5b8ae0ffe78e0276f08a935afac98cf2fce6bd6f00a0188e90a7d1462db03",
				manufacturer: "Sony",
				partNumber: "PN123456",
				obdDid: "did:obada:obd:1234",
				ownerDid: "did:obada:owner:123456",
				status: "STOLEN",
				alternateIDS: []string{"1", "2"},
				modifiedOn: 1628244835,
				metadata: map[string]string{
					"type": "phone",
				},
				structuredData: map[string]string{
					"color": "red",
				},
				checksum: "4946b22e3d8aa7285709bdcd87d61c5bae837acf9ee97cdbd3c733febf15920c",
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

		if tc.want.status != obit.GetStatus().GetValue() {
			t.Errorf("obit.GetStatus() = %q, want %q", obit.GetStatus().GetValue(), tc.want.status)
		}

		if !reflect.DeepEqual(tc.want.alternateIDS, obit.GetAlternateIDS().GetValue()) {
			t.Errorf("obit.GetAlternateIDS() = %q, want %q", obit.GetAlternateIDS().GetValue(), tc.want.alternateIDS)
		}

		if tc.want.modifiedOn != obit.GetModifiedOn().GetValue() {
			t.Errorf("obit.GetModifiedOn() = %q, want %q", obit.GetModifiedOn().GetValue(), tc.want.modifiedOn)
		}

		m := obit.GetMetadata()
		for _, rec := range m.GetAll() {
			mv, ok := tc.want.metadata[rec.GetKey().GetValue()]

			if !ok {
				t.Errorf("obit.GetMetadata() = %v and doesn't contain %v", m.GetAll(), tc.want.metadata)
			}

			if mv != rec.GetValue().GetValue() {
				t.Errorf("obit.GetMetadata() = %q and doesn't match %q", rec.GetValue().GetValue(), mv)
			}
		}

		sd := obit.GetStructuredData()
		for _, rec := range sd.GetAll() {
			sdv, ok := tc.want.structuredData[rec.GetKey().GetValue()]

			if !ok {
				t.Errorf("obit.GetStructuredData() = %v and doesn't contain %v", sd.GetAll(), tc.want.structuredData)
			}

			if sdv != rec.GetValue().GetValue() {
				t.Errorf("obit.GetStructuredData() = %q and doesn't match %q", rec.GetValue().GetValue(), sdv)
			}
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
