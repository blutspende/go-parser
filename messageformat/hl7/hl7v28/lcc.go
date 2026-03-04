package hl7v28

// LCC - Location Charge Code
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/LCC
type LCC struct {
	PrimaryKeyValueLcc PL `hl7:"POS=2"`
	LocationDepartment CWE `hl7:"POS=3"`
	AccommodationType []CWE `hl7:"POS=4"`
	ChargeCode []CWE `hl7:"POS=5"`
}

