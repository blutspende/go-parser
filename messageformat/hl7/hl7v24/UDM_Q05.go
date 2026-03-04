package hl7v24

// UDM_Q05 - Unsolicited display update message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/UDM_Q05
type UDM_Q05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	ResultsUpdateDefinition URD `hl7:"TAG=URD"`
	UnsolicitedSelection URS `hl7:"TAG=URS;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

