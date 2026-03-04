package hl7v22

// ORR_O02_PATIENT - Group struct
type ORR_O02_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Order []ORR_O02_PATIENT_ORDER `hl7:"GROUP"`
}

// ORR_O02_PATIENT_ORDER - Group struct
type ORR_O02_PATIENT_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail ORR_O02_PATIENT_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// ORR_O02_PATIENT_ORDER_ORDER_DETAIL - Group struct
type ORR_O02_PATIENT_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment ORR_O02_PATIENT_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORR_O02_PATIENT_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type ORR_O02_PATIENT_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstruction ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORR_O02 - Response Message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ORR_O02
type ORR_O02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient ORR_O02_PATIENT `hl7:"GROUP;ATR=optional"`
}

