package hl7v251

// UDM_Q05 - Unsolicited Display Update
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/UDM_Q05
type UDM_Q05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	ResultsUpdateDefinition URD `hl7:"TAG=URD"`
	UnsolicitedSelection URS `hl7:"TAG=URS;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

