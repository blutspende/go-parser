package hl7v27

// VND - Purchasing Vendor
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/VND
type VND struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	VendorIdentifier EI `hl7:"POS=3"`
	VendorName string `hl7:"POS=4"`
	VendorCatalogNumber EI `hl7:"POS=5"`
	PrimaryVendorIndicator CNE `hl7:"POS=6"`
}

