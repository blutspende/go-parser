package hl7v25

// ADT_A44_PATIENT_ID - Group struct
type ADT_A44_PATIENT_ID struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
}

// ADT_A44 - Move Account Information - Patient Account Number
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ADT_A44
type ADT_A44 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientID []ADT_A44_PATIENT_ID `hl7:"GROUP"`
}

