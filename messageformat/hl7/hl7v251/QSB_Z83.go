package hl7v251

// QSB_Z83 - Publish ORU Subscription
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QSB_Z83
type QSB_Z83 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

