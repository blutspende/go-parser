package hl7v27

// RXR - Pharmacy/treatment Route
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/RXR
type RXR struct {
	Route CWE `hl7:"POS=2"`
	AdministrationSite CWE `hl7:"POS=3"`
	AdministrationDevice CWE `hl7:"POS=4"`
	AdministrationMethod CWE `hl7:"POS=5"`
	RoutingInstruction CWE `hl7:"POS=6"`
	AdministrationSiteModifier CWE `hl7:"POS=7"`
}

