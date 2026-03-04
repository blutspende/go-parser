package hl7v271

// OM6 - Observations That Are Calculated From Other Observations
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/OM6
type OM6 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	DerivationRule string `hl7:"POS=3"`
}

