package hl7v24

import "time"

type FT1 struct {
	SetID                     *int      `hl7:"POS=2;ATR=sequence" json:"SetID,omitempty"`
	TransactionID             string    `hl7:"POS=3" json:"TransactionID,omitempty"`
	TransactionBatchID        string    `hl7:"POS=4" json:"TransactionBatchID,omitempty"`
	TransactionDate           time.Time `hl7:"POS=5" json:"TransactionDate,omitempty"`
	TransactionPostingDate    time.Time `hl7:"POS=6" json:"TransactionPostingDate,omitempty"`
	TransactionType           string    `hl7:"POS=7" json:"TransactionType,omitempty"`
	TransactionCode           CE        `hl7:"POS=8" json:"TransactionCode,omitempty"`
	TransactionDescription    string    `hl7:"POS=9" json:"TransactionDescription,omitempty"`
	TransactionDescriptionAlt string    `hl7:"POS=10" json:"TransactionDescriptionAlt,omitempty"`
	TransactionQuantity       *float32  `hl7:"POS=11" json:"TransactionQuantity,omitempty"`
	TransactionAmountExtended CP        `hl7:"POS=12" json:"TransactionAmountExtended,omitempty"`
	TransactionAmountUnit     CP        `hl7:"POS=13" json:"TransactionAmountUnit,omitempty"`
	DepartmentCode            CE        `hl7:"POS=14" json:"DepartmentCode,omitempty"`
	InsurancePlanID           CE        `hl7:"POS=15" json:"InsurancePlanID,omitempty"`
	InsuranceAmount           CP        `hl7:"POS=16" json:"InsuranceAmount,omitempty"`
	AssignedPatientLocaton    PL        `hl7:"POS=17" json:"AssignedPatientLocaton,omitempty"`
	FeeSchedule               string    `hl7:"POS=18" json:"FeeSchedule,omitempty"`
	PatientType               string    `hl7:"POS=19" json:"PatientType,omitempty"`
	DiagnosisCodeFT1          []CE      `hl7:"POS=20" json:"DiagnosisCodeFT1,omitempty"`
	PerformedByCode           []XCN     `hl7:"POS=21" json:"PerformedByCode,omitempty"`
	OrderedByCode             []XCN     `hl7:"POS=22" json:"OrderedByCode,omitempty"`
	UnitCost                  CP        `hl7:"POS=23" json:"UnitCost,omitempty"`
	FillerOrderNUmber         EI        `hl7:"POS=24" json:"FillerOrderNUmber,omitempty"`
	EnteredByCode             []XCN     `hl7:"POS=25" json:"EnteredByCode,omitempty"`
	ProcedureCode             CE        `hl7:"POS=26" json:"ProcedureCode,omitempty"`
	ProcedureCodeModifier     []CE      `hl7:"POS=27" json:"ProcedureCodeModifier,omitempty"`
}
