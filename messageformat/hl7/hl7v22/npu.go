package hl7v22

// NPU - Bed Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/NPU
type NPU struct {
	BedLocation CM_INTERNAL_LOCATION `hl7:"POS=2"`
	BedStatus string `hl7:"POS=3"`
}

