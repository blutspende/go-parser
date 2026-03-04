package hl7v23

// SIU_S21_PATIENT - Group struct
type SIU_S21_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SIU_S21_RESOURCES - Group struct
type SIU_S21_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SIU_S21_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SIU_S21_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SIU_S21_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SIU_S21_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SIU_S21_RESOURCES_SERVICE - Group struct
type SIU_S21_RESOURCES_SERVICE struct {
	AppointmentInformationService AIS `hl7:"TAG=AIS"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S21_RESOURCES_GENERAL_RESOURCE - Group struct
type SIU_S21_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S21_RESOURCES_LOCATION_RESOURCE - Group struct
type SIU_S21_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocation AIL `hl7:"TAG=AIL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S21_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SIU_S21_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S21 - Notification of discontinuation of service/resource on appointment
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/SIU_S21
type SIU_S21 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	ScheduleActivityInformation SCH `hl7:"TAG=SCH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SIU_S21_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SIU_S21_RESOURCES `hl7:"GROUP"`
}

