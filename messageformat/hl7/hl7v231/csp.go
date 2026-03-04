package hl7v231

import "time"

// CSP - Clinical study phase segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/CSP
type CSP struct {
	StudyPhaseIdentifier CE `hl7:"POS=2"`
	DateTimeStudyPhaseBegan time.Time `hl7:"POS=3"`
	DateTimeStudyPhaseEnded time.Time `hl7:"POS=4"`
	StudyPhaseEvaluability CE `hl7:"POS=5"`
}

