package hl7v26

// ADT_A60 - Update Adverse Reaction Information 
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/ADT_A60
type ADT_A60 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	AccessRestrictions1 []ARV `hl7:"TAG=ARV;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	AccessRestrictions2 []ARV `hl7:"TAG=ARV;ATR=optional"`
	PatientAdverseReactionInformation []IAM `hl7:"TAG=IAM;ATR=optional"`
}

