package hl7v22

// DSR_Q03 - Deferred Response to A Query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/DSR_Q03
type DSR_Q03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA;ATR=optional"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

