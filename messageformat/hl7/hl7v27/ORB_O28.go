package hl7v27

// ORB_O28_RESPONSE - Group struct
type ORB_O28_RESPONSE struct {
	Patient ORB_O28_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
}

// ORB_O28_RESPONSE_PATIENT - Group struct
type ORB_O28_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	Order []ORB_O28_RESPONSE_PATIENT_ORDER `hl7:"GROUP;ATR=optional"`
}

// ORB_O28_RESPONSE_PATIENT_ORDER - Group struct
type ORB_O28_RESPONSE_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []ORB_O28_RESPONSE_PATIENT_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO;ATR=optional"`
}

// ORB_O28_RESPONSE_PATIENT_ORDER_TIMING - Group struct
type ORB_O28_RESPONSE_PATIENT_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORB_O28 - Blood product order acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/ORB_O28
type ORB_O28 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORB_O28_RESPONSE `hl7:"GROUP;ATR=optional"`
}

