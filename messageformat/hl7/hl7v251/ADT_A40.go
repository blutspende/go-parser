package hl7v251

// ADT_A40_PATIENT_ID - Group struct
type ADT_A40_PATIENT_ID struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A40 - Merge Patient - Patient Identifier List
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ADT_A40
type ADT_A40 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientID []ADT_A40_PATIENT_ID `hl7:"GROUP"`
}

