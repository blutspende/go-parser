package hl7v22

// QRY_Q26 - Pharmacy Prescription Order Query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/QRY_Q26
type QRY_Q26 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

