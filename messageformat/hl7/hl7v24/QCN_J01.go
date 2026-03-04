package hl7v24

// QCN_J01 - Cancel query/acknowledge message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/QCN_J01
type QCN_J01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

