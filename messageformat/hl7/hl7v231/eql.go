package hl7v231

// EQL - Embedded query language segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/EQL
type EQL struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseFormatCode string `hl7:"POS=3"`
	EqlQueryName CE `hl7:"POS=4"`
	EqlQueryStatement string `hl7:"POS=5"`
}

