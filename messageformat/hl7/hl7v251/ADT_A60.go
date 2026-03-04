package hl7v251

// ADT_A60 - Update Adverse Reaction Information
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ADT_A60
type ADT_A60 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	PatientAdverseReactionInformation []IAM `hl7:"TAG=IAM;ATR=optional"`
}

