package hl7v26

// RSP_E03_QUERY_ACK - Group struct
type RSP_E03_QUERY_ACK struct {
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	InvoiceProcessingResultsInfo []RSP_E03_QUERY_ACK_INVOICE_PROCESSING_RESULTS_INFO `hl7:"GROUP;ATR=optional"`
}

// RSP_E03_QUERY_ACK_INVOICE_PROCESSING_RESULTS_INFO - Group struct
type RSP_E03_QUERY_ACK_INVOICE_PROCESSING_RESULTS_INFO struct {
	InvoiceProcessingResults IPR `hl7:"TAG=IPR"`
}

// RSP_E03 - HealthCare Services Invoice Status Query Response
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/RSP_E03
type RSP_E03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential []UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAck RSP_E03_QUERY_ACK `hl7:"GROUP"`
}

