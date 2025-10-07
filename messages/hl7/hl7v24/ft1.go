package hl7v24

import "time"

type FT1 struct {
	SetID                     int       `hl7:"1,sequence" json:"SetID,omitempty"`
	TransactionID             string    `hl7:"2" json:"TransactionID,omitempty"`
	TransactionBatchID        string    `hl7:"3" json:"TransactionBatchID,omitempty"`
	TransactionDate           time.Time `hl7:"4,longdate" json:"TransactionDate,omitempty"`
	TransactionPostingDate    time.Time `hl7:"5,longdate" json:"TransactionPostingDate,omitempty"`
	TransactionType           string    `hl7:"6" json:"TransactionType,omitempty"`
	TransactionCode           CE        `hl7:"7" json:"TransactionCode,omitempty"`
	TransactionDescription    string    `hl7:"8" json:"TransactionDescription,omitempty"`
	TransactionDescriptionAlt string    `hl7:"9" json:"TransactionDescriptionAlt,omitempty"`
	TransactionQuantity       float32   `hl7:"10" json:"TransactionQuantity,omitempty"`
	TransactionAmountExtended CP        `hl7:"11" json:"TransactionAmountExtended,omitempty"`
	TransactionAmountUnit     CP        `hl7:"12" json:"TransactionAmountUnit,omitempty"`
	DepartmentCode            CE        `hl7:"13" json:"DepartmentCode,omitempty"`
	InsurancePlanID           CE        `hl7:"14" json:"InsurancePlanID,omitempty"`
	InsuranceAmount           CP        `hl7:"15" json:"InsuranceAmount,omitempty"`
	AssignedPatientLocaton    PL        `hl7:"16" json:"AssignedPatientLocaton,omitempty"`
	FeeSchedule               string    `hl7:"17" json:"FeeSchedule,omitempty"`
	PatientType               string    `hl7:"18" json:"PatientType,omitempty"`
	DiagnosisCodeFT1          []CE      `hl7:"19" json:"DiagnosisCodeFT1,omitempty"`
	PerformedByCode           []XCN     `hl7:"20" json:"PerformedByCode,omitempty"`
	OrderedByCode             []XCN     `hl7:"21" json:"OrderedByCode,omitempty"`
	UnitCost                  CP        `hl7:"22" json:"UnitCost,omitempty"`
	FillerOrderNUmber         EI        `hl7:"23" json:"FillerOrderNUmber,omitempty"`
	EnteredByCode             []XCN     `hl7:"24" json:"EnteredByCode,omitempty"`
	ProcedureCode             CE        `hl7:"25" json:"ProcedureCode,omitempty"`
	ProcedureCodeModifier     []CE      `hl7:"26" json:"ProcedureCodeModifier,omitempty"`
}
