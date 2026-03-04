package hl7v28

// EHC_E20_AUTHORIZATION_REQUEST - Group struct
type EHC_E20_AUTHORIZATION_REQUEST struct {
	InvoiceSegment []IVC `hl7:"TAG=IVC;ATR=optional"`
	ContactData []CTD `hl7:"TAG=CTD"`
	LocationIdentification []LOC `hl7:"TAG=LOC;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	Pat_info []EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO `hl7:"GROUP"`
	Psl_item_info []EHC_E20_AUTHORIZATION_REQUEST_PSL_ITEM_INFO `hl7:"GROUP"`
}

// EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO - Group struct
type EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO struct {
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	Accident []ACC `hl7:"TAG=ACC;ATR=optional"`
	Insurance []EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_INSURANCE `hl7:"GROUP"`
	Diagnosis EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_DIAGNOSIS `hl7:"GROUP;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
}

// EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_INSURANCE - Group struct
type EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1;ATR=optional"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
}

// EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_DIAGNOSIS - Group struct
type EHC_E20_AUTHORIZATION_REQUEST_PAT_INFO_DIAGNOSIS struct {
	Diagnosis DG1 `hl7:"TAG=DG1"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// EHC_E20_AUTHORIZATION_REQUEST_PSL_ITEM_INFO - Group struct
type EHC_E20_AUTHORIZATION_REQUEST_PSL_ITEM_INFO struct {
	ProductServiceLineItem PSL `hl7:"TAG=PSL"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Adjustment []ADJ `hl7:"TAG=ADJ;ATR=optional"`
	LocationIdentification []LOC `hl7:"TAG=LOC;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// EHC_E20 - Submit Authorization Request
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/EHC_E20
type EHC_E20 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment []UAC `hl7:"TAG=UAC;ATR=optional"`
	Authorization_request EHC_E20_AUTHORIZATION_REQUEST `hl7:"GROUP"`
}

