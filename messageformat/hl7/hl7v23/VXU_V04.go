package hl7v23

// VXU_V04_PATIENT - Group struct
type VXU_V04_PATIENT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
}

// VXU_V04_INSURANCE - Group struct
type VXU_V04_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// VXU_V04_ORDER - Group struct
type VXU_V04_ORDER struct {
	CommonOrderSegment ORC `hl7:"TAG=ORC;ATR=optional"`
	PharmacyAdministrationSegment RXA `hl7:"TAG=RXA"`
	PharmacyRouteSegment RXR `hl7:"TAG=RXR;ATR=optional"`
	Observation []VXU_V04_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// VXU_V04_ORDER_OBSERVATION - Group struct
type VXU_V04_ORDER_OBSERVATION struct {
	ObservationSegment OBX `hl7:"TAG=OBX"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// VXU_V04 - Unsolicited vaccination record update
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/VXU_V04
type VXU_V04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Patient VXU_V04_PATIENT `hl7:"GROUP;ATR=optional"`
	Insurance []VXU_V04_INSURANCE `hl7:"GROUP;ATR=optional"`
	Order []VXU_V04_ORDER `hl7:"GROUP;ATR=optional"`
}

