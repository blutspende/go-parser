package hl7v271

// QBP_Z85_QBP - Group struct
type QBP_Z85_QBP struct {
	AnyHl7Segment Hxx `hl7:"TAG=Hxx;ATR=optional"`
}

// QBP_Z85 - Pharmacy Information Comprehensive
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/QBP_Z85
type QBP_Z85 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Qbp QBP_Z85_QBP `hl7:"GROUP;ATR=optional"`
	ResponseControlParameter RCP `hl7:"TAG=RCP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

