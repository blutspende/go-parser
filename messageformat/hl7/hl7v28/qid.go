package hl7v28

// QID - Query Identification
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/QID
type QID struct {
	QueryTag string `hl7:"POS=2"`
	MessageQueryName CWE `hl7:"POS=3"`
}

