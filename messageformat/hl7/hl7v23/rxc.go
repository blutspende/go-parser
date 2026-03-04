package hl7v23

// RXC - Pharmacy component order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RXC
type RXC struct {
	RxComponentType string `hl7:"POS=2"`
	ComponentCode CE `hl7:"POS=3"`
	ComponentAmount *int `hl7:"POS=4"`
	ComponentUnits CE `hl7:"POS=5"`
	ComponentStrength *int `hl7:"POS=6"`
	ComponentStrengthUnits CE `hl7:"POS=7"`
}

