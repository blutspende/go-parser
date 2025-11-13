package hl7v24

// HL7 v2.4 - BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/BLG
type BLG struct {
	WhenToCharge CCD    `hl7:"POS=2" json:"WhenToCharge,omitempty"`
	ChargeType   string `hl7:"POS=3" json:"ChargeType,omitempty"`
	AccountID    CX     `hl7:"POS=4" json:"AccountID,omitempty"`
}
