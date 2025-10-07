package hl7v23

import "time"

// RQD - Requisition detail
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RQD
type RQD struct {
	RequisitionLineNumber    string    `hl7:"1" json:"requisitionLineNumber,omitempty"`
	ItemCodeInternal         CE        `hl7:"2" json:"itemCodeInternal,omitempty"`
	ItemCodeExternal         CE        `hl7:"3" json:"itemCodeExternal,omitempty"`
	HospitalItemCode         CE        `hl7:"4" json:"hospitalItemCode,omitempty"`
	RequisitionQuantity      float32   `hl7:"5" json:"requisitionQuantity,omitempty"`
	RequisitionUnitOfMeasure CE        `hl7:"6" json:"requisitionUnitOfMeasure,omitempty"`
	DepartmentCostCenter     string    `hl7:"7" json:"departmentCostCenter,omitempty"`
	ItemNaturalAccountCode   string    `hl7:"8" json:"itemNaturalAccountCode,omitempty"`
	DeliverToID              CE        `hl7:"9" json:"deliverToID,omitempty"`
	DateNeeded               time.Time `hl7:"10,shortdate" json:"dateNeeded,omitempty"`
}
