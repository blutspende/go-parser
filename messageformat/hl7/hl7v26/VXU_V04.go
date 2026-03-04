package hl7v26

// VXU_V04_PATIENT - Group struct
type VXU_V04_PATIENT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXU_V04_INSURANCE - Group struct
type VXU_V04_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXU_V04_ORDER - Group struct
type VXU_V04_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []VXU_V04_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentAdministration RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXU_V04_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXU_V04_ORDER_TIMING - Group struct
type VXU_V04_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// VXU_V04_ORDER_OBSERVATION - Group struct
type VXU_V04_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXU_V04 - Unsolicited Vaccination Record Update
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/VXU_V04
type VXU_V04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Patient VXU_V04_PATIENT `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []VXU_V04_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXU_V04_ORDER `hl7:"GROUP;ATR=optional"`
}

