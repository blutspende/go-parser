package hl7v28

// ORX_O43_RESPONSE - Group struct
type ORX_O43_RESPONSE struct {
	Patient ORX_O43_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORX_O43_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORX_O43_RESPONSE_PATIENT - Group struct
type ORX_O43_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
}

// ORX_O43_RESPONSE_ORDER - Group struct
type ORX_O43_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORX_O43 - General Order Message with Document Payload Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ORX_O43
type ORX_O43 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORX_O43_RESPONSE `hl7:"GROUP;ATR=optional"`
}

