package hl7v25

// SQM_S25_REQUEST - Group struct
type SQM_S25_REQUEST struct {
	AppointmentRequest ARQ `hl7:"TAG=ARQ"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	Resources []SQM_S25_REQUEST_RESOURCES `hl7:"GROUP"`
}

// SQM_S25_REQUEST_RESOURCES - Group struct
type SQM_S25_REQUEST_RESOURCES struct {
	ResourceGroup RGS `hl7:"TAG=RGS"`
	Service []SQM_S25_REQUEST_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_SERVICE - Group struct
type SQM_S25_REQUEST_RESOURCES_SERVICE struct {
	AppointmentInformation AIS `hl7:"TAG=AIS"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResource AIG `hl7:"TAG=AIG"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResource AIP `hl7:"TAG=AIP"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResource AIL `hl7:"TAG=AIL"`
	AppointmentPreferences APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25 - Schedule query message and response
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/SQM_S25
type SQM_S25 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Request SQM_S25_REQUEST `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

