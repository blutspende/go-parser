package hl7v23

// DSR_Q03 - Deferred response to a query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/DSR_Q03
type DSR_Q03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA;ATR=optional"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

