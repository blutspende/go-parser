package hl7v271

// BTS - Batch Trailer
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/BTS
type BTS struct {
	BatchMessageCount string `hl7:"POS=2"`
	BatchComment string `hl7:"POS=3"`
	BatchTotals []*int `hl7:"POS=4"`
}

