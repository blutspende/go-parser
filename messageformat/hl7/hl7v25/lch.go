package hl7v25

// LCH - Location Characteristic
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/LCH
type LCH struct {
	PrimaryKeyValueLch PL `hl7:"POS=2"`
	SegmentActionCode string `hl7:"POS=3"`
	SegmentUniqueKey EI `hl7:"POS=4"`
	LocationCharacteristicID CE `hl7:"POS=5"`
	LocationCharacteristicValueLch CE `hl7:"POS=6"`
}

