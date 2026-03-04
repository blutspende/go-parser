package hl7v22

import "time"

// PV2 - Patient Visit - Additional Information
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/PV2
type PV2 struct {
	PriorPendingLocation CM_INTERNAL_LOCATION `hl7:"POS=2"`
	AccommodationCode CE `hl7:"POS=3"`
	AdmitReason CE `hl7:"POS=4"`
	TransferReason CE `hl7:"POS=5"`
	PatientValuables []string `hl7:"POS=6"`
	PatientValuablesLocation string `hl7:"POS=7"`
	VisitUserCode string `hl7:"POS=8"`
	ExpectedAdmitDate time.Time `hl7:"POS=9;ATR=date"`
	ExpectedDischargeDate time.Time `hl7:"POS=10;ATR=date"`
}

