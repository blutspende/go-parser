package hl7v23

// NPU - Bed status update
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/NPU
type NPU struct {
	BedLocation PL `hl7:"POS=2"`
	BedStatus string `hl7:"POS=3"`
}

