package spdx

import (
	"reflect"
	"testing"
)

// Tests here depend on the spdx.org website functioning.

func TestList(t *testing.T) {
	var c Client
	list, err := c.List()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if list == nil {
		t.Fatal("list should not be nil")
	}

	var license *LicenseInfo
	for _, l := range list.Licenses {
		if l.ID == "MIT" {
			license = l
			break
		}
	}
	if license == nil {
		t.Fatal("MIT license not found")
	}
}

func TestList_exactURL(t *testing.T) {
	var c Client
	c.ListURL = "https://raw.githubusercontent.com/spdx/license-list-data/76b71ab7bc787f68e4847c6fe83a81ca3c82c0ef/json/licenses.json"

	list, err := c.List()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if list == nil {
		t.Fatal("list should not be nil")
	}

	var license *LicenseInfo
	for _, l := range list.Licenses {
		if l.ID == "MIT" {
			license = l
			break
		}
	}
	if license == nil {
		t.Fatal("MIT license not found")
	}

	expected := &LicenseInfo{
		ID:          "MIT",
		Name:        "MIT License",
		OSIApproved: true,
		SeeAlso:     []string{"http://www.opensource.org/licenses/MIT"},
	}
	if !reflect.DeepEqual(license, expected) {
		t.Fatalf("bad: %#v", license)
	}
}

func TestLicense(t *testing.T) {
	var c Client
	license, err := c.License("MIT")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if license == nil {
		t.Fatal("license should not be nil")
	}
	if license.Text == "" {
		t.Fatal("should have license text")
	}
}
