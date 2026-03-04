package hl7v24

// QRY_T12 - Document query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QRY_T12
type QRY_T12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

