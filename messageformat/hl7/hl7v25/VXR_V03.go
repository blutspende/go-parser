package hl7v25

// VXR_V03_PATIENT_VISIT - Group struct
type VXR_V03_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXR_V03_INSURANCE - Group struct
type VXR_V03_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXR_V03_ORDER - Group struct
type VXR_V03_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Timing []VXR_V03_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	PharmacyTreatmentAdministration RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXR_V03_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXR_V03_ORDER_TIMING - Group struct
type VXR_V03_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// VXR_V03_ORDER_OBSERVATION - Group struct
type VXR_V03_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXR_V03 - Vaccination Record Response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/VXR_V03
type VXR_V03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit VXR_V03_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []VXR_V03_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXR_V03_ORDER `hl7:"GROUP;ATR=optional"`
}

