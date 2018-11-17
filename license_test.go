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
	ls.Licenses = []*LicenseInfo{nil}
	if v := ls.License("MIT"); v != nil {
		t.Fatalf("should be nil: %#v", v)
	}

	// should find it
	ls.Licenses = []*LicenseInfo{
		&LicenseInfo{ID: "0BSD"},
		&LicenseInfo{ID: "MIT"},
	}
	if v := ls.License("MIT"); v == nil || v.ID != "MIT" {
		t.Fatalf("wrong: %#v", v)
	}
}
