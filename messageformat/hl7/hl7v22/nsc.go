package hl7v22

// NSC - Status Change
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/NSC
type NSC struct {
	NetworkChangeType string `hl7:"POS=2"`
	CurrentCpu string `hl7:"POS=3"`
	CurrentFileserver string `hl7:"POS=4"`
	CurrentApplication string `hl7:"POS=5"`
	CurrentFacility string `hl7:"POS=6"`
	NewCpu string `hl7:"POS=7"`
	NewFileserver string `hl7:"POS=8"`
	NewApplication string `hl7:"POS=9"`
	NewFacility string `hl7:"POS=10"`
}

