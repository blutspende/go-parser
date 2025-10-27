package hl7v23

// RQ1 - Requisition detail-1 segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RQ1
type RQ1 struct {
	AnticipatedPrice     string `hl7:"POS=2" json:"anticipatedPrice,omitempty"`
	ManufacturedId       CE     `hl7:"POS=3" json:"manufacturedId,omitempty"`
	ManufacturersCatalog string `hl7:"POS=4" json:"manufacturersCatalog,omitempty"`
	VendorID             CE     `hl7:"POS=5" json:"vendorID,omitempty"`
	VendorCatalog        string `hl7:"POS=6" json:"vendorCatalog,omitempty"`
	Taxable              string `hl7:"POS=7" json:"taxable,omitempty"`
	SubstituteAllowed    string `hl7:"POS=8" json:"substituteAllowed,omitempty"`
}
