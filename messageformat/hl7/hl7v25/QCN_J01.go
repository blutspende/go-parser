package hl7v25

// QCN_J01 - Cancel Query
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/QCN_J01
type QCN_J01 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryIdentification QID `hl7:"TAG=QID"`
}

