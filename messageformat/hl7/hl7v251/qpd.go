package hl7v251

// QPD - Query Parameter Definition
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/QPD
type QPD struct {
	MessageQueryName CE `hl7:"POS=2"`
	QueryTag string `hl7:"POS=3"`
	UserParametersInSuccessiveFields string `hl7:"POS=4"`
}

