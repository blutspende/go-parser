package hl7v271

// NSC - Application Status Change
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/NSC
type NSC struct {
	ApplicationChangeType CWE `hl7:"POS=2"`
	CurrentCpu string `hl7:"POS=3"`
	CurrentFileserver string `hl7:"POS=4"`
	CurrentApplication HD `hl7:"POS=5"`
	CurrentFacility HD `hl7:"POS=6"`
	NewCpu string `hl7:"POS=7"`
	NewFileserver string `hl7:"POS=8"`
	NewApplication HD `hl7:"POS=9"`
	NewFacility HD `hl7:"POS=10"`
}

