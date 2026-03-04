package hl7v251

// DSR_Q03 - Deferred Response to a Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/DSR_Q03
type DSR_Q03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA;ATR=optional"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

