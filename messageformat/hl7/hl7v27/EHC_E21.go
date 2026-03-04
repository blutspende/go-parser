package hl7v27

// EHC_E21_AUTHORIZATION_REQUEST - Group struct
type EHC_E21_AUTHORIZATION_REQUEST struct {
	InvoiceSegment IVC `hl7:"TAG=IVC"`
	PslItemInfo []EHC_E21_AUTHORIZATION_REQUEST_PSL_ITEM_INFO `hl7:"GROUP"`
}

// EHC_E21_AUTHORIZATION_REQUEST_PSL_ITEM_INFO - Group struct
type EHC_E21_AUTHORIZATION_REQUEST_PSL_ITEM_INFO struct {
	ProductServiceLineItem PSL `hl7:"TAG=PSL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	AuthorizationInformation AUT `hl7:"TAG=AUT;ATR=optional"`
}

// EHC_E21 - Cancel Authorization Request
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/EHC_E21
type EHC_E21 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	AuthorizationRequest EHC_E21_AUTHORIZATION_REQUEST `hl7:"GROUP"`
}

