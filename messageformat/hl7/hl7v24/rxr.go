package hl7v24

// RXR - Pharmacy/Treatment Route
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/RXR
type RXR struct {
	Route CE `hl7:"POS=2"`
	AdministrationSite CE `hl7:"POS=3"`
	AdministrationDevice CE `hl7:"POS=4"`
	AdministrationMethod CE `hl7:"POS=5"`
	RoutingInstruction CE `hl7:"POS=6"`
}

