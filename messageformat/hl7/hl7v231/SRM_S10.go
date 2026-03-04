package hl7v231

// SRM_S10_PATIENT - Group struct
type SRM_S10_PATIENT struct {
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SRM_S10_RESOURCES - Group struct
type SRM_S10_RESOURCES struct {
	ResourceGroupSegment RGS `hl7:"TAG=RGS"`
	Service []SRM_S10_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SRM_S10_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SRM_S10_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SRM_S10_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SRM_S10_RESOURCES_SERVICE - Group struct
type SRM_S10_RESOURCES_SERVICE struct {
	AppointmentInformationServiceSegment AIS `hl7:"TAG=AIS"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_GENERAL_RESOURCE - Group struct
type SRM_S10_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResourceSegment AIG `hl7:"TAG=AIG"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_LOCATION_RESOURCE - Group struct
type SRM_S10_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResourceSegment AIL `hl7:"TAG=AIL"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SRM_S10_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResourceSegment AIP `hl7:"TAG=AIP"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10 - Request discontinuation of service/resource on appointment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/SRM_S10
type SRM_S10 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	AppointmentRequestSegment ARQ `hl7:"TAG=ARQ"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SRM_S10_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SRM_S10_RESOURCES `hl7:"GROUP"`
}

