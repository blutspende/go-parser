package hl7v24

// ORG - Practitioner Organization Unit
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ORG
type ORG struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	OrganizationUnitCode CE `hl7:"POS=3"`
	OrganizationUnitTypeCodeOrg CE `hl7:"POS=4"`
	PrimaryOrgUnitIndicator string `hl7:"POS=5"`
	PractitionerOrgUnitIdentifier CX `hl7:"POS=6"`
	HealthCareProviderTypeCode CE `hl7:"POS=7"`
	HealthCareProviderClassificationCode CE `hl7:"POS=8"`
	HealthCareProviderAreaOfSpecializationCode CE `hl7:"POS=9"`
	EffectiveDateRange DR `hl7:"POS=10"`
	EmploymentStatusCode CE `hl7:"POS=11"`
	BoardApprovalIndicator string `hl7:"POS=12"`
	PrimaryCarePhysicianIndicator string `hl7:"POS=13"`
}

