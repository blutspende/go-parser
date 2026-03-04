package hl7v28

// EHC_E02_INVOICE_INFORMATION_CANCEL - Group struct
type EHC_E02_INVOICE_INFORMATION_CANCEL struct {
	InvoiceSegment IVC `hl7:"TAG=IVC"`
	PayeeInformation PYE `hl7:"TAG=PYE"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Product_service_section []EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION `hl7:"GROUP;ATR=optional"`
}

// EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION - Group struct
type EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION struct {
	ProductServiceSection PSS `hl7:"TAG=PSS"`
	Psg []EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION_PSG `hl7:"GROUP;ATR=optional"`
}

// EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION_PSG - Group struct
type EHC_E02_INVOICE_INFORMATION_CANCEL_PRODUCT_SERVICE_SECTION_PSG struct {
	ProductServiceGroup PSG `hl7:"TAG=PSG"`
	ProductServiceLineItem []PSL `hl7:"TAG=PSL;ATR=optional"`
}

// EHC_E02 - Cancel HealthCare Services Invoice
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/EHC_E02
type EHC_E02 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	Invoice_information_cancel EHC_E02_INVOICE_INFORMATION_CANCEL `hl7:"GROUP"`
}

