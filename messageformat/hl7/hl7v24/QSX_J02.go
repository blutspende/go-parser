package hl7v24

// QSX_J02 - Cancel subscription/acknowledge message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QSX_J02
type QSX_J02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

