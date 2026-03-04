package hl7v26

// OM3 - Categorical Service/Test/Observation
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/OM3
type OM3 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	PreferredCodingSystem CWE `hl7:"POS=3"`
	ValidCodedAnswers []CWE `hl7:"POS=4"`
	NormalTextCodesForCategoricalObservations []CWE `hl7:"POS=5"`
	AbnormalTextCodesForCategoricalObservations []CWE `hl7:"POS=6"`
	CriticalTextCodesForCategoricalObservations []CWE `hl7:"POS=7"`
	ValueType string `hl7:"POS=8"`
}

