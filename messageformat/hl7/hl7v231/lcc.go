package hl7v231

// LCC - Location charge code segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/LCC
type LCC struct {
	PrimaryKeyValueLcc PL `hl7:"POS=2"`
	LocationDepartment string `hl7:"POS=3"`
	AccommodationType []CE `hl7:"POS=4"`
	ChargeCode []CE `hl7:"POS=5"`
}

