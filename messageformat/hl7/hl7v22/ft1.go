package hl7v22

import "time"

// FT1 - Financial Transaction
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/FT1
type FT1 struct {
	SetID                           int                  `hl7:"POS=2;ATR=sequence"`
	TransactionID                   string               `hl7:"POS=3"`
	TransactionBatchID              string               `hl7:"POS=4"`
	TransactionDate                 time.Time            `hl7:"POS=5;ATR=date"`
	TransactionPostingDate          time.Time            `hl7:"POS=6;ATR=date"`
	TransactionType                 string               `hl7:"POS=7"`
	TransactionCode                 CE                   `hl7:"POS=8"`
	TransactionDescription          string               `hl7:"POS=9"`
	TransactionDescriptionAlternate string               `hl7:"POS=10"`
	TransactionQuantity             *int                 `hl7:"POS=11"`
	TransactionAmountExtended       *int                 `hl7:"POS=12"`
	TransactionAmountUnit           *int                 `hl7:"POS=13"`
	DepartmentCode                  CE                   `hl7:"POS=14"`
	InsurancePlanID                 string               `hl7:"POS=15"`
	InsuranceAmount                 *int                 `hl7:"POS=16"`
	AssignedPatientLocation         CM_INTERNAL_LOCATION `hl7:"POS=17"`
	FeeSchedule                     string               `hl7:"POS=18"`
	PatientType                     string               `hl7:"POS=19"`
	DiagnosisCode                   []CE                 `hl7:"POS=20"`
	PerformedByCode                 CN_PERSON            `hl7:"POS=21"`
	OrderedByCode                   CN_PERSON            `hl7:"POS=22"`
	UnitCost                        *int                 `hl7:"POS=23"`
	FillerOrderNumber               CM_FILLER            `hl7:"POS=24"`
}
