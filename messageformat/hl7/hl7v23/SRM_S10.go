package hl7v23

// SRM_S10_PATIENT - Group struct
type SRM_S10_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SRM_S10_RESOURCES - Group struct
type SRM_S10_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SRM_S10_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SRM_S10_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SRM_S10_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SRM_S10_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SRM_S10_RESOURCES_SERVICE - Group struct
type SRM_S10_RESOURCES_SERVICE struct {
	AppointmentInformationService AIS `hl7:"TAG=AIS"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_GENERAL_RESOURCE - Group struct
type SRM_S10_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_LOCATION_RESOURCE - Group struct
type SRM_S10_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocation AIL `hl7:"TAG=AIL"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SRM_S10_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S10 - Request discontinuation of service/resource on appointment
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/SRM_S10
type SRM_S10 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	AppointmentRequest ARQ `hl7:"TAG=ARQ"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SRM_S10_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SRM_S10_RESOURCES `hl7:"GROUP"`
}

