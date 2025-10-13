package hl7v23

import "time"

//	PV1 - Patient visit
//
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   string      `hl7:"2" json:"setID,omitempty"`
	PatientClass            string      `hl7:"3" json:"patientClass,omitempty"`
	AssignedPatientLocation PL          `hl7:"4" json:"assignedPatientLocation,omitempty"`
	AdmissionType           string      `hl7:"5" json:"admissionType,omitempty"`
	PreadmitNumber          CX          `hl7:"6" json:"preadmitNumber,omitempty"`
	PriorPatientLocation    PL          `hl7:"7" json:"priorPatientLocation,omitempty"`
	AttendingDoctor         []XCN       `hl7:"8" json:"attendingDoctor,omitempty"`
	ReferringDoctor         []XCN       `hl7:"9" json:"referringDoctor,omitempty"`
	ConsultingDoctor        []XCN       `hl7:"10" json:"consultingDoctor,omitempty"`
	HospitalService         string      `hl7:"11" json:"hospitalService,omitempty"`
	TemporaryLocation       PL          `hl7:"12" json:"temporaryLocation,omitempty"`
	PreadmitTestIndicator   string      `hl7:"13" json:"preadmitTestIndicator,omitempty"`
	ReadmissionIndicator    string      `hl7:"14" json:"readmissionIndicator,omitempty"`
	AdmitSource             string      `hl7:"15" json:"admitSource,omitempty"`
	AmbulatoryStatus        []string    `hl7:"16" json:"ambulatoryStatus,omitempty"`
	VIPIndicator            string      `hl7:"17" json:"vipIndicator,omitempty"`
	AdmittingDoctor         []XCN       `hl7:"18" json:"admittingDoctor,omitempty"`
	PatientType             string      `hl7:"19" json:"patientType,omitempty"`
	VisitNumber             CX          `hl7:"20" json:"visitNumber,omitempty"`
	FinancialClass          FC          `hl7:"21" json:"financialClass,omitempty"`
	ChargePriceIndicator    string      `hl7:"22" json:"chargePriceIndicator,omitempty"`
	CourtesyCode            string      `hl7:"23" json:"courtesyCode,omitempty"`
	CreditRating            string      `hl7:"24" json:"creditRating,omitempty"`
	ContractCode            []string    `hl7:"25" json:"contractCode,omitempty"`
	ContractEffectiveDate   []time.Time `hl7:"26" json:"contractEffectiveDate,omitempty"`
	ContractAmount          []float32   `hl7:"27" json:"contractAmount,omitempty"`
	ContractPeriod          []float32   `hl7:"28" json:"contractPeriod,omitempty"`
	InterestCode            string      `hl7:"29" json:"interestCode,omitempty"`
	TransferToBadDebtCode   string      `hl7:"30" json:"transferToBadDebtCode,omitempty"`
	TransferToBadDebtDate   time.Time   `hl7:"31" json:"transferToBadDebtDate,omitempty"`
	BadDebtAgencyCode       string      `hl7:"32" json:"badDebtAgencyCode,omitempty"`
	BadDebtTransferAmount   float32     `hl7:"33" json:"badDebtTransferAmount,omitempty"`
	BadDebtRecoveryAmount   float32     `hl7:"34" json:"badDebtRecoveryAmount,omitempty"`
	DeleteAccountIndicator  string      `hl7:"35" json:"deleteAccountIndicator,omitempty"`
	DeleteAccountDate       time.Time   `hl7:"36" json:"deleteAccountDate,omitempty"`
	DischargeDisposition    string      `hl7:"37" json:"dischargeDisposition,omitempty"`
	DischargedToLocation    CM_DLD      `hl7:"38" json:"dischargedToLocation,omitempty"`
	DietType                string      `hl7:"39" json:"dietType,omitempty"`
	ServicingFacility       string      `hl7:"40" json:"servicingFacility,omitempty"`
	BedStatus               string      `hl7:"41" json:"bedStatus,omitempty"`
	AccountStatus           string      `hl7:"42" json:"accountStatus,omitempty"`
	PendingLocation         PL          `hl7:"43" json:"pendingLocation,omitempty"`
	PriorTemporaryLocation  PL          `hl7:"44" json:"priorTemporaryLocation,omitempty"`
	AdmitDateTime           time.Time   `hl7:"45" json:"admitDateTime,omitempty"`
	DischargeDateTime       time.Time   `hl7:"46" json:"dischargeDateTime,omitempty"`
	CurrentPatientBalance   float32     `hl7:"47" json:"currentPatientBalance,omitempty"`
	TotalCharges            float32     `hl7:"48" json:"totalCharges,omitempty"`
	TotalAdjustments        float32     `hl7:"49" json:"totalAdjustments,omitempty"`
	TotalPayments           float32     `hl7:"50" json:"totalPayments,omitempty"`
	AlternateVisitID        CX          `hl7:"51" json:"alternateVisitID,omitempty"`
	VisitIndicator1         string      `hl7:"52" json:"visitIndicator1,omitempty"`
	OtherHealthcareProvider []XCN       `hl7:"53" json:"otherHealthcareProvider,omitempty"`
}
