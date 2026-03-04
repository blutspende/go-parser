package hl7v28

// EHC_E13_REQUEST - Group struct
type EHC_E13_REQUEST struct {
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
	Response []EHC_E13_REQUEST_RESPONSE `hl7:"GROUP"`
}

// EHC_E13_REQUEST_RESPONSE - Group struct
type EHC_E13_REQUEST_RESPONSE struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
	TranscriptionDocumentHeader TXA `hl7:"TAG=TXA;ATR=optional"`
}

// EHC_E13 - Additional Information Response
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/EHC_E13
type EHC_E13 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error []ERR `hl7:"TAG=ERR;ATR=optional"`
	RequestForInformation RFI `hl7:"TAG=RFI"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
	InvoiceSegment IVC `hl7:"TAG=IVC"`
	ProductServiceSection PSS `hl7:"TAG=PSS"`
	ProductServiceGroup PSG `hl7:"TAG=PSG"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	ProductServiceLineItem PSL `hl7:"TAG=PSL;ATR=optional"`
	Request []EHC_E13_REQUEST `hl7:"GROUP"`
}

