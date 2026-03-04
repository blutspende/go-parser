package hl7v251

// FAC - Facility
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/FAC
type FAC struct {
	FacilityIDFac EI `hl7:"POS=2"`
	FacilityType string `hl7:"POS=3"`
	FacilityAddress []XAD `hl7:"POS=4"`
	FacilityTelecommunication XTN `hl7:"POS=5"`
	ContactPerson []XCN `hl7:"POS=6"`
	ContactTitle []string `hl7:"POS=7"`
	ContactAddress []XAD `hl7:"POS=8"`
	ContactTelecommunication []XTN `hl7:"POS=9"`
	SignatureAuthority []XCN `hl7:"POS=10"`
	SignatureAuthorityTitle string `hl7:"POS=11"`
	SignatureAuthorityAddress []XAD `hl7:"POS=12"`
	SignatureAuthorityTelecommunication XTN `hl7:"POS=13"`
}

