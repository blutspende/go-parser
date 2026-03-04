package hl7v27

// SRR_S03_SCHEDULE - Group struct
type SRR_S03_SCHEDULE struct {
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	TimingQuantity []TQ1 `hl7:"TAG=TQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
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
	AppointmentInformation AIS `hl7:"TAG=AIS"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_GENERAL_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_LOCATION_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03_SCHEDULE_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SRR_S03_SCHEDULE_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRR_S03 - Scheduled Request Response - Request Appointment Modification 
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/SRR_S03
type SRR_S03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	Schedule SRR_S03_SCHEDULE `hl7:"GROUP;ATR=optional"`
}

