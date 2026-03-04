package hl7v27

// QBP_Z93_QBP - Group struct
type QBP_Z93_QBP struct {
	AnyHl7Segment Hxx `hl7:"TAG=Hxx;ATR=optional"`
}

// QBP_Z93 - Tabular Dispense History
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/QBP_Z93
type QBP_Z93 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Qbp QBP_Z93_QBP `hl7:"GROUP;ATR=optional"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

