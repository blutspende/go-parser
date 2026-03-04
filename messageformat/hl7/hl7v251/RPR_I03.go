package hl7v251

// RPR_I03_PROVIDER - Group struct
type RPR_I03_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPR_I03 - Request/receipt of patient selection list acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/RPR_I03
type RPR_I03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Provider []RPR_I03_PROVIDER `hl7:"GROUP"`
	PatientIdentification []PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

