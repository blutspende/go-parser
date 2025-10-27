package hl7v23

import "time"

// ORC - Common order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ORC
type ORC struct {
	OrderControl           string    `hl7:"POS=2" json:"orderControl,omitempty"`
	PlacerOrderNumber      EI        `hl7:"POS=3" json:"placerOrderNumber,omitempty"`
	FillerOrderNumber      EI        `hl7:"POS=4" json:"fillerOrderNumber,omitempty"`
	PlacerGroupNumber      EI        `hl7:"POS=5" json:"placerGroupNumber,omitempty"`
	OrderStatus            string    `hl7:"POS=6" json:"orderStatus,omitempty"`
	ResponseFlag           string    `hl7:"POS=7" json:"responseFlag,omitempty"`
	QuantityTiming         TQ        `hl7:"POS=8" json:"quantityTiming,omitempty"`
	ParentOrder            CM_EIP    `hl7:"POS=9" json:"parentOrder,omitempty"`
	DateTimeOfTransaction  time.Time `hl7:"POS=10" json:"dateTimeOfTransaction,omitempty"`
	EnteredBy              XCN       `hl7:"POS=11" json:"enteredBy,omitempty"`
	VerifiedBy             XCN       `hl7:"POS=12" json:"verifiedBy,omitempty"`
	OrderingProvider       XCN       `hl7:"POS=13" json:"orderingProvider,omitempty"`
	EnterersLocation       PL        `hl7:"POS=14" json:"enterersLocation,omitempty"`
	CallBackPhoneNumber    []XTN     `hl7:"POS=15" json:"callBackPhoneNumber,omitempty"`
	OrderEffectiveDateTime time.Time `hl7:"POS=16" json:"orderEffectiveDateTime,omitempty"`
	OrderControlCodeReason CE        `hl7:"POS=17" json:"orderControlCodeReason,omitempty"`
	EnteringOrganization   CE        `hl7:"POS=18" json:"enteringOrganization,omitempty"`
	EnteringDevice         CE        `hl7:"POS=19" json:"enteringDevice,omitempty"`
	ActionBy               XCN       `hl7:"POS=20" json:"actionBy,omitempty"`
}
