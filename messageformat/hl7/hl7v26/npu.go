package hl7v26

// NPU - Bed Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/NPU
type NPU struct {
	BedLocation PL `hl7:"POS=2"`
	BedStatus string `hl7:"POS=3"`
}

