package hl7v23

import "time"

// RQD - Requisition detail
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RQD
type RQD struct {
	RequisitionLineNumber    string    `hl7:"POS=2" json:"requisitionLineNumber,omitempty"`
	ItemCodeInternal         CE        `hl7:"POS=3" json:"itemCodeInternal,omitempty"`
	ItemCodeExternal         CE        `hl7:"POS=4" json:"itemCodeExternal,omitempty"`
	HospitalItemCode         CE        `hl7:"POS=5" json:"hospitalItemCode,omitempty"`
	RequisitionQuantity      *float32  `hl7:"POS=6" json:"requisitionQuantity,omitempty"`
	RequisitionUnitOfMeasure CE        `hl7:"POS=7" json:"requisitionUnitOfMeasure,omitempty"`
	DepartmentCostCenter     string    `hl7:"POS=8" json:"departmentCostCenter,omitempty"`
	ItemNaturalAccountCode   string    `hl7:"POS=9" json:"itemNaturalAccountCode,omitempty"`
	DeliverToID              CE        `hl7:"POS=10" json:"deliverToID,omitempty"`
	DateNeeded               time.Time `hl7:"POS=11;ATR=date" json:"dateNeeded,omitempty"`
}
