package hl7v28

// BRP_O30_RESPONSE - Group struct
type BRP_O30_RESPONSE struct {
	Patient BRP_O30_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
}

// BRP_O30_RESPONSE_PATIENT - Group struct
type BRP_O30_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Order []BRP_O30_RESPONSE_PATIENT_ORDER `hl7:"GROUP;ATR=optional"`
}

// BRP_O30_RESPONSE_PATIENT_ORDER - Group struct
type BRP_O30_RESPONSE_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []BRP_O30_RESPONSE_PATIENT_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO;ATR=optional"`
	BloodProductDispenseStatus []BPX `hl7:"TAG=BPX;ATR=optional"`
}

// BRP_O30_RESPONSE_PATIENT_ORDER_TIMING - Group struct
type BRP_O30_RESPONSE_PATIENT_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// BRP_O30 - Blood product dispense status acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/BRP_O30
type BRP_O30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response BRP_O30_RESPONSE `hl7:"GROUP;ATR=optional"`
}

