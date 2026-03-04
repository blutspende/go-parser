package hl7v24

// OM3 - Categorical Service/Test/Observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OM3
type OM3 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	PreferredCodingSystem CE `hl7:"POS=3"`
	ValidCodedAnswers CE `hl7:"POS=4"`
	NormalTextCodesForCategoricalObservations []CE `hl7:"POS=5"`
	AbnormalTextCodesForCategoricalObservations CE `hl7:"POS=6"`
	CriticalTextCodesForCategoricalObservations CE `hl7:"POS=7"`
	ValueType string `hl7:"POS=8"`
}

