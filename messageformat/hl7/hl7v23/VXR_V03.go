package hl7v23

// VXR_V03_PATIENT_VISIT - Group struct
type VXR_V03_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXR_V03_INSURANCE - Group struct
type VXR_V03_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXR_V03_ORDER - Group struct
type VXR_V03_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	PharmacyAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXR_V03_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXR_V03_ORDER_OBSERVATION - Group struct
type VXR_V03_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXR_V03 - Vaccination record response
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/VXR_V03
type VXR_V03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit VXR_V03_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []VXR_V03_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXR_V03_ORDER `hl7:"GROUP;ATR=optional"`
}

