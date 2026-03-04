package hl7v27

// EHC_E24_AUTHORIZATION_RESPONSE_INFO - Group struct
type EHC_E24_AUTHORIZATION_RESPONSE_INFO struct {
	InvoiceSegment IVC `hl7:"TAG=IVC"`
	PslItemInfo []EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO `hl7:"GROUP"`
}

// EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO - Group struct
type EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO struct {
	ProductServiceLineItem PSL `hl7:"TAG=PSL"`
	AuthorizationInformation AUT `hl7:"TAG=AUT;ATR=optional"`
	PayerAdjustment []EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO_PAYER_ADJUSTMENT `hl7:"GROUP;ATR=optional"`
}

// EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO_PAYER_ADJUSTMENT - Group struct
type EHC_E24_AUTHORIZATION_RESPONSE_INFO_PSL_ITEM_INFO_PAYER_ADJUSTMENT struct {
	Adjustment ADJ `hl7:"TAG=ADJ"`
}

// EHC_E24 - Authorization Response
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/EHC_E24
type EHC_E24 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	AuthorizationResponseInfo EHC_E24_AUTHORIZATION_RESPONSE_INFO `hl7:"GROUP"`
}

