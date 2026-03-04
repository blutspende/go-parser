package hl7v22

// DSR_R03 - oriented results, query/unsolicited update
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/DSR_R03
type DSR_R03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA;ATR=optional"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

