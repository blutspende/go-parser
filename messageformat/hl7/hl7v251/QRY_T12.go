package hl7v251

// QRY_T12 - Document Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QRY_T12
type QRY_T12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

