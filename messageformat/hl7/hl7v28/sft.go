package hl7v28

import "time"

// SFT - Software Segment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/SFT
type SFT struct {
	SoftwareVendorOrganization XON `hl7:"POS=2"`
	SoftwareCertifiedVersionOrReleaseNumber string `hl7:"POS=3"`
	SoftwareProductName string `hl7:"POS=4"`
	SoftwareBinaryID string `hl7:"POS=5"`
	SoftwareProductInformation string `hl7:"POS=6"`
	SoftwareInstallDate time.Time `hl7:"POS=7"`
}

