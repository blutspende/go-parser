package hl7v231

// CM1 - Clinical study phase master segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/CM1
type CM1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	StudyPhaseIdentifier CE `hl7:"POS=3"`
	DescriptionOfStudyPhase string `hl7:"POS=4"`
}

