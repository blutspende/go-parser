package hl7v27

// RCL_I06_PROVIDER - Group struct
type RCL_I06_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RCL_I06 - Return clinical list
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/RCL_I06
type RCL_I06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA;ATR=optional"`
	Provider []RCL_I06_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

