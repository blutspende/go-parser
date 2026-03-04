package hl7v28

// RSP_E03_QUERY_ACK_IPR - Group struct
type RSP_E03_QUERY_ACK_IPR struct {
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	InvoiceProcessingResults []IPR `hl7:"TAG=IPR;ATR=optional"`
}

// RSP_E03 - HealthCare Services Invoice Status Query Response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_E03
type RSP_E03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	Query_ack_ipr RSP_E03_QUERY_ACK_IPR `hl7:"GROUP;ATR=optional"`
}

