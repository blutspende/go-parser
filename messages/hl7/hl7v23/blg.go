package hl7v23

// BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/BLG
type BLG struct {
	WhenToCharge CM_CCD `hl7:"2" json:"whenToCharge,omitempty"`
	ChargeType   string `hl7:"3" json:"chargeType,omitempty"`
	AccountID    CK     `hl7:"4" json:"accountID,omitempty"`
}
