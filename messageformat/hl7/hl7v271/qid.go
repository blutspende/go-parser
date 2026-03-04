package hl7v271

// QID - Query Identification
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/QID
type QID struct {
	QueryTag string `hl7:"POS=2"`
	MessageQueryName CWE `hl7:"POS=3"`
}

