package hl7v24

// QRY_Q26 - Pharmacy/treatment order query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QRY_Q26
type QRY_Q26 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

