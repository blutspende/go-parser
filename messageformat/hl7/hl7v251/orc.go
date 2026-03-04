package hl7v251

import "time"

// ORC - Common Order
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/ORC
type ORC struct {
	OrderControl string `hl7:"POS=2"`
	PlacerOrderNumber EI `hl7:"POS=3"`
	FillerOrderNumber EI `hl7:"POS=4"`
	PlacerGroupNumber EI `hl7:"POS=5"`
	OrderStatus string `hl7:"POS=6"`
	ResponseFlag string `hl7:"POS=7"`
	QuantityTiming []TQ `hl7:"POS=8"`
	ParentOrder EIP `hl7:"POS=9"`
	DateTimeOfTransaction time.Time `hl7:"POS=10"`
	EnteredBy []XCN `hl7:"POS=11"`
	VerifiedBy []XCN `hl7:"POS=12"`
	OrderingProvider []XCN `hl7:"POS=13"`
	EnterersLocation PL `hl7:"POS=14"`
	CallBackPhoneNumber []XTN `hl7:"POS=15"`
	OrderEffectiveDateTime time.Time `hl7:"POS=16"`
	OrderControlCodeReason CE `hl7:"POS=17"`
	EnteringOrganization CE `hl7:"POS=18"`
	EnteringDevice CE `hl7:"POS=19"`
	ActionBy []XCN `hl7:"POS=20"`
	AdvancedBeneficiaryNoticeCode CE `hl7:"POS=21"`
	OrderingFacilityName []XON `hl7:"POS=22"`
	OrderingFacilityAddress []XAD `hl7:"POS=23"`
	OrderingFacilityPhoneNumber []XTN `hl7:"POS=24"`
	OrderingProviderAddress []XAD `hl7:"POS=25"`
	OrderStatusModifier CWE `hl7:"POS=26"`
	AdvancedBeneficiaryNoticeOverrideReason CWE `hl7:"POS=27"`
	FillersExpectedAvailabilityDateTime time.Time `hl7:"POS=28"`
	ConfidentialityCode CWE `hl7:"POS=29"`
	OrderType CWE `hl7:"POS=30"`
	EntererAuthorizationMode CNE `hl7:"POS=31"`
	ParentUniversalServiceIdentifier CWE `hl7:"POS=32"`
}

