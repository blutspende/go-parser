package hl7v28

// RDE_O11_PATIENT - Group struct
type RDE_O11_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Patient_visit RDE_O11_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []RDE_O11_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// RDE_O11_PATIENT_PATIENT_VISIT - Group struct
type RDE_O11_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
}

// RDE_O11_PATIENT_INSURANCE - Group struct
type RDE_O11_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RDE_O11_ORDER - Group struct
type RDE_O11_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation1 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []RDE_O11_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Order_detail RDE_O11_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentEncodedOrder RXE `hl7:"TAG=RXE"`
	ParticipationInformation2 []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Timing_encoded []RDE_O11_ORDER_TIMING_ENCODED `hl7:"GROUP"`
	Infusion_order RDE_O11_ORDER_INFUSION_ORDER `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	CumulativeDosage []CDO `hl7:"TAG=CDO;ATR=optional"`
	Observation []RDE_O11_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
}

// RDE_O11_ORDER_TIMING - Group struct
type RDE_O11_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDE_O11_ORDER_ORDER_DETAIL - Group struct
type RDE_O11_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Component []RDE_O11_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// RDE_O11_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type RDE_O11_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O11_ORDER_TIMING_ENCODED - Group struct
type RDE_O11_ORDER_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDE_O11_ORDER_INFUSION_ORDER - Group struct
type RDE_O11_ORDER_INFUSION_ORDER struct {
	PharmacyTreatmentInfusion RXV `hl7:"TAG=RXV"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Timing_encoded []RDE_O11_ORDER_INFUSION_ORDER_TIMING_ENCODED `hl7:"GROUP"`
}

// RDE_O11_ORDER_INFUSION_ORDER_TIMING_ENCODED - Group struct
type RDE_O11_ORDER_INFUSION_ORDER_TIMING_ENCODED struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// RDE_O11_ORDER_OBSERVATION - Group struct
type RDE_O11_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RDE_O11 - Pharmacy/treatment encoded order
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RDE_O11
type RDE_O11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient RDE_O11_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RDE_O11_ORDER `hl7:"GROUP"`
}

