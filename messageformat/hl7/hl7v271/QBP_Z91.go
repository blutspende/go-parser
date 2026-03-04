package hl7v271

// QBP_Z91_QBP - Group struct
type QBP_Z91_QBP struct {
	AnyHl7Segment Hxx `hl7:"TAG=Hxx;ATR=optional"`
}

// QBP_Z91 - Who Am I
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/QBP_Z91
type QBP_Z91 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Qbp QBP_Z91_QBP `hl7:"GROUP;ATR=optional"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

