package hl7v24

// ORP_O10_RESPONSE - Group struct
type ORP_O10_RESPONSE struct {
	Patient ORP_O10_RESPONSE_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []ORP_O10_RESPONSE_ORDER `hl7:"GROUP"`
}

// ORP_O10_RESPONSE_PATIENT - Group struct
type ORP_O10_RESPONSE_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORP_O10_RESPONSE_ORDER - Group struct
type ORP_O10_RESPONSE_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	OrderDetail ORP_O10_RESPONSE_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// ORP_O10_RESPONSE_ORDER_ORDER_DETAIL - Group struct
type ORP_O10_RESPONSE_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments1 []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	PharmacyTreatmentComponentOrder []RXC `hl7:"TAG=RXC;ATR=optional"`
	NotesAndComments2 []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORP_O10 - Pharmacy/treatment order acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORP_O10
type ORP_O10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORP_O10_RESPONSE `hl7:"GROUP;ATR=optional"`
}

