package hl7v26

// SLT - Sterilization Lot Segment
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/SLT
type SLT struct {
	DeviceNumber EI `hl7:"POS=2"`
	DeviceName string `hl7:"POS=3"`
	LotNumber EI `hl7:"POS=4"`
	ItemIdentifier EI `hl7:"POS=5"`
	BarCode string `hl7:"POS=6"`
}

