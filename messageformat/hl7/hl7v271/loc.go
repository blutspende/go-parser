package hl7v271

// LOC - Location Identification
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/LOC
type LOC struct {
	PrimaryKeyValueLoc PL `hl7:"POS=2"`
	LocationDescription string `hl7:"POS=3"`
	LocationTypeLoc []CWE `hl7:"POS=4"`
	OrganizationNameLoc []XON `hl7:"POS=5"`
	LocationAddress []XAD `hl7:"POS=6"`
	LocationPhone []XTN `hl7:"POS=7"`
	LicenseNumber []CWE `hl7:"POS=8"`
	LocationEquipment []CWE `hl7:"POS=9"`
	LocationServiceCode CWE `hl7:"POS=10"`
}

