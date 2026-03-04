package hl7v28

// SRM_S11_PATIENT - Group struct
type SRM_S11_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Observation []SRM_S11_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
}

// SRM_S11_PATIENT_OBSERVATION - Group struct
type SRM_S11_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// SRM_S11_RESOURCES - Group struct
type SRM_S11_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SRM_S11_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	General_resource []SRM_S11_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	Location_resource []SRM_S11_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
	Personnel_resource []SRM_S11_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SRM_S11_RESOURCES_SERVICE - Group struct
type SRM_S11_RESOURCES_SERVICE struct {
	AppointmentInformation AIS `hl7:"TAG=AIS"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S11_RESOURCES_GENERAL_RESOURCE - Group struct
type SRM_S11_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S11_RESOURCES_LOCATION_RESOURCE - Group struct
type SRM_S11_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S11_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SRM_S11_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// SRM_S11 - Request deletion of service/resource on appointment
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/SRM_S11
type SRM_S11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	AppointmentRequest ARQ `hl7:"TAG=ARQ"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient []SRM_S11_PATIENT `hl7:"GROUP;ATR=optional"`
	Resources []SRM_S11_RESOURCES `hl7:"GROUP"`
}

