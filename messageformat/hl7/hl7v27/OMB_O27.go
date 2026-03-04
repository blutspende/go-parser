package hl7v27

// OMB_O27_PATIENT - Group struct
type OMB_O27_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient_visit OMB_O27_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OMB_O27_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OMB_O27_PATIENT_PATIENT_VISIT - Group struct
type OMB_O27_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// OMB_O27_PATIENT_INSURANCE - Group struct
type OMB_O27_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OMB_O27_ORDER - Group struct
type OMB_O27_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []OMB_O27_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	BloodProductOrder BPO `hl7:"TAG=BPO"`
	Specimen SPM `hl7:"TAG=SPM;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []OMB_O27_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OMB_O27_ORDER_TIMING - Group struct
type OMB_O27_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OMB_O27_ORDER_OBSERVATION - Group struct
type OMB_O27_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OMB_O27 - Blood product order
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/OMB_O27
type OMB_O27 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OMB_O27_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OMB_O27_ORDER `hl7:"GROUP"`
}

