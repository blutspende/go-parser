package hl7v26

// OSR_Q06_RESPONSE - Group struct
type OSR_Q06_RESPONSE struct {
	Patient OSR_Q06_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OSR_Q06_RESPONSE_ORDER `hl7:"GROUP"`
}

// OSR_Q06_RESPONSE_PATIENT - Group struct
type OSR_Q06_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OSR_Q06_RESPONSE_ORDER - Group struct
type OSR_Q06_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []OSR_Q06_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetailSegment OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// OSR_Q06_RESPONSE_ORDER_TIMING - Group struct
type OSR_Q06_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT - Group struct
type OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// OSR_Q06 - Query Response for Order Status
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/OSR_Q06
type OSR_Q06 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Response OSR_Q06_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

