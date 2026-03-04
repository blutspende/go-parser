package hl7v24

// RDY_K15 - Display based response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RDY_K15
type RDY_K15 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	DisplayData []DSP `hl7:"TAG=DSP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

