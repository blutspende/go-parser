package hl7v271

// RSP_E22_QUERY_ACK - Group struct
type RSP_E22_QUERY_ACK struct {
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	AuthorizationInfo RSP_E22_QUERY_ACK_AUTHORIZATION_INFO `hl7:"GROUP;ATR=optional"`
}

// RSP_E22_QUERY_ACK_AUTHORIZATION_INFO - Group struct
type RSP_E22_QUERY_ACK_AUTHORIZATION_INFO struct {
	InvoiceSegment IVC `hl7:"TAG=IVC"`
	ProductServiceGroup PSG `hl7:"TAG=PSG"`
	PslItemInfo []RSP_E22_QUERY_ACK_AUTHORIZATION_INFO_PSL_ITEM_INFO `hl7:"GROUP"`
}

// RSP_E22_QUERY_ACK_AUTHORIZATION_INFO_PSL_ITEM_INFO - Group struct
type RSP_E22_QUERY_ACK_AUTHORIZATION_INFO_PSL_ITEM_INFO struct {
	ProductServiceLineItem PSL `hl7:"TAG=PSL"`
}

// RSP_E22 - Authorization Request StatusResponse 
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/RSP_E22
type RSP_E22 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAck RSP_E22_QUERY_ACK `hl7:"GROUP"`
}

