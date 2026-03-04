package hl7v23

// BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/BLG
type BLG struct {
	WhenToCharge CM_CCD `hl7:"POS=2"`
	ChargeType string `hl7:"POS=3"`
	AccountID CK `hl7:"POS=4"`
}

