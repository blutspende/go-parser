package hl7v231

// RQ1 - Requisition detail-1 segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/RQ1
type RQ1 struct {
	AnticipatedPrice string `hl7:"POS=2"`
	ManufacturerID CE `hl7:"POS=3"`
	ManufacturersCatalog string `hl7:"POS=4"`
	VendorID CE `hl7:"POS=5"`
	VendorCatalog string `hl7:"POS=6"`
	Taxable string `hl7:"POS=7"`
	SubstituteAllowed string `hl7:"POS=8"`
}

