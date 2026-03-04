package hl7v23

// QAK - Query Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/QAK
type QAK struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseStatus string `hl7:"POS=3"`
}

