package hl7v271

// PIN_I07_PROVIDER - Group struct
type PIN_I07_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// PIN_I07_GUARANTOR_INSURANCE - Group struct
type PIN_I07_GUARANTOR_INSURANCE struct {
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []PIN_I07_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// PIN_I07_GUARANTOR_INSURANCE_INSURANCE - Group struct
type PIN_I07_GUARANTOR_INSURANCE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// PIN_I07 - Unsolicited insurance information
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/PIN_I07
type PIN_I07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Provider []PIN_I07_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Guarantor_insurance PIN_I07_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

