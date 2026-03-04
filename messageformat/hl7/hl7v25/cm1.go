package hl7v25

// CM1 - Clinical Study Phase Master
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/CM1
type CM1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	StudyPhaseIdentifier CE `hl7:"POS=3"`
	DescriptionOfStudyPhase string `hl7:"POS=4"`
}

