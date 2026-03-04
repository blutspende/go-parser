package hl7v22

// QRY_Q28 - Pharmacy Dispense Information Query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/QRY_Q28
type QRY_Q28 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

