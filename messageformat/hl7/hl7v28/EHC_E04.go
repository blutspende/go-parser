package hl7v28

// EHC_E04_REASSESSMENT_REQUEST_INFO - Group struct
type EHC_E04_REASSESSMENT_REQUEST_INFO struct {
	InvoiceSegment []IVC `hl7:"TAG=IVC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Product_service_section []EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION `hl7:"GROUP;ATR=optional"`
}

// EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION - Group struct
type EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION struct {
	ProductServiceSection PSS `hl7:"TAG=PSS"`
	Product_service_group []EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION_PRODUCT_SERVICE_GROUP `hl7:"GROUP;ATR=optional"`
}

// EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION_PRODUCT_SERVICE_GROUP - Group struct
type EHC_E04_REASSESSMENT_REQUEST_INFO_PRODUCT_SERVICE_SECTION_PRODUCT_SERVICE_GROUP struct {
	ProductServiceGroup PSG `hl7:"TAG=PSG"`
	ProductServiceLineItem []PSL `hl7:"TAG=PSL;ATR=optional"`
}

// EHC_E04 - Assess HealthCare Services Invoice Request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/EHC_E04
type EHC_E04 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	Reassessment_request_info EHC_E04_REASSESSMENT_REQUEST_INFO `hl7:"GROUP"`
}

