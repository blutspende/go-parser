package hl7v28

import "time"

// ADJ - Adjustment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/ADJ
type ADJ struct {
	ProviderAdjustmentNumber EI `hl7:"POS=2"`
	PayerAdjustmentNumber EI `hl7:"POS=3"`
	AdjustmentSequenceNumber int `hl7:"POS=4;ATR=sequence"`
	AdjustmentCategory CWE `hl7:"POS=5"`
	AdjustmentAmount CP `hl7:"POS=6"`
	AdjustmentQuantity CQ `hl7:"POS=7"`
	AdjustmentReasonPa CWE `hl7:"POS=8"`
	AdjustmentDescription string `hl7:"POS=9"`
	OriginalValue *int `hl7:"POS=10"`
	SubstituteValue *int `hl7:"POS=11"`
	AdjustmentAction CWE `hl7:"POS=12"`
	ProviderAdjustmentNumberCrossReference EI `hl7:"POS=13"`
	ProviderProductServiceLineItemNumberCrossReference EI `hl7:"POS=14"`
	AdjustmentDate time.Time `hl7:"POS=15"`
	ResponsibleOrganization XON `hl7:"POS=16"`
}

