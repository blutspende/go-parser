package hl7v26

// QBP_Q13 - Query by Parameter Requesting an RTB Tabular Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/QBP_Q13
type QBP_Q13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	TableRowDefinition RDF `hl7:"TAG=RDF;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

