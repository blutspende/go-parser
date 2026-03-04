package hl7v251

// QAK - Query Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/QAK
type QAK struct {
	QueryTag string `hl7:"POS=2"`
	QueryResponseStatus string `hl7:"POS=3"`
	MessageQueryName CE `hl7:"POS=4"`
	HitCount *int `hl7:"POS=5"`
	ThisPayload *int `hl7:"POS=6"`
	HitsRemaining *int `hl7:"POS=7"`
}

