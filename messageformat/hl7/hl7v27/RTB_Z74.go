package hl7v27

// RTB_Z74_ROW_DEFINITION - Group struct
type RTB_Z74_ROW_DEFINITION struct {
	TableRowDefinition RDF `hl7:"TAG=RDF"`
	TableRowData []RDT `hl7:"TAG=RDT;ATR=optional"`
}

// RTB_Z74 - Information about Phone Calls (Response)
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/RTB_Z74
type RTB_Z74 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Row_definition RTB_Z74_ROW_DEFINITION `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

