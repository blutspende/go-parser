package hl7v23

import "time"

// ORC - Common order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ORC
type ORC struct {
	OrderControl           string    `hl7:"2" json:"orderControl,omitempty"`
	PlacerOrderNumber      EI        `hl7:"3" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber      EI        `hl7:"4" json:"fillerOrderNumber,omitempty"`
	PlacerGroupNumber      EI        `hl7:"5" json:"placerGroupNumber,omitempty"`
	OrderStatus            string    `hl7:"6" json:"orderStatus,omitempty"`
	ResponseFlag           string    `hl7:"7" json:"responseFlag,omitempty"`
	QuantityTiming         TQ        `hl7:"8" json:"quantityTiming,omitempty"`
	ParentOrder            CM_EIP    `hl7:"9" json:"parentOrder,omitempty"`
	DateTimeOfTransaction  time.Time `hl7:"10,longdate" json:"dateTimeOfTransaction,omitempty"`
	EnteredBy              XCN       `hl7:"11" json:"enteredBy,omitempty"`
	VerifiedBy             XCN       `hl7:"12" json:"verifiedBy,omitempty"`
	OrderingProvider       XCN       `hl7:"13" json:"orderingProvider,omitempty"`
	EnterersLocation       PL        `hl7:"14" json:"enterersLocation,omitempty"`
	CallBackPhoneNumber    []XTN     `hl7:"15" json:"callBackPhoneNumber,omitempty"`
	OrderEffectiveDateTime time.Time `hl7:"16" json:"orderEffectiveDateTime,omitempty"`
	OrderControlCodeReason CE        `hl7:"17" json:"orderControlCodeReason,omitempty"`
	EnteringOrganization   CE        `hl7:"18" json:"enteringOrganization,omitempty"`
	EnteringDevice         CE        `hl7:"19" json:"enteringDevice,omitempty"`
	ActionBy               XCN       `hl7:"20" json:"actionBy,omitempty"`
}
