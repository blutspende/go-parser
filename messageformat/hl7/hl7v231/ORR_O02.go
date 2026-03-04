package hl7v231

// ORR_O02_RESPONSE - Group struct
type ORR_O02_RESPONSE struct {
	Patient ORR_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORR_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORR_O02_RESPONSE_PATIENT - Group struct
type ORR_O02_RESPONSE_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORR_O02_RESPONSE_ORDER - Group struct
type ORR_O02_RESPONSE_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetailSegment ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT - Group struct
type ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORR_O02 - Order response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORR_O02
type ORR_O02 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORR_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

