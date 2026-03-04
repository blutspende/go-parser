package hl7v27

// CM1 - Clinical Study Phase Master
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/CM1
type CM1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	StudyPhaseIdentifier CWE `hl7:"POS=3"`
	DescriptionOfStudyPhase string `hl7:"POS=4"`
}

