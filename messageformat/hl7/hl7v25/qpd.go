package hl7v25

// QPD - Query Parameter Definition
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/QPD
type QPD struct {
	MessageQueryName CE `hl7:"POS=2"`
	QueryTag string `hl7:"POS=3"`
	UserParametersInSuccessiveFields string `hl7:"POS=4"`
}

