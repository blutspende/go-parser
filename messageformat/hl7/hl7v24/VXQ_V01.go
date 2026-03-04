package hl7v24

// VXQ_V01 - Query for vaccination record
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/VXQ_V01
type VXQ_V01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
}

