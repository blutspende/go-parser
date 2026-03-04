package hl7v23

// OM6 - Observations that are calculated from other observations
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OM6
type OM6 struct {
	SequenceNumber *int `hl7:"POS=2"`
	DerivationRule string `hl7:"POS=3"`
}

