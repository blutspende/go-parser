package hl7v23

// RQ1 - Requisition detail-1 segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RQ1
type RQ1 struct {
	AnticipatedPrice     string `hl7:"1" json:"anticipatedPrice,omitempty"`
	ManufacturedId       CE     `hl7:"2" json:"manufacturedId,omitempty"`
	ManufacturersCatalog string `hl7:"3" json:"manufacturersCatalog,omitempty"`
	VendorID             CE     `hl7:"4" json:"vendorID,omitempty"`
	VendorCatalog        string `hl7:"5" json:"vendorCatalog,omitempty"`
	Taxable              string `hl7:"6" json:"taxable,omitempty"`
	SubstituteAllowed    string `hl7:"7" json:"substituteAllowed,omitempty"`
}
