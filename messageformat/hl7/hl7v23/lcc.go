package hl7v23

// LCC - Location Charge Code
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/LCC
type LCC struct {
	PrimaryKeyValue PL `hl7:"POS=2"`
	LocationDepartment string `hl7:"POS=3"`
	AccommodationType []CE `hl7:"POS=4"`
	ChargeCode []CE `hl7:"POS=5"`
}

