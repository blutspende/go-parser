package hl7v28

// DER_O44_DONOR - Group struct
type DER_O44_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration DER_O44_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DER_O44_DONOR_DONOR_REGISTRATION - Group struct
type DER_O44_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DER_O44_DONOR_ORDER - Group struct
type DER_O44_DONOR_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DER_O44 - Donor Eligibility Request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DER_O44
type DER_O44 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DER_O44_DONOR `hl7:"GROUP;ATR=optional"`
	Donor_order []DER_O44_DONOR_ORDER `hl7:"GROUP"`
}

