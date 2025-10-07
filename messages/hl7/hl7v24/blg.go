package hl7v24

// HL7 v2.4 - BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/BLG
type BLG struct {
	WhenToCharge CCD    `hl7:"1" json:"WhenToCharge,omitempty"`
	ChargeType   string `hl7:"2" json:"ChargeType,omitempty"`
	AccountID    CX     `hl7:"3" json:"AccountID,omitempty"`
}
