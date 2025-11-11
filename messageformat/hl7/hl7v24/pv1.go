package hl7v24

import "time"

// PV1 - Patient visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   int         `hl7:"POS=2;ATR=sequence" json:"SetID,omitempty"`
	PatientClass            string      `hl7:"POS=3" json:"PatientClass,omitempty"`
	AssignedPatientLocation PL          `hl7:"POS=4" json:"AssignedPatientLocation,omitempty"`
	AdmissionType           string      `hl7:"POS=5" json:"AdmissionType,omitempty"`
	PreadmitNumber          CX          `hl7:"POS=6" json:"PreadmitNumber,omitempty"`
	PriorPatientLocation    PL          `hl7:"POS=7" json:"PriorPatientLocation,omitempty"`
	AttendingDoctor         []XCN       `hl7:"POS=8" json:"AttendingDoctor,omitempty"`
	ReferringDoctor         []XCN       `hl7:"POS=9" json:"ReferringDoctor,omitempty"`
	ConsultingDoctor        []XCN       `hl7:"POS=10" json:"ConsultingDoctor,omitempty"`
	HospitalService         string      `hl7:"POS=11" json:"HospitalService,omitempty"`
	TemporaryLocation       PL          `hl7:"POS=12" json:"TemporaryLocation,omitempty"`
	PreadmitTestIndicator   string      `hl7:"POS=13" json:"PreadmitTestIndicator,omitempty"`
	ReadmissionIndicator    string      `hl7:"POS=14" json:"ReadmissionIndicator,omitempty"`
	AdmitSource             string      `hl7:"POS=15" json:"AdmitSource,omitempty"`
	AmbulatoryStatus        []string    `hl7:"POS=16" json:"AmbulatoryStatus,omitempty"`
	VIPIndicator            string      `hl7:"POS=17" json:"VIPIndicator,omitempty"`
	AdmittingDoctor         []XCN       `hl7:"POS=18" json:"AdmittingDoctor,omitempty"`
	PatientType             string      `hl7:"POS=19" json:"PatientType,omitempty"`
	VisitNumber             CX          `hl7:"POS=20" json:"VisitNumber,omitempty"`
	FinancialClass          []FC        `hl7:"POS=21" json:"FinancialClass,omitempty"`
	ChargePriceIndicator    string      `hl7:"POS=22" json:"ChargePriceIndicator,omitempty"`
	CourtesyCode            string      `hl7:"POS=23" json:"CourtesyCode,omitempty"`
	CreditRating            string      `hl7:"POS=24" json:"CreditRating,omitempty"`
	ContractCode            []string    `hl7:"POS=25" json:"ContractCode,omitempty"`
	ContractEffectiveDate   []time.Time `hl7:"POS=26;ATR=date" json:"ContractEffectiveDate,omitempty"`
	ContractAmount          []float32   `hl7:"POS=27" json:"ContractAmount,omitempty"`
	ContractPeriod          []float32   `hl7:"POS=28" json:"ContractPeriod,omitempty"`
	InterestCode            string      `hl7:"POS=29" json:"InterestCode,omitempty"`
	TransferToBadDebtCode   string      `hl7:"POS=30" json:"TransferToBadDebtCode,omitempty"`
	TransferToBadDebtDate   time.Time   `hl7:"POS=31;ATR=date" json:"TransferToBadDebtDate,omitempty"`
	BadDebtAgencyCode       string      `hl7:"POS=32" json:"BadDebtAgencyCode,omitempty"`
	BadDebtTransferAmount   *float32    `hl7:"POS=33" json:"BadDebtTransferAmount,omitempty"`
	BadDebtRecoveryAmount   *float32    `hl7:"POS=34" json:"BadDebtRecoveryAmount,omitempty"`
	DeleteAccountIndicator  string      `hl7:"POS=35" json:"DeleteAccountIndicator,omitempty"`
	DeleteAccountDate       time.Time   `hl7:"POS=36;ATR=date" json:"DeleteAccountDate,omitempty"`
	DischargeDisposition    string      `hl7:"POS=37" json:"DischargeDisposition,omitempty"`
	DischargedToLocation    DLD         `hl7:"POS=38" json:"DischargedToLocation,omitempty"`
	DietType                string      `hl7:"POS=39" json:"DietType,omitempty"`
	ServicingFacility       string      `hl7:"POS=40" json:"ServicingFacility,omitempty"`
	BedStatus               string      `hl7:"POS=41" json:"BedStatus,omitempty"`
	AccountStatus           string      `hl7:"POS=42" json:"AccountStatus,omitempty"`
	PendingLocation         PL          `hl7:"POS=43" json:"PendingLocation,omitempty"`
	PriorTemporaryLocation  PL          `hl7:"POS=44" json:"PriorTemporaryLocation,omitempty"`
	AdmitDateTime           time.Time   `hl7:"POS=45" json:"AdmitDateTime,omitempty"`
	DischargeDateTime       []time.Time `hl7:"POS=46" json:"DischargeDateTime,omitempty"`
	CurrentPatientBalance   *float32    `hl7:"POS=47" json:"CurrentPatientBalance,omitempty"`
	TotalCharges            *float32    `hl7:"POS=48" json:"TotalCharges,omitempty"`
	TotalAdjustments        *float32    `hl7:"POS=49" json:"TotalAdjustments,omitempty"`
	TotalPayments           *float32    `hl7:"POS=50" json:"TotalPayments,omitempty"`
	AlternateVisitID        CX          `hl7:"POS=51" json:"AlternateVisitID,omitempty"`
	VisitIndicator1         string      `hl7:"POS=52" json:"VisitIndicator1,omitempty"`
	OtherHealthcareProvider []XCN       `hl7:"POS=53" json:"OtherHealthcareProvider,omitempty"`
}
