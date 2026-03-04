package hl7v231

// SQM_S25_REQUEST - Group struct
type SQM_S25_REQUEST struct {
	AppointmentRequestSegment ARQ `hl7:"TAG=ARQ"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
	PatientIdentificationSegment PID `hl7:"TAG=PID;ATR=optional"`
	Resources []SQM_S25_REQUEST_RESOURCES `hl7:"GROUP"`
}

// SQM_S25_REQUEST_RESOURCES - Group struct
type SQM_S25_REQUEST_RESOURCES struct {
	ResourceGroupSegment RGS `hl7:"TAG=RGS"`
	Service []SQM_S25_REQUEST_RESOURCES_SERVICE `hl7:"GROUP;ATR=optional"`
	GeneralResource []SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE `hl7:"GROUP;ATR=optional"`
	PersonnelResource []SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE `hl7:"GROUP;ATR=optional"`
	LocationResource []SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE `hl7:"GROUP;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_SERVICE - Group struct
type SQM_S25_REQUEST_RESOURCES_SERVICE struct {
	AppointmentInformationServiceSegment AIS `hl7:"TAG=AIS"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_GENERAL_RESOURCE struct {
	AppointmentInformationGeneralResourceSegment AIG `hl7:"TAG=AIG"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_PERSONNEL_RESOURCE struct {
	AppointmentInformationPersonnelResourceSegment AIP `hl7:"TAG=AIP"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE - Group struct
type SQM_S25_REQUEST_RESOURCES_LOCATION_RESOURCE struct {
	AppointmentInformationLocationResourceSegment AIL `hl7:"TAG=AIL"`
	AppointmentPreferencesSegment APR `hl7:"TAG=APR;ATR=optional"`
}

// SQM_S25 - Schedule query message and response
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/SQM_S25
type SQM_S25 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	OriginalStyleQueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	Request SQM_S25_REQUEST `hl7:"GROUP;ATR=optional"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

