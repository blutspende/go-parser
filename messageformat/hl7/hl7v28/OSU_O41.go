package hl7v28

// OSU_O41_ORDER_STATUS - Group struct
type OSU_O41_ORDER_STATUS struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSU_O41 - Order Status Update
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/OSU_O41
type OSU_O41 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Order_status []OSU_O41_ORDER_STATUS `hl7:"GROUP"`
}

