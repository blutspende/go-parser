package hl7v23

import "time"

// ORC - Common order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ORC
type ORC struct {
	OrderControl string `hl7:"POS=2"`
	PlacerOrderNumber EI `hl7:"POS=3"`
	FillerOrderNumber EI `hl7:"POS=4"`
	PlacerGroupNumber EI `hl7:"POS=5"`
	OrderStatus string `hl7:"POS=6"`
	ResponseFlag string `hl7:"POS=7"`
	QuantityTiming TQ `hl7:"POS=8"`
	ParentOrder CM_EIP `hl7:"POS=9"`
	DateTimeOfTransaction time.Time `hl7:"POS=10"`
	EnteredBy XCN `hl7:"POS=11"`
	VerifiedBy XCN `hl7:"POS=12"`
	OrderingProvider XCN `hl7:"POS=13"`
	EnterersLocation PL `hl7:"POS=14"`
	CallBackPhoneNumber []XTN `hl7:"POS=15"`
	OrderEffectiveDateTime time.Time `hl7:"POS=16"`
	OrderControlCodeReason CE `hl7:"POS=17"`
	EnteringOrganization CE `hl7:"POS=18"`
	EnteringDevice CE `hl7:"POS=19"`
	ActionBy XCN `hl7:"POS=20"`
}

