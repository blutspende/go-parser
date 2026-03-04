package hl7v24

// QID - Query Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/QID
type QID struct {
	QueryTag string `hl7:"POS=2"`
	MessageQueryName CE `hl7:"POS=3"`
}

