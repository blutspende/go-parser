package hl7v25

// BTS - Batch Trailer Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/BTS
type BTS struct {
	BatchMessageCount string `hl7:"POS=2"`
	BatchComment string `hl7:"POS=3"`
	BatchTotals []*int `hl7:"POS=4"`
}

