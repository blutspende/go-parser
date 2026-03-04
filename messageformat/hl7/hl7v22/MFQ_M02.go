package hl7v22

// MFQ_M02 - Staff Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/MFQ_M02
type MFQ_M02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

