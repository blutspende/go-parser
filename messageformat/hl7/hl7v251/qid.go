package hl7v251

// QID - Query Identification
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/QID
type QID struct {
	QueryTag string `hl7:"POS=2"`
	MessageQueryName CE `hl7:"POS=3"`
}

