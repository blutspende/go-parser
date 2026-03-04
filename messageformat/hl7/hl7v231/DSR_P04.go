package hl7v231

// DSR_P04 - Display response message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/DSR_P04
type DSR_P04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

