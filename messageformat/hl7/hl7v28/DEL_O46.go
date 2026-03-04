package hl7v28

// DEL_O46_DONOR - Group struct
type DEL_O46_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration DEL_O46_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DEL_O46_DONOR_DONOR_REGISTRATION - Group struct
type DEL_O46_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DEL_O46 - Donor Eligibility
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DEL_O46
type DEL_O46 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DEL_O46_DONOR `hl7:"GROUP;ATR=optional"`
	DonationSegment DON `hl7:"TAG=DON"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

