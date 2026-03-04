package hl7v23

// RPI_I01_PROVIDER - Group struct
type RPI_I01_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RPI_I01_GUARANTOR_INSURANCE - Group struct
type RPI_I01_GUARANTOR_INSURANCE struct {
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RPI_I01_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RPI_I01_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RPI_I01_GUARANTOR_INSURANCE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RPI_I01 - Response for insurance information
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RPI_I01
type RPI_I01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	Provider []RPI_I01_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RPI_I01_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

