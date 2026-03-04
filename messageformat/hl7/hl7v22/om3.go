package hl7v22

// OM3 - Categorical Test/observation
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM3
type OM3 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	PreferredCodingSystem string `hl7:"POS=4"`
	ValidCodedAnswers CE `hl7:"POS=5"`
	NormalTestCodesForCategoricalObservations []CE `hl7:"POS=6"`
	AbnormalTestCodesForCategoricalObservations CE `hl7:"POS=7"`
	CriticalTestCodesForCategoricalObservations CE `hl7:"POS=8"`
	DataType string `hl7:"POS=9"`
}

