package hl7v24

// ADT_A10 - Patient arriving - tracking
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A10
type ADT_A10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

