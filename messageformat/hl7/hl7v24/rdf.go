package hl7v24

// RDF - Table Row Definition
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/RDF
type RDF struct {
	NumberOfColumnsPerRow *int `hl7:"POS=2"`
	ColumnDescription []RCD `hl7:"POS=3"`
}

