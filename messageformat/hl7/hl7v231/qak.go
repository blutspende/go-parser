package hl7v231

// QAK - Query Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/QAK
type QAK struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseStatus string `hl7:"POS=3"`
}

