package hl7v271

// OMS_O05_PATIENT - Group struct
type OMS_O05_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient_visit OMS_O05_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMS_O05_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMS_O05_PATIENT_PATIENT_VISIT - Group struct
type OMS_O05_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OMS_O05_PATIENT_INSURANCE - Group struct
type OMS_O05_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMS_O05_ORDER - Group struct
type OMS_O05_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []OMS_O05_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	RequisitionDetail RQD `hl7:"TAG=RQD"`
	RequisitionDetail1 RQ1 `hl7:"TAG=RQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Observation []OMS_O05_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OMS_O05_ORDER_TIMING - Group struct
type OMS_O05_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMS_O05_ORDER_OBSERVATION - Group struct
type OMS_O05_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMS_O05 - Stock requisition order
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/OMS_O05
type OMS_O05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMS_O05_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OMS_O05_ORDER `hl7:"GROUP"`
}

