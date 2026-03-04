package hl7v24

// QPD - Query Parameter Definition
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/QPD
type QPD struct {
	MessageQueryName CE `hl7:"POS=2"`
	QueryTag string `hl7:"POS=3"`
	UserParameters string `hl7:"POS=4"`
}

