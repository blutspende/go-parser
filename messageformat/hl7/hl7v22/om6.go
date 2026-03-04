package hl7v22

// OM6 - Observations That Are Calculated From Other Observations
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM6
type OM6 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	DerivationRule string `hl7:"POS=4"`
}

