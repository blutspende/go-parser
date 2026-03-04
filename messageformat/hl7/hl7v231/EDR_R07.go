package hl7v231

// EDR_R07 - Enhanced display response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/EDR_R07
type EDR_R07 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgement QAK `hl7:"TAG=QAK"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

