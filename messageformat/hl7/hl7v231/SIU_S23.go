package hl7v231

// SIU_S23_PATIENT - Group struct
type SIU_S23_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SIU_S23_RESOURCES - Group struct
type SIU_S23_RESOURCES struct {
	ResourceGroupSegment RGS `hl7:"TAG=RGS"`
	Service []SIU_S23_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SIU_S23_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SIU_S23_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SIU_S23_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SIU_S23_RESOURCES_SERVICE - Group struct
type SIU_S23_RESOURCES_SERVICE struct {
	AppointmentInformationServiceSegment AIS `hl7:"TAG=AIS"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S23_RESOURCES_GENERAL_RESOURCE - Group struct
type SIU_S23_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResourceSegment AIG `hl7:"TAG=AIG"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S23_RESOURCES_LOCATION_RESOURCE - Group struct
type SIU_S23_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResourceSegment AIL `hl7:"TAG=AIL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S23_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SIU_S23_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResourceSegment AIP `hl7:"TAG=AIP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S23 - Notification of blocked schedule time slot(s)
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/SIU_S23
type SIU_S23 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ScheduleActivityInformationSegment SCH `hl7:"TAG=SCH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SIU_S23_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SIU_S23_RESOURCES `hl7:"GROUP"`
}

