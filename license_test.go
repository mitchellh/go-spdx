package spdx

import (
	"testing"
)

func TestLicenseListLicense(t *testing.T) {
	ls := &LicenseList{}
	if v := ls.License("MIT"); v != nil {
		t.Fatalf("should be nil: %#v", v)
	}

	// should not panic on a nil license
	ls.Licenses = []*License{nil}
	if v := ls.License("MIT"); v != nil {
		t.Fatalf("should be nil: %#v", v)
	}

	// should find it
	ls.Licenses = []*License{
		&License{ID: "0BSD"},
		&License{ID: "MIT"},
	}
	if v := ls.License("MIT"); v == nil || v.ID != "MIT" {
		t.Fatalf("wrong: %#v", v)
	}
}
