package hl7v24

// ADT_A41_PATIENT_ID - Group struct
type ADT_A41_PATIENT_ID struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	MergePatientInformation MRG `hl7:"TAG=MRG"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ADT_A41 - Merge account - patient account number
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A41
type ADT_A41 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientID []ADT_A41_PATIENT_ID `hl7:"GROUP"`
}

