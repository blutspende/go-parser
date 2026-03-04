package hl7v27

// RXC - Pharmacy/treatment Component Order
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/RXC
type RXC struct {
	RxComponentType string `hl7:"POS=2"`
	ComponentCode CWE `hl7:"POS=3"`
	ComponentAmount *int `hl7:"POS=4"`
	ComponentUnits CWE `hl7:"POS=5"`
	ComponentStrength *int `hl7:"POS=6"`
	ComponentStrengthUnits CWE `hl7:"POS=7"`
	SupplementaryCode []CWE `hl7:"POS=8"`
	ComponentDrugStrengthVolume *int `hl7:"POS=9"`
	ComponentDrugStrengthVolumeUnits CWE `hl7:"POS=10"`
}

