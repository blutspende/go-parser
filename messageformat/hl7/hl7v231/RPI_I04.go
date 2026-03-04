package hl7v231

// RPI_I04_PROVIDER - Group struct
type RPI_I04_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPI_I04_GUARANTOR_INSURANCE - Group struct
type RPI_I04_GUARANTOR_INSURANCE struct {
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RPI_I04_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RPI_I04_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RPI_I04_GUARANTOR_INSURANCE_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RPI_I04 - Return patient information
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RPI_I04
type RPI_I04 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgmentSegment MSA `hl7:"TAG=MSA"`
	Provider []RPI_I04_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RPI_I04_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

