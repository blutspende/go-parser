package hl7v22

// ORM_O01_PATIENT - Group struct
type ORM_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
}

// ORM_O01_ORDER - Group struct
type ORM_O01_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Order_detail ORM_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL - Group struct
type ORM_O01_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyPrescriptionOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstruction ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORM_O01 - Order Message
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient ORM_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORM_O01_ORDER `hl7:"GROUP"`
}

