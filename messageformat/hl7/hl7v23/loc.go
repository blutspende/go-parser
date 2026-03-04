package hl7v23

// LOC - Location Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/LOC
type LOC struct {
	PrimaryKeyValue PL `hl7:"POS=2"`
	LocationDescription string `hl7:"POS=3"`
	LocationType []string `hl7:"POS=4"`
	OrganizationName XON `hl7:"POS=5"`
	LocationAddress XAD `hl7:"POS=6"`
	LocationPhone []XTN `hl7:"POS=7"`
	LicenseNumber []CE `hl7:"POS=8"`
	LocationEquipment []string `hl7:"POS=9"`
}

