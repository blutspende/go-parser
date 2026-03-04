package hl7v251

// RXR - Pharmacy/Treatment Route
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/RXR
type RXR struct {
	Route CE `hl7:"POS=2"`
	AdministrationSite CWE `hl7:"POS=3"`
	AdministrationDevice CE `hl7:"POS=4"`
	AdministrationMethod CWE `hl7:"POS=5"`
	RoutingInstruction CE `hl7:"POS=6"`
	AdministrationSiteModifier CWE `hl7:"POS=7"`
}

