package hl7v24

// RRA_O18_RESPONSE - Group struct
type RRA_O18_RESPONSE struct {
	Patient RRA_O18_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []RRA_O18_RESPONSE_ORDER `hl7:"GROUP"`
}

// RRA_O18_RESPONSE_PATIENT - Group struct
type RRA_O18_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RRA_O18_RESPONSE_ORDER - Group struct
type RRA_O18_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	Administration RRA_O18_RESPONSE_ORDER_ADMINISTRATION `hl7:"GROUP;ATR=optional"`
}

// RRA_O18_RESPONSE_ORDER_ADMINISTRATION - Group struct
type RRA_O18_RESPONSE_ORDER_ADMINISTRATION struct {
	PharmacyTreatmentAdministration []RXA `hl7:"TAG=RXA"`
	PharmacyTreatmentRoute RXR `hl7:"TAG=RXR"`
}

// RRA_O18 - Pharmacy/treatment administration acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/RRA_O18
type RRA_O18 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response RRA_O18_RESPONSE `hl7:"GROUP;ATR=optional"`
}

