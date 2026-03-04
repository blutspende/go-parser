package hl7v22

// QRY_A19 - Patient Query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/QRY_A19
type QRY_A19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

