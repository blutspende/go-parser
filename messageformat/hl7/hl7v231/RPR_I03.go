package hl7v231

// RPR_I03_PROVIDER - Group struct
type RPR_I03_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPR_I03 - Return patient list
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RPR_I03
type RPR_I03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	Provider []RPR_I03_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment []PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

