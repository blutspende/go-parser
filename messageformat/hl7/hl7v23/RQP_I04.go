package hl7v23

// RQP_I04_PROVIDER - Group struct
type RQP_I04_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQP_I04 - Request patient demographics
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RQP_I04
type RQP_I04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Provider []RQP_I04_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

