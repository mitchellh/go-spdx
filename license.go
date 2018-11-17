package spdx

// LicenseList is the structure for a list of licenses provided by the SPDX API
type LicenseList struct {
	// Version is the raw version string of the license list.
	Version string `json:"licenseListVersion"`

	// Licenses is the list of known licenses.
	Licenses []*License `json:"licenses"`
}

// License is a single software license.
//
// Basic descriptions are documented in the fields below. For a full
// description of the fields, see the official SPDX specification here:
// https://github.com/spdx/license-list-data/blob/master/accessingLicenses.md
type License struct {
	ID          string   `json:"licenseId"`
	Name        string   `json:"name"`
	Text        string   `json:"licenseText"`
	Deprecated  bool     `json:"isDeprecatedLicenseId"`
	OSIApproved bool     `json:"isOsiApproved"`
	SeeAlso     []string `json:"seeAlso"`
}
