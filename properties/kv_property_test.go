package properties

import "testing"

func TestNewRecord(t *testing.T) {
	r, err := NewRecord("color", "red")

	if err != nil {
		t.Fatalf(err.Error())
	}

	k := r.GetKey()

	if k.GetValue() != "color" {
		t.Fatalf("Expected to to get value %q but received %q", "color", k.GetValue())
	}

	v := r.GetValue()

	if v.GetValue() != "red" {
		t.Fatalf("Expected to to get value %q but received %q", "color", v.GetValue())
	}

	h := r.GetHash()

	if h.GetHash() != "709082f0e18e52ae4d519c1753f4ed40d28d6ddae5c6157e1718fa56f1a21e4e" {
		t.Fatalf("Expected to to get hash %q but received %q", "709082f0e18e52ae4d519c1753f4ed40d28d6ddae5c6157e1718fa56f1a21e4e", h.GetHash())
	}
}
