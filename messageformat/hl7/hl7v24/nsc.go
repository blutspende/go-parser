package hl7v24

// NSC - Application status change
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NSC
type NSC struct {
	ApplicationChangeType string `hl7:"POS=2"`
	CurrentCpu string `hl7:"POS=3"`
	CurrentFileserver string `hl7:"POS=4"`
	CurrentApplication HD `hl7:"POS=5"`
	CurrentFacility HD `hl7:"POS=6"`
	NewCpu string `hl7:"POS=7"`
	NewFileserver string `hl7:"POS=8"`
	NewApplication HD `hl7:"POS=9"`
	NewFacility HD `hl7:"POS=10"`
}

