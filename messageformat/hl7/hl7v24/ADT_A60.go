package hl7v24

// ADT_A60 - Update allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ADT_A60
type ADT_A60 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	PatientAdverseReactionInformationUniqueIden []IAM `hl7:"TAG=IAM;ATR=optional"`
}

