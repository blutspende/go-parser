package hl7v28

// CDO - Cumulative Dosage
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/CDO
type CDO struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	ActionCode string `hl7:"POS=3"`
	CumulativeDosageLimit CQ `hl7:"POS=4"`
	CumulativeDosageLimitTimeInterval CQ `hl7:"POS=5"`
}

