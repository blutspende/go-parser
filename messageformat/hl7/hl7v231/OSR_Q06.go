package hl7v231

// OSR_Q06_RESPONSE - Group struct
type OSR_Q06_RESPONSE struct {
	Patient OSR_Q06_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OSR_Q06_RESPONSE_ORDER `hl7:"GROUP"`
}

// OSR_Q06_RESPONSE_PATIENT - Group struct
type OSR_Q06_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OSR_Q06_RESPONSE_ORDER - Group struct
type OSR_Q06_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetailSegment OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT - Group struct
type OSR_Q06_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// OSR_Q06 - Query response for order status
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/OSR_Q06
type OSR_Q06 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Response OSR_Q06_RESPONSE `hl7:"GROUP;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

