package hl7v28

// DRC_O47_DONOR - Group struct
type DRC_O47_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration DRC_O47_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DRC_O47_DONOR_DONOR_REGISTRATION - Group struct
type DRC_O47_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DRC_O47_DONOR_ORDER - Group struct
type DRC_O47_DONOR_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DRC_O47 - Donor Request to Collect
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DRC_O47
type DRC_O47 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DRC_O47_DONOR `hl7:"GROUP;ATR=optional"`
	Donor_order []DRC_O47_DONOR_ORDER `hl7:"GROUP"`
}

