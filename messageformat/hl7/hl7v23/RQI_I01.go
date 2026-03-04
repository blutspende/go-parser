package hl7v23

// RQI_I01_PROVIDER - Group struct
type RQI_I01_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQI_I01_GUARANTOR_INSURANCE - Group struct
type RQI_I01_GUARANTOR_INSURANCE struct {
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RQI_I01_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RQI_I01_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RQI_I01_GUARANTOR_INSURANCE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQI_I01 - Request for insurance information
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/RQI_I01
type RQI_I01 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Provider []RQI_I01_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RQI_I01_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

