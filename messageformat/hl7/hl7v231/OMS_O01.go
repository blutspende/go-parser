package hl7v231

// OMS_O01_PATIENT - Group struct
type OMS_O01_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit OMS_O01_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMS_O01_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	GuarantorSegment GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformationSegment []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMS_O01_PATIENT_PATIENT_VISIT - Group struct
type OMS_O01_PATIENT_PATIENT_VISIT struct {
	PatientVisitSegment PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OMS_O01_PATIENT_INSURANCE - Group struct
type OMS_O01_PATIENT_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMS_O01_ORDER - Group struct
type OMS_O01_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC"`
	OrderDetail OMS_O01_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	BillingSegment BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OMS_O01_ORDER_ORDER_DETAIL - Group struct
type OMS_O01_ORDER_ORDER_DETAIL struct {
	RequisitionDetailSegment RQD `hl7:"TAG=RQD"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []OMS_O01_ORDER_ORDER_DETAIL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OMS_O01_ORDER_ORDER_DETAIL_OBSERVATION - Group struct
type OMS_O01_ORDER_ORDER_DETAIL_OBSERVATION struct {
	ObservationResultSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMS_O01 - Stock requisition order message
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/OMS_O01
type OMS_O01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMS_O01_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OMS_O01_ORDER `hl7:"GROUP"`
}

