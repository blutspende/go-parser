package hl7v27

import "time"

// FT1 - Financial Transaction
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/FT1
type FT1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	TransactionID string `hl7:"POS=3"`
	TransactionBatchID string `hl7:"POS=4"`
	TransactionDate DR `hl7:"POS=5"`
	TransactionPostingDate time.Time `hl7:"POS=6"`
	TransactionType CWE `hl7:"POS=7"`
	TransactionCode CWE `hl7:"POS=8"`
	TransactionQuantity *int `hl7:"POS=11"`
	TransactionAmountExtended CP `hl7:"POS=12"`
	TransactionAmountUnit CP `hl7:"POS=13"`
	DepartmentCode CWE `hl7:"POS=14"`
	HealthPlanID CWE `hl7:"POS=15"`
	InsuranceAmount CP `hl7:"POS=16"`
	AssignedPatientLocation PL `hl7:"POS=17"`
	FeeSchedule CWE `hl7:"POS=18"`
	PatientType CWE `hl7:"POS=19"`
	DiagnosisCodeFt1 []CWE `hl7:"POS=20"`
	PerformedByCode []XCN `hl7:"POS=21"`
	OrderedByCode []XCN `hl7:"POS=22"`
	UnitCost CP `hl7:"POS=23"`
	FillerOrderNumber EI `hl7:"POS=24"`
	EnteredByCode []XCN `hl7:"POS=25"`
	ProcedureCode CNE `hl7:"POS=26"`
	ProcedureCodeModifier []CNE `hl7:"POS=27"`
	AdvancedBeneficiaryNoticeCode CWE `hl7:"POS=28"`
	MedicallyNecessaryDuplicateProcedureReason CWE `hl7:"POS=29"`
	NdcCode CWE `hl7:"POS=30"`
	PaymentReferenceID CX `hl7:"POS=31"`
	TransactionReferenceKey []int `hl7:"POS=32;ATR=sequence"`
	PerformingFacility []XON `hl7:"POS=33"`
	OrderingFacility XON `hl7:"POS=34"`
	ItemNumber CWE `hl7:"POS=35"`
	ModelNumber string `hl7:"POS=36"`
	SpecialProcessingCode []CWE `hl7:"POS=37"`
	ClinicCode CWE `hl7:"POS=38"`
	ReferralNumber CX `hl7:"POS=39"`
	AuthorizationNumber CX `hl7:"POS=40"`
	ServiceProviderTaxonomyCode CWE `hl7:"POS=41"`
	RevenueCode CWE `hl7:"POS=42"`
	PrescriptionNumber string `hl7:"POS=43"`
	NdcQtyAndUom CQ `hl7:"POS=44"`
}

