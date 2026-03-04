package hl7v28

// DEO_O45_DONOR - Group struct
type DEO_O45_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Donor_registration DEO_O45_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DEO_O45_DONOR_DONOR_REGISTRATION - Group struct
type DEO_O45_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DEO_O45_DONOR_ORDER - Group struct
type DEO_O45_DONOR_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Donation_observation []DEO_O45_DONOR_ORDER_DONATION_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// DEO_O45_DONOR_ORDER_DONATION_OBSERVATION - Group struct
type DEO_O45_DONOR_ORDER_DONATION_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DEO_O45 - Donor Eligibility Observations
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DEO_O45
type DEO_O45 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DEO_O45_DONOR `hl7:"GROUP;ATR=optional"`
	Donor_order []DEO_O45_DONOR_ORDER `hl7:"GROUP"`
}

