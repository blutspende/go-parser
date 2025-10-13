package hl7v24

import "time"

type FT1 struct {
	SetID                     int       `hl7:"2,sequence" json:"SetID,omitempty"`
	TransactionID             string    `hl7:"3" json:"TransactionID,omitempty"`
	TransactionBatchID        string    `hl7:"4" json:"TransactionBatchID,omitempty"`
	TransactionDate           time.Time `hl7:"5,longdate" json:"TransactionDate,omitempty"`
	TransactionPostingDate    time.Time `hl7:"6,longdate" json:"TransactionPostingDate,omitempty"`
	TransactionType           string    `hl7:"7" json:"TransactionType,omitempty"`
	TransactionCode           CE        `hl7:"8" json:"TransactionCode,omitempty"`
	TransactionDescription    string    `hl7:"9" json:"TransactionDescription,omitempty"`
	TransactionDescriptionAlt string    `hl7:"10" json:"TransactionDescriptionAlt,omitempty"`
	TransactionQuantity       float32   `hl7:"11" json:"TransactionQuantity,omitempty"`
	TransactionAmountExtended CP        `hl7:"12" json:"TransactionAmountExtended,omitempty"`
	TransactionAmountUnit     CP        `hl7:"13" json:"TransactionAmountUnit,omitempty"`
	DepartmentCode            CE        `hl7:"14" json:"DepartmentCode,omitempty"`
	InsurancePlanID           CE        `hl7:"15" json:"InsurancePlanID,omitempty"`
	InsuranceAmount           CP        `hl7:"16" json:"InsuranceAmount,omitempty"`
	AssignedPatientLocaton    PL        `hl7:"17" json:"AssignedPatientLocaton,omitempty"`
	FeeSchedule               string    `hl7:"18" json:"FeeSchedule,omitempty"`
	PatientType               string    `hl7:"19" json:"PatientType,omitempty"`
	DiagnosisCodeFT1          []CE      `hl7:"20" json:"DiagnosisCodeFT1,omitempty"`
	PerformedByCode           []XCN     `hl7:"21" json:"PerformedByCode,omitempty"`
	OrderedByCode             []XCN     `hl7:"22" json:"OrderedByCode,omitempty"`
	UnitCost                  CP        `hl7:"23" json:"UnitCost,omitempty"`
	FillerOrderNUmber         EI        `hl7:"24" json:"FillerOrderNUmber,omitempty"`
	EnteredByCode             []XCN     `hl7:"25" json:"EnteredByCode,omitempty"`
	ProcedureCode             CE        `hl7:"26" json:"ProcedureCode,omitempty"`
	ProcedureCodeModifier     []CE      `hl7:"27" json:"ProcedureCodeModifier,omitempty"`
}
