package hl7v22

import "time"

// ORC - Common Order
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/ORC
type ORC struct {
	OrderControl string `hl7:"POS=2"`
	PlacerOrderNumber CM_PLACER `hl7:"POS=3"`
	FillerOrderNumber CM_FILLER `hl7:"POS=4"`
	PlacerGroupNumber CM_GROUP_ID `hl7:"POS=5"`
	OrderStatus string `hl7:"POS=6"`
	ResponseFlag string `hl7:"POS=7"`
	QuantityTiming []TQ `hl7:"POS=8"`
	Parent CM_EIP `hl7:"POS=9"`
	DateTimeOfTransaction time.Time `hl7:"POS=10"`
	EnteredBy CN_PERSON `hl7:"POS=11"`
	VerifiedBy CN_PERSON `hl7:"POS=12"`
	OrderingProvider CN_PERSON `hl7:"POS=13"`
	EnterersLocation PL `hl7:"POS=14"`
	CallBackPhoneNumber []string `hl7:"POS=15"`
	OrderEffectiveDateTime time.Time `hl7:"POS=16"`
	OrderControlCodeReason CE `hl7:"POS=17"`
	EnteringOrganization CE `hl7:"POS=18"`
	EnteringDevice CE `hl7:"POS=19"`
	ActionBy CN_PERSON `hl7:"POS=20"`
}

