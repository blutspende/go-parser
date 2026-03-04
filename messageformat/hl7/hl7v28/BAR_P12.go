package hl7v28

// BAR_P12_PROCEDURE - Group struct
type BAR_P12_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// BAR_P12 - Update Diagnosis/Procedure
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/BAR_P12
type BAR_P12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []BAR_P12_PROCEDURE `hl7:"GROUP;ATR=optional"`
	ObservationResult OBX `hl7:"TAG=OBX;ATR=optional"`
}

