package hl7v24

import "time"

// PV1 - Patient visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   string      `hl7:"2" json:"SetID,omitempty"`
	PatientClass            string      `hl7:"3" json:"PatientClass,omitempty"`
	AssignedPatientLocation PL          `hl7:"4" json:"AssignedPatientLocation,omitempty"`
	AdmissionType           string      `hl7:"5" json:"AdmissionType,omitempty"`
	PreadmitNumber          CX          `hl7:"6" json:"PreadmitNumber,omitempty"`
	PriorPatientLocation    PL          `hl7:"7" json:"PriorPatientLocation,omitempty"`
	AttendingDoctor         []XCN       `hl7:"8" json:"AttendingDoctor,omitempty"`
	ReferringDoctor         []XCN       `hl7:"9" json:"ReferringDoctor,omitempty"`
	ConsultingDoctor        []XCN       `hl7:"10" json:"ConsultingDoctor,omitempty"`
	HospitalService         string      `hl7:"11" json:"HospitalService,omitempty"`
	TemporaryLocation       PL          `hl7:"12" json:"TemporaryLocation,omitempty"`
	PreadmitTestIndicator   string      `hl7:"13" json:"PreadmitTestIndicator,omitempty"`
	ReadmissionIndicator    string      `hl7:"14" json:"ReadmissionIndicator,omitempty"`
	AdmitSource             string      `hl7:"15" json:"AdmitSource,omitempty"`
	AmbulatoryStatus        []string    `hl7:"16" json:"AmbulatoryStatus,omitempty"`
	VIPIndicator            string      `hl7:"17" json:"VIPIndicator,omitempty"`
	AdmittingDoctor         []XCN       `hl7:"18" json:"AdmittingDoctor,omitempty"`
	PatientType             string      `hl7:"19" json:"PatientType,omitempty"`
	VisitNumber             CX          `hl7:"20" json:"VisitNumber,omitempty"`
	FinancialClass          []FC        `hl7:"21" json:"FinancialClass,omitempty"`
	ChargePriceIndicator    string      `hl7:"22" json:"ChargePriceIndicator,omitempty"`
	CourtesyCode            string      `hl7:"23" json:"CourtesyCode,omitempty"`
	CreditRating            string      `hl7:"24" json:"CreditRating,omitempty"`
	ContractCode            []string    `hl7:"25" json:"ContractCode,omitempty"`
	ContractEffectiveDate   []time.Time `hl7:"26" json:"ContractEffectiveDate,omitempty"`
	ContractAmount          []float32   `hl7:"27" json:"ContractAmount,omitempty"`
	ContractPeriod          []float32   `hl7:"28" json:"ContractPeriod,omitempty"`
	InterestCode            string      `hl7:"29" json:"InterestCode,omitempty"`
	TransferToBadDebtCode   string      `hl7:"30" json:"TransferToBadDebtCode,omitempty"`
	TransferToBadDebtDate   time.Time   `hl7:"31" json:"TransferToBadDebtDate,omitempty"`
	BadDebtAgencyCode       string      `hl7:"32" json:"BadDebtAgencyCode,omitempty"`
	BadDebtTransferAmount   float32     `hl7:"33" json:"BadDebtTransferAmount,omitempty"`
	BadDebtRecoveryAmount   float32     `hl7:"34" json:"BadDebtRecoveryAmount,omitempty"`
	DeleteAccountIndicator  string      `hl7:"35" json:"DeleteAccountIndicator,omitempty"`
	DeleteAccountDate       time.Time   `hl7:"36" json:"DeleteAccountDate,omitempty"`
	DischargeDisposition    string      `hl7:"37" json:"DischargeDisposition,omitempty"`
	DischargedToLocation    DLD         `hl7:"38" json:"DischargedToLocation,omitempty"`
	DietType                string      `hl7:"39" json:"DietType,omitempty"`
	ServicingFacility       string      `hl7:"40" json:"ServicingFacility,omitempty"`
	BedStatus               string      `hl7:"41" json:"BedStatus,omitempty"`
	AccountStatus           string      `hl7:"42" json:"AccountStatus,omitempty"`
	PendingLocation         PL          `hl7:"43" json:"PendingLocation,omitempty"`
	PriorTemporaryLocation  PL          `hl7:"44" json:"PriorTemporaryLocation,omitempty"`
	AdmitDateTime           time.Time   `hl7:"45,longdate" json:"AdmitDateTime,omitempty"`
	DischargeDateTime       []time.Time `hl7:"46,longdate" json:"DischargeDateTime,omitempty"`
	CurrentPatientBalance   float32     `hl7:"47" json:"CurrentPatientBalance,omitempty"`
	TotalCharges            float32     `hl7:"48" json:"TotalCharges,omitempty"`
	TotalAdjustments        float32     `hl7:"49" json:"TotalAdjustments,omitempty"`
	TotalPayments           float32     `hl7:"50" json:"TotalPayments,omitempty"`
	AlternateVisitID        CX          `hl7:"51" json:"AlternateVisitID,omitempty"`
	VisitIndicator1         string      `hl7:"52" json:"VisitIndicator1,omitempty"`
	OtherHealthcareProvider []XCN       `hl7:"53" json:"OtherHealthcareProvider,omitempty"`
}
