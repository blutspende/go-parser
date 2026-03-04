package hl7v23

// LCH - Location Characteristic
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/LCH
type LCH struct {
	PrimaryKeyValue PL `hl7:"POS=2"`
	SegmentActionCode string `hl7:"POS=3"`
	SegmentUniqueKey EI `hl7:"POS=4"`
	LocationCharacteristicID CE `hl7:"POS=5"`
	LocationCharacteristicValue CE `hl7:"POS=6"`
}

