package hl7v25

// RDY_Z80 - Dispense Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/RDY_Z80
type RDY_Z80 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	DisplayData []DSP `hl7:"TAG=DSP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

