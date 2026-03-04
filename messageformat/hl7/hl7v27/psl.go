package hl7v27

import "time"

// PSL - Product/service Line Item
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/PSL
type PSL struct {
	ProviderProductServiceLineItemNumber EI `hl7:"POS=2"`
	PayerProductServiceLineItemNumber EI `hl7:"POS=3"`
	ProductServiceLineItemSequenceNumber int `hl7:"POS=4;ATR=sequence"`
	ProviderTrackingID EI `hl7:"POS=5"`
	PayerTrackingID EI `hl7:"POS=6"`
	ProductServiceLineItemStatus CWE `hl7:"POS=7"`
	ProductServiceCode CWE `hl7:"POS=8"`
	ProductServiceCodeModifier CWE `hl7:"POS=9"`
	ProductServiceCodeDescription string `hl7:"POS=10"`
	ProductServiceEffectiveDate time.Time `hl7:"POS=11"`
	ProductServiceExpirationDate time.Time `hl7:"POS=12"`
	ProductServiceQuantity CQ `hl7:"POS=13"`
	ProductServiceUnitCost CP `hl7:"POS=14"`
	NumberOfItemsPerUnit *int `hl7:"POS=15"`
	ProductServiceGrossAmount CP `hl7:"POS=16"`
	ProductServiceBilledAmount CP `hl7:"POS=17"`
	ProductServiceClarificationCodeType CWE `hl7:"POS=18"`
	ProductServiceClarificationCodeValue string `hl7:"POS=19"`
	HealthDocumentReferenceIdentifier EI `hl7:"POS=20"`
	ProcessingConsiderationCode CWE `hl7:"POS=21"`
	RestrictedDisclosureIndicator string `hl7:"POS=22"`
	RelatedProductServiceCodeIndicator CWE `hl7:"POS=23"`
	ProductServiceAmountForPhysician CP `hl7:"POS=24"`
	ProductServiceCostFactor *int `hl7:"POS=25"`
	CostCenter CX `hl7:"POS=26"`
	BillingPeriod DR `hl7:"POS=27"`
	DaysWithoutBilling *int `hl7:"POS=28"`
	SessionNo *int `hl7:"POS=29"`
	ExecutingPhysicianID XCN `hl7:"POS=30"`
	ResponsiblePhysicianID XCN `hl7:"POS=31"`
	RoleExecutingPhysician CWE `hl7:"POS=32"`
	MedicalRoleExecutingPhysician CWE `hl7:"POS=33"`
	SideOfBody CWE `hl7:"POS=34"`
	NumberOfTpsPp *int `hl7:"POS=35"`
	TpValuePp CP `hl7:"POS=36"`
	InternalScalingFactorPp *int `hl7:"POS=37"`
	ExternalScalingFactorPp *int `hl7:"POS=38"`
	AmountPp CP `hl7:"POS=39"`
	NumberOfTpsTechnicalPart *int `hl7:"POS=40"`
	TpValueTechnicalPart CP `hl7:"POS=41"`
	InternalScalingFactorTechnicalPart *int `hl7:"POS=42"`
	ExternalScalingFactorTechnicalPart *int `hl7:"POS=43"`
	AmountTechnicalPart CP `hl7:"POS=44"`
	TotalAmountProfessionalPartTechnicalPart CP `hl7:"POS=45"`
	VatRate *int `hl7:"POS=46"`
	MainService string `hl7:"POS=47"`
	Validation string `hl7:"POS=48"`
	Comment string `hl7:"POS=49"`
}

