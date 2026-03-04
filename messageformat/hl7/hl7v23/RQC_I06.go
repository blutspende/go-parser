package hl7v23

// RQC_I06_PROVIDER - Group struct
type RQC_I06_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQC_I06 - Request/receipt of clinical data listing
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RQC_I06
type RQC_I06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Provider []RQC_I06_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

