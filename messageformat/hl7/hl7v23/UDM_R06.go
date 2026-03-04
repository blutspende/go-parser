package hl7v23

// UDM_R06 - Continuation of unsolicited display update message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/UDM_R06
type UDM_R06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ResultsUpdateDefinitionSegment URD `hl7:"TAG=URD"`
	UnsolicitedSelectionSegment URS `hl7:"TAG=URS;ATR=optional"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

