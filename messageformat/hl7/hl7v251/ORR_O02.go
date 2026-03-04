package hl7v251

// ORR_O02_RESPONSE - Group struct
type ORR_O02_RESPONSE struct {
	Patient ORR_O02_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORR_O02_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORR_O02_RESPONSE_PATIENT - Group struct
type ORR_O02_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORR_O02_RESPONSE_ORDER - Group struct
type ORR_O02_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetailSegment ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT - Group struct
type ORR_O02_RESPONSE_ORDER_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORR_O02 - General Order Response
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ORR_O02
type ORR_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORR_O02_RESPONSE `hl7:"GROUP;ATR=optional"`
}

