package hl7v23

// RPL_I02_PROVIDER - Group struct
type RPL_I02_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPL_I02 - Return patient display list
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RPL_I02
type RPL_I02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	Provider []RPL_I02_PROVIDER `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	DisplayDataSegment []DSP `hl7:"TAG=DSP;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

