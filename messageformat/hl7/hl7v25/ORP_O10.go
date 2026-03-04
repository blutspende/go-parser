package hl7v25

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
	Timing []ORP_O10_RESPONSE_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	OrderDetail ORP_O10_RESPONSE_ORDER_ORDER_DETAIL `hl7:"GROUP;ATR=optional"`
}

// ORP_O10_RESPONSE_ORDER_TIMING - Group struct
type ORP_O10_RESPONSE_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// ORP_O10_RESPONSE_ORDER_ORDER_DETAIL - Group struct
type ORP_O10_RESPONSE_ORDER_ORDER_DETAIL struct {
	PharmacyTreatmentOrder RXO `hl7:"TAG=RXO"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	PharmacyTreatmentRoute []RXR `hl7:"TAG=RXR"`
	Component []ORP_O10_RESPONSE_ORDER_ORDER_DETAIL_COMPONENT `hl7:"GROUP;ATR=optional"`
}

// ORP_O10_RESPONSE_ORDER_ORDER_DETAIL_COMPONENT - Group struct
type ORP_O10_RESPONSE_ORDER_ORDER_DETAIL_COMPONENT struct {
	PharmacyTreatmentComponentOrder RXC `hl7:"TAG=RXC"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// ORP_O10 - Pharmacy/Treatment Order Acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/ORP_O10
type ORP_O10 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Response ORP_O10_RESPONSE `hl7:"GROUP;ATR=optional"`
}

