package properties

import (
	"fmt"
	"testing"

	"github.com/obada-foundation/sdkgo/tests"
)

func TestNewDocumentProperty(t *testing.T) {
	type given struct {
		name     string
		link     string
		checksum string
	}

	type want struct {
		hash string
	}

	type testCase struct {
		given given
		want  want
	}

	testCases := []testCase{
		{
			given: given{
				name:     "wipe report",
				link:     "https://google.com",
				checksum: "",
			},
			want: want{
				hash: "c9893dae3d4c7440d273fc21ae36b3a3b61c69c2687c5004e064c9327d6f735e",
			},
		},
	}

	for _, tc := range testCases {
		logger, buff := tests.CreateSdkTestLogger()

		doc, err := NewDocument(tc.given.name, tc.given.link, tc.given.checksum, logger, true)
		if err != nil {
			fmt.Println(buff.String())
			t.Fatal(err.Error())
		}

		if tc.given.name != doc.GetName().GetValue() {
			fmt.Println(buff.String())
			t.Errorf("doc.GetName() = %q, want %q", doc.GetName().GetValue(), tc.given.name)
		}

		if tc.given.link != doc.GetLink().GetValue() {
			fmt.Println(buff.String())
			t.Errorf("doc.GetLink() = %q, want %q", doc.GetLink().GetValue(), tc.given.link)
		}

		if tc.given.checksum != doc.GetChecksum().GetValue() {
			fmt.Println(buff.String())
			t.Errorf("doc.GetChecksum() = %q, want %q", doc.GetChecksum().GetValue(), tc.given.checksum)
		}

		if tc.given.checksum != doc.GetChecksum().GetValue() {
			fmt.Println(buff.String())
			t.Errorf("doc.GetHash() = %q, want %q", doc.GetHash().GetHash(), tc.given.checksum)
		}

		hash := doc.GetHash()
		if tc.want.hash != hash.GetHash() {
			fmt.Println(buff.String())
			t.Errorf("doc.GetHash() = %q, want %q", hash.GetHash(), tc.want.hash)
		}
	}
}

func TestNewDocumentsProperty(t *testing.T) {
	type given struct {
		name     string
		link     string
		checksum string
	}

	type want struct {
		docLength int
		hash      string
	}

	type testCase struct {
		given []given
		want  want
	}

	testCases := []testCase{
		{
			given: []given{
				{
					name:     "wipe report",
					link:     "https://google.com",
					checksum: "",
				},
			},
			want: want{
				docLength: 1,
				hash:      "c9893dae3d4c7440d273fc21ae36b3a3b61c69c2687c5004e064c9327d6f735e",
			},
		},
		{
			given: []given{
				{
					name:     "wipe report",
					link:     "https://google.com",
					checksum: "",
				},
				{
					name:     "wipe report2",
					link:     "https://google.com",
					checksum: "",
				},
			},
			want: want{
				docLength: 2,
				hash:      "6dd558d089f76c9fd6e4b75ad5313127c7479dca85c7749582123bab279b74e7",
			},
		},
	}

	for _, tc := range testCases {
		logger, buff := tests.CreateSdkTestLogger()

		docs := NewDocumentsCollection(logger, true)

		for _, doc := range tc.given {
			d, err := NewDocument(doc.name, doc.link, doc.checksum, logger, true)
			if err != nil {
				fmt.Println(buff.String())
				t.Fatal(err.Error())
			}

			docs.AddDocument(d)
		}

		hash, err := docs.GetHash()
		if err != nil {
			fmt.Println(buff.String())
			t.Fatal(err.Error())
		}

		if hash.GetHash() != tc.want.hash {
			t.Errorf("docs.GetHash() = %q, want %q", hash.GetHash(), tc.want.hash)
		}

		if len(docs.GetAll()) != tc.want.docLength {
			t.Errorf("docs.GetAll() = %d, want %d", len(docs.GetAll()), tc.want.docLength)
		}
	}
}
