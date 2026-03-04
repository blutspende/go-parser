package hl7v22

// BTS - Batch Trailer
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/BTS
type BTS struct {
	BatchMessageCount string `hl7:"POS=2"`
	BatchComment string `hl7:"POS=3"`
	BatchTotals []CM_BATCH_TOTAL `hl7:"POS=4"`
}

