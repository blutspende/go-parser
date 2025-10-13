package hl7v24

import "time"

// HL7 v2.4 - ORC - Common Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ORC
type ORC struct {
	OrderControl                  string    `hl7:"2" json:"OrderControl,omitempty"`
	PlacerOrderNumber             EI        `hl7:"3" json:"PlacerOrderNumber,omitempty"`
	FillerOrderNumber             EI        `hl7:"4" json:"FillerOrderNumber,omitempty"`
	PlacerGroupNumber             EI        `hl7:"5" json:"PlacerGroupNumber,omitempty"`
	OrderStatus                   string    `hl7:"6" json:"OrderStatus,omitempty"`
	ResponseFlag                  string    `hl7:"7" json:"ResponseFlag,omitempty"`
	QuantityTiming                []TQ      `hl7:"8" json:"QuantityTiming,omitempty"`
	ParentOrder                   EIP       `hl7:"9" json:"ParentOrder,omitempty"`
	DateTimeOfTransaction         time.Time `hl7:"10" json:"DateTimeOfTransaction,omitempty"`
	EnteredBy                     []XCN     `hl7:"11" json:"EnteredBy,omitempty"`
	VerifiedBy                    []XCN     `hl7:"12" json:"VerifiedBy,omitempty"`
	OrderingProvider              []XCN     `hl7:"13" json:"OrderingProvider,omitempty"`
	EnterersLocation              PL        `hl7:"14" json:"EnterersLocation,omitempty"`
	CallBackPhoneNumber           []XTN     `hl7:"15" json:"CallBackPhoneNumber,omitempty"`
	OrderEffectiveDateTime        time.Time `hl7:"16" json:"OrderEffectiveDateTime,omitempty"`
	OrderControlCodeReason        CE        `hl7:"17" json:"OrderControlCodeReason,omitempty"`
	EnteringOrganization          CE        `hl7:"18" json:"EnteringOrganization,omitempty"`
	EnteringDevice                CE        `hl7:"19" json:"EnteringDevice,omitempty"`
	ActionBy                      []XCN     `hl7:"20" json:"ActionBy,omitempty"`
	AdvancedBeneficiaryNoticeCode CE        `hl7:"21" json:"AdvancedBeneficiaryNoticeCode,omitempty"`
	OrderingFacilityName          []XON     `hl7:"22" json:"OrderingFacilityName,omitempty"`
	OrderingFacilityPhoneNumber   []XTN     `hl7:"23" json:"OrderingFacilityPhoneNumber,omitempty"`
	OrderingProviderAddress       []XAD     `hl7:"24" json:"OrderingProviderAddress,omitempty"`
	OrderStatusModifier           CWE       `hl7:"25" json:"OrderStatusModifier,omitempty"`
}
