package hl7v271

// ORG - Practitioner Organization Unit Segment
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/ORG
type ORG struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	OrganizationUnitCode CWE `hl7:"POS=3"`
	OrganizationUnitTypeCode CWE `hl7:"POS=4"`
	PrimaryOrgUnitIndicator string `hl7:"POS=5"`
	PractitionerOrgUnitIdentifier CX `hl7:"POS=6"`
	HealthCareProviderTypeCode CWE `hl7:"POS=7"`
	HealthCareProviderClassificationCode CWE `hl7:"POS=8"`
	HealthCareProviderAreaOfSpecializationCode CWE `hl7:"POS=9"`
	EffectiveDateRange DR `hl7:"POS=10"`
	EmploymentStatusCode CWE `hl7:"POS=11"`
	BoardApprovalIndicator string `hl7:"POS=12"`
	PrimaryCarePhysicianIndicator string `hl7:"POS=13"`
	CostCenterCode []CWE `hl7:"POS=14"`
}

