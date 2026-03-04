package hl7v27

// SIU_S16_PATIENT - Group struct
type SIU_S16_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SIU_S16_RESOURCES - Group struct
type SIU_S16_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SIU_S16_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	General_resource []SIU_S16_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	Location_resource []SIU_S16_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	Personnel_resource []SIU_S16_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SIU_S16_RESOURCES_SERVICE - Group struct
type SIU_S16_RESOURCES_SERVICE struct {
	AppointmentInformation AIS `hl7:"TAG=AIS"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S16_RESOURCES_GENERAL_RESOURCE - Group struct
type SIU_S16_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S16_RESOURCES_LOCATION_RESOURCE - Group struct
type SIU_S16_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S16_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SIU_S16_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SIU_S16 - Notification of appointment discontinuation
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/SIU_S16
type SIU_S16 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SchedulingActivityInformation SCH `hl7:"TAG=SCH"`
	TimingQuantity []TQ1 `hl7:"TAG=TQ1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SIU_S16_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SIU_S16_RESOURCES `hl7:"GROUP"`
}

