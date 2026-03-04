package hl7v26

// QRI - Query Response Instance
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/QRI
type QRI struct {
	CandidateConfidence *int `hl7:"POS=2"`
	MatchReasonCode []string `hl7:"POS=3"`
	AlgorithmDescriptor CWE `hl7:"POS=4"`
}

