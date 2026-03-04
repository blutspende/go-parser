package hl7v231

// ORM_O01_PATIENT - Group struct
type ORM_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit ORM_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []ORM_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	GuarantorSegment GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// ORM_O01_PATIENT_PATIENT_VISIT - Group struct
type ORM_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// ORM_O01_PATIENT_INSURANCE - Group struct
type ORM_O01_PATIENT_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment []IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ORM_O01_ORDER - Group struct
type ORM_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail ORM_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	ClinicalTrialIdentificationSegment []CTI `hl7:"TAG=CTI;ATR=optional"`
	BillingSegment BLG `hl7:"TAG=BLG;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL - Group struct
type ORM_O01_ORDER_ORDER_DETAIL struct {
	OrderDetailSegment ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT `hl7:"GROUP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_ORDER_DETAIL_SEGMENT struct {
	ObservationRequestSegment OBR `hl7:"TAG=OBR;ATR=optional"`
	RequisitionDetailSegment RQD `hl7:"TAG=RQD;ATR=optional"`
	RequisitionDetail1Segment RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	PharmacyTreatmentOrderSegment RXO `hl7:"TAG=RXO;ATR=optional"`
	DietaryOrdersSupplementsAndPreferencesSegment ODS `hl7:"TAG=ODS;ATR=optional"`
	DietTrayInstructionsSegment ODT `hl7:"TAG=ODT;ATR=optional"`
}

// ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION - Group struct
type ORM_O01_ORDER_ORDER_DETAIL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORM_O01 - Order message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient ORM_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORM_O01_ORDER `hl7:"GROUP"`
}

