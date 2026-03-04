package hl7v23

// UDM_Q05 - Unsolicited display update message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/UDM_Q05
type UDM_Q05 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ResultsUpdateDefinitionSegment URD `hl7:"TAG=URD"`
	UnsolicitedSelectionSegment URS `hl7:"TAG=URS;ATR=optional"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

