package hl7v231

// RDF - Table row definition segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/RDF
type RDF struct {
	NumberOfColumnsPerRow *int `hl7:"POS=2"`
	ColumnDescription []RCD `hl7:"POS=3"`
}

