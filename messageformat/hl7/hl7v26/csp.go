package hl7v26

import "time"

// CSP - Clinical Study Phase
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/CSP
type CSP struct {
	StudyPhaseIdentifier CWE `hl7:"POS=2"`
	DateTimeStudyPhaseBegan time.Time `hl7:"POS=3"`
	DateTimeStudyPhaseEnded time.Time `hl7:"POS=4"`
	StudyPhaseEvaluability CWE `hl7:"POS=5"`
}

