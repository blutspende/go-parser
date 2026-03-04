package hl7v24

// QRY_A19 - Patient query
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QRY_A19
type QRY_A19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

