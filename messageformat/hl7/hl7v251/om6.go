package hl7v251

// OM6 - Observations that are Calculated from Other Observations
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/OM6
type OM6 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	DerivationRule string `hl7:"POS=3"`
}

