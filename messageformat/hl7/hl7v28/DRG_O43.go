package hl7v28

// DRG_O43_DONOR - Group struct
type DRG_O43_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration DRG_O43_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DRG_O43_DONOR_DONOR_REGISTRATION - Group struct
type DRG_O43_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DRG_O43 - Donor Registration
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DRG_O43
type DRG_O43 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DRG_O43_DONOR `hl7:"GROUP;ATR=optional"`
}

