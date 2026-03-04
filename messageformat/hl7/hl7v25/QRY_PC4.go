package hl7v25

// QRY_PC4 - Patient care problem query
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QRY_PC4
type QRY_PC4 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

