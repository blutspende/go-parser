package hl7v24

// QRY_Q30 - Pharmacy/treatment dose information query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QRY_Q30
type QRY_Q30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

