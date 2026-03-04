package hl7v28

// DPR_O48_DONOR - Group struct
type DPR_O48_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration DPR_O48_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// DPR_O48_DONOR_DONOR_REGISTRATION - Group struct
type DPR_O48_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DPR_O48_DONOR_ORDER - Group struct
type DPR_O48_DONOR_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DPR_O48_DONATION - Group struct
type DPR_O48_DONATION struct {
	DonationSegment DON `hl7:"TAG=DON"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Blood_unit DPR_O48_DONATION_BLOOD_UNIT `hl7:"GROUP;ATR=optional"`
}

// DPR_O48_DONATION_BLOOD_UNIT - Group struct
type DPR_O48_DONATION_BLOOD_UNIT struct {
	BloodUnitInformationSegment []BUI `hl7:"TAG=BUI;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DPR_O48 - Donation Procedure
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/DPR_O48
type DPR_O48 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Donor DPR_O48_DONOR `hl7:"GROUP;ATR=optional"`
	Donor_order []DPR_O48_DONOR_ORDER `hl7:"GROUP"`
	Donation DPR_O48_DONATION `hl7:"GROUP;ATR=optional"`
}

