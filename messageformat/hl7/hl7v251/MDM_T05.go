package hl7v251

// MDM_T05_COMMON_ORDER - Group struct
type MDM_T05_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []MDM_T05_COMMON_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// MDM_T05_COMMON_ORDER_TIMING - Group struct
type MDM_T05_COMMON_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// MDM_T05 - Document addendum notification
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/MDM_T05
type MDM_T05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	CommonOrder []MDM_T05_COMMON_ORDER `hl7:"GROUP;ATR=optional"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA"`
}

