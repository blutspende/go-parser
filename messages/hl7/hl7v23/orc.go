package hl7v23

import "time"

// ORC - Common order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ORC
type ORC struct {
	OrderControl           string    `hl7:"1" json:"orderControl,omitempty"`
	PlacerOrderNumber      EI        `hl7:"2" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber      EI        `hl7:"3" json:"fillerOrderNumber,omitempty"`
	PlacerGroupNumber      EI        `hl7:"4" json:"placerGroupNumber,omitempty"`
	OrderStatus            string    `hl7:"5" json:"orderStatus,omitempty"`
	ResponseFlag           string    `hl7:"6" json:"responseFlag,omitempty"`
	QuantityTiming         TQ        `hl7:"7" json:"quantityTiming,omitempty"`
	ParentOrder            CM_EIP    `hl7:"8" json:"parentOrder,omitempty"`
	DateTimeOfTransaction  time.Time `hl7:"9" json:"dateTimeOfTransaction,omitempty"`
	EnteredBy              XCN       `hl7:"10" json:"enteredBy,omitempty"`
	VerifiedBy             XCN       `hl7:"11" json:"verifiedBy,omitempty"`
	OrderingProvider       XCN       `hl7:"12" json:"orderingProvider,omitempty"`
	EnterersLocation       PL        `hl7:"13" json:"enterersLocation,omitempty"`
	CallBackPhoneNumber    []XTN     `hl7:"14" json:"callBackPhoneNumber,omitempty"`
	OrderEffectiveDateTime time.Time `hl7:"15" json:"orderEffectiveDateTime,omitempty"`
	OrderControlCodeReason CE        `hl7:"16" json:"orderControlCodeReason,omitempty"`
	EnteringOrganization   CE        `hl7:"17" json:"enteringOrganization,omitempty"`
	EnteringDevice         CE        `hl7:"18" json:"enteringDevice,omitempty"`
	ActionBy               XCN       `hl7:"19" json:"actionBy,omitempty"`
}
