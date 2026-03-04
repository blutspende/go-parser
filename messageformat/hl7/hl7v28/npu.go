package hl7v28

// NPU - Bed Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/NPU
type NPU struct {
	BedLocation PL `hl7:"POS=2"`
	BedStatus CWE `hl7:"POS=3"`
}

