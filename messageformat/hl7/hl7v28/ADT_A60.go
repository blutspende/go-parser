package hl7v28

// ADT_A60_ADVERSE_REACTION_GROUP - Group struct
type ADT_A60_ADVERSE_REACTION_GROUP struct {
	PatientAdverseReactionInformation IAM `hl7:"TAG=IAM"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	AllergyReaction []IAR `hl7:"TAG=IAR;ATR=optional"`
}

// ADT_A60 - Update allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/ADT_A60
type ADT_A60 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	AccessRestriction1 []ARV `hl7:"TAG=ARV;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	AccessRestriction2 []ARV `hl7:"TAG=ARV;ATR=optional"`
	Adverse_reaction_group []ADT_A60_ADVERSE_REACTION_GROUP `hl7:"GROUP;ATR=optional"`
}

