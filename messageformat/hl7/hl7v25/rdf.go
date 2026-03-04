package hl7v25

// RDF - Table Row Definition
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/RDF
type RDF struct {
	NumberOfColumnsPerRow *int `hl7:"POS=2"`
	ColumnDescription []RCD `hl7:"POS=3"`
}

