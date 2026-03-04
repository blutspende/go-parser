package hl7v22

// RXC - Pharmacy Component Order
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/RXC
type RXC struct {
	RxComponentType string `hl7:"POS=2"`
	ComponentCode CE `hl7:"POS=3"`
	ComponentAmount *int `hl7:"POS=4"`
	ComponentUnits CE `hl7:"POS=5"`
}

