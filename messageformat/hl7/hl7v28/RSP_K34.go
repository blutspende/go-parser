package hl7v28

// RSP_K34_DONOR - Group struct
type RSP_K34_DONOR struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Donor_registration RSP_K34_DONOR_DONOR_REGISTRATION `hl7:"GROUP;ATR=optional"`
}

// RSP_K34_DONOR_DONOR_REGISTRATION - Group struct
type RSP_K34_DONOR_DONOR_REGISTRATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_K34_DONATION - Group struct
type RSP_K34_DONATION struct {
	DonationSegment DON `hl7:"TAG=DON"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RSP_K34 - Get Donor Record Response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_K34
type RSP_K34 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	StaffIdentification []STF `hl7:"TAG=STF;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Donor RSP_K34_DONOR `hl7:"GROUP;ATR=optional"`
	Donation RSP_K34_DONATION `hl7:"GROUP;ATR=optional"`
}

