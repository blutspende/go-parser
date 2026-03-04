package hl7v27

// QRI - Query Response Instance
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/QRI
type QRI struct {
	CandidateConfidence *int `hl7:"POS=2"`
	MatchReasonCode []CWE `hl7:"POS=3"`
	AlgorithmDescriptor CWE `hl7:"POS=4"`
}

