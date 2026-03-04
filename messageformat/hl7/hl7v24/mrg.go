package hl7v24

// MRG - Merge patient information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MRG
type MRG struct {
	PriorPatientIdentifierList []CX `hl7:"POS=2"`
	PriorAlternatePatientID []CX `hl7:"POS=3"`
	PriorPatientAccountNumber CX `hl7:"POS=4"`
	PriorPatientID CX `hl7:"POS=5"`
	PriorVisitNumber CX `hl7:"POS=6"`
	PriorAlternateVisitID CX `hl7:"POS=7"`
	PriorPatientName []XPN `hl7:"POS=8"`
}

