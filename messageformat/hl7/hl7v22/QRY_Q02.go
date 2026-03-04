package hl7v22

// QRY_Q02 - A Query is Sent for Deferred Response
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/QRY_Q02
type QRY_Q02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

