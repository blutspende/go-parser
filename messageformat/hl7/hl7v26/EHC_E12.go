package hl7v26

// EHC_E12_REQUEST - Group struct
type EHC_E12_REQUEST struct {
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments NTE `hl7:"TAG=NTE;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// EHC_E12 - Request Additional Information
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/EHC_E12
type EHC_E12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential []UAC `hl7:"TAG=UAC;ATR=optional"`
	RequestForInformation RFI `hl7:"TAG=RFI"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
	Invoice IVC `hl7:"TAG=IVC"`
	ProductServiceSection PSS `hl7:"TAG=PSS"`
	ProductServiceGroup PSG `hl7:"TAG=PSG"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	ProductServiceLineItem []PSL `hl7:"TAG=PSL;ATR=optional"`
	Request []EHC_E12_REQUEST `hl7:"GROUP"`
}

