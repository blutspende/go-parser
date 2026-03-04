package hl7v23

// SRR_S03_SCHEDULE - Group struct
type SRR_S03_SCHEDULE struct {
	ScheduleActivityInformation SCH `hl7:"TAG=SCH"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SRR_S03_SCHEDULE_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SRR_S03_SCHEDULE_RESOURCES `hl7:"GROUP"`
}

// SRR_S03_SCHEDULE_PATIENT - Group struct
type SRR_S03_SCHEDULE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES - Group struct
type SRR_S03_SCHEDULE_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SRR_S03_SCHEDULE_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SRR_S03_SCHEDULE_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SRR_S03_SCHEDULE_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SRR_S03_SCHEDULE_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_SERVICE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_SERVICE struct {
	AppointmentInformationService AIS `hl7:"TAG=AIS"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_GENERAL_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_LOCATION_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocation AIL `hl7:"TAG=AIL"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03 - Response appointment modification
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/SRR_S03
type SRR_S03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	Schedule SRR_S03_SCHEDULE `hl7:"GROUP;ATR=optional"`
}

