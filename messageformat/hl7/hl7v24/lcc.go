package hl7v24

// LCC - Location Charge Code
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/LCC
type LCC struct {
	PrimaryKeyValueLcc PL `hl7:"POS=2"`
	LocationDepartment CE `hl7:"POS=3"`
	AccommodationType []CE `hl7:"POS=4"`
	ChargeCode []CE `hl7:"POS=5"`
}

