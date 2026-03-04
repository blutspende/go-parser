package hl7v24

// OM6 - Observations that are Calculated from Other Observ
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OM6
type OM6 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	DerivationRule string `hl7:"POS=3"`
}

