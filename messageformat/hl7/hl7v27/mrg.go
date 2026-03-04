package hl7v27

// MRG - Merge Patient Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/MRG
type MRG struct {
	PriorPatientIdentifierList []CX `hl7:"POS=2"`
	PriorPatientAccountNumber CX `hl7:"POS=4"`
	PriorVisitNumber CX `hl7:"POS=6"`
	PriorAlternateVisitID CX `hl7:"POS=7"`
	PriorPatientName []XPN `hl7:"POS=8"`
}

