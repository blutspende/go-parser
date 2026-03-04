package hl7v27

// RQI_I03_PROVIDER - Group struct
type RQI_I03_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQI_I03_GUARANTOR_INSURANCE - Group struct
type RQI_I03_GUARANTOR_INSURANCE struct {
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RQI_I03_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RQI_I03_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RQI_I03_GUARANTOR_INSURANCE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQI_I03 - Request/receipt of patient selection list
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/RQI_I03
type RQI_I03 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Provider []RQI_I03_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor_insurance RQI_I03_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

