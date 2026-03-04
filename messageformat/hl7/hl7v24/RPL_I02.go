package hl7v24

// RPL_I02_PROVIDER - Group struct
type RPL_I02_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPL_I02 - Request/receipt of patient selection display list acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RPL_I02
type RPL_I02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Provider []RPL_I02_PROVIDER `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

