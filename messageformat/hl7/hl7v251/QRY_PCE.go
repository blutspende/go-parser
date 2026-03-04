package hl7v251

// QRY_PCE - Patient pathway (problem-oriented) query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QRY_PCE
type QRY_PCE struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

