package hl7v231

// RQI_I03_PROVIDER - Group struct
type RQI_I03_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RQI_I03_GUARANTOR_INSURANCE - Group struct
type RQI_I03_GUARANTOR_INSURANCE struct {
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []RQI_I03_GUARANTOR_INSURANCE_INSURANCE `hl7:"GROUP"`
}

// RQI_I03_GUARANTOR_INSURANCE_INSURANCE - Group struct
type RQI_I03_GUARANTOR_INSURANCE_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// RQI_I03 - Request/receipt of patient selection list
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/RQI_I03
type RQI_I03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	Provider []RQI_I03_PROVIDER `hl7:"GROUP"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	NextOfKinAssociatedPartiesSegment []NK1 `hl7:"TAG=NK1;ATR=optional"`
	GuarantorInsurance RQI_I03_GUARANTOR_INSURANCE `hl7:"GROUP;ATR=optional"`
	NotesAndCommentsSegment []NTE `hl7:"TAG=NTE;ATR=optional"`
}

