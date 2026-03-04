package hl7v231

// NPU - Bed status update segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/NPU
type NPU struct {
	BedLocation PL `hl7:"POS=2"`
	BedStatus string `hl7:"POS=3"`
}

