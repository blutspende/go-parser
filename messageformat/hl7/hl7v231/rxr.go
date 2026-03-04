package hl7v231

// RXR - Pharmacy/treatment route segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/RXR
type RXR struct {
	Route CE `hl7:"POS=2"`
	Site CE `hl7:"POS=3"`
	AdministrationDevice CE `hl7:"POS=4"`
	AdministrationMethod CE `hl7:"POS=5"`
	RoutingInstruction CE `hl7:"POS=6"`
}

