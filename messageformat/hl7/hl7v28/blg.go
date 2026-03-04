package hl7v28

// BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/BLG
type BLG struct {
	WhenToCharge CCD `hl7:"POS=2"`
	ChargeType string `hl7:"POS=3"`
	AccountID CX `hl7:"POS=4"`
	ChargeTypeReason CWE `hl7:"POS=5"`
}

