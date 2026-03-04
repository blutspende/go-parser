package hl7v251

// ORM_O01_PATIENT - Group struct
type ORM_O01_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit ORM_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []ORM_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// ORM_O01_PATIENT_PATIENT_VISIT - Group struct
type ORM_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORM_O01_PATIENT_INSURANCE - Group struct
type ORM_O01_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ORM_O01_ORDER - Group struct
type ORM_O01_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail ORM_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL - Group struct
type ORM_O01_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequest OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferences ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructions ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORM_O01 - General Order
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient ORM_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORM_O01_ORDER `hl7:"GROUP"`
}

