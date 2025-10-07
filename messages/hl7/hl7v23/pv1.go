package hl7v23

import "time"

//	PV1 - Patient visit
//
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   string      `hl7:"1" json:"setID,omitempty"`
	PatientClass            string      `hl7:"2" json:"patientClass,omitempty"`
	AssignedPatientLocation PL          `hl7:"3" json:"assignedPatientLocation,omitempty"`
	AdmissionType           string      `hl7:"4" json:"admissionType,omitempty"`
	PreadmitNumber          CX          `hl7:"5" json:"preadmitNumber,omitempty"`
	PriorPatientLocation    PL          `hl7:"6" json:"priorPatientLocation,omitempty"`
	AttendingDoctor         []XCN       `hl7:"7" json:"attendingDoctor,omitempty"`
	ReferringDoctor         []XCN       `hl7:"8" json:"referringDoctor,omitempty"`
	ConsultingDoctor        []XCN       `hl7:"9" json:"consultingDoctor,omitempty"`
	HospitalService         string      `hl7:"10" json:"hospitalService,omitempty"`
	TemporaryLocation       PL          `hl7:"11" json:"temporaryLocation,omitempty"`
	PreadmitTestIndicator   string      `hl7:"12" json:"preadmitTestIndicator,omitempty"`
	ReadmissionIndicator    string      `hl7:"13" json:"readmissionIndicator,omitempty"`
	AdmitSource             string      `hl7:"14" json:"admitSource,omitempty"`
	AmbulatoryStatus        []string    `hl7:"15" json:"ambulatoryStatus,omitempty"`
	VIPIndicator            string      `hl7:"16" json:"vipIndicator,omitempty"`
	AdmittingDoctor         []XCN       `hl7:"17" json:"admittingDoctor,omitempty"`
	PatientType             string      `hl7:"18" json:"patientType,omitempty"`
	VisitNumber             CX          `hl7:"19" json:"visitNumber,omitempty"`
	FinancialClass          FC          `hl7:"20" json:"financialClass,omitempty"`
	ChargePriceIndicator    string      `hl7:"21" json:"chargePriceIndicator,omitempty"`
	CourtesyCode            string      `hl7:"22" json:"courtesyCode,omitempty"`
	CreditRating            string      `hl7:"23" json:"creditRating,omitempty"`
	ContractCode            []string    `hl7:"24" json:"contractCode,omitempty"`
	ContractEffectiveDate   []time.Time `hl7:"25,shortdate" json:"contractEffectiveDate,omitempty"`
	ContractAmount          []float32   `hl7:"26" json:"contractAmount,omitempty"`
	ContractPeriod          []float32   `hl7:"27" json:"contractPeriod,omitempty"`
	InterestCode            string      `hl7:"28" json:"interestCode,omitempty"`
	TransferToBadDebtCode   string      `hl7:"29" json:"transferToBadDebtCode,omitempty"`
	TransferToBadDebtDate   time.Time   `hl7:"30,shortdate" json:"transferToBadDebtDate,omitempty"`
	BadDebtAgencyCode       string      `hl7:"31" json:"badDebtAgencyCode,omitempty"`
	BadDebtTransferAmount   float32     `hl7:"32" json:"badDebtTransferAmount,omitempty"`
	BadDebtRecoveryAmount   float32     `hl7:"33" json:"badDebtRecoveryAmount,omitempty"`
	DeleteAccountIndicator  string      `hl7:"34" json:"deleteAccountIndicator,omitempty"`
	DeleteAccountDate       time.Time   `hl7:"35" json:"deleteAccountDate,omitempty"`
	DischargeDisposition    string      `hl7:"36" json:"dischargeDisposition,omitempty"`
	DischargedToLocation    CM_DLD      `hl7:"37" json:"dischargedToLocation,omitempty"`
	DietType                string      `hl7:"38" json:"dietType,omitempty"`
	ServicingFacility       string      `hl7:"39" json:"servicingFacility,omitempty"`
	BedStatus               string      `hl7:"40" json:"bedStatus,omitempty"`
	AccountStatus           string      `hl7:"41" json:"accountStatus,omitempty"`
	PendingLocation         PL          `hl7:"42" json:"pendingLocation,omitempty"`
	PriorTemporaryLocation  PL          `hl7:"43" json:"priorTemporaryLocation,omitempty"`
	AdmitDateTime           time.Time   `hl7:"44" json:"admitDateTime,omitempty"`
	DischargeDateTime       time.Time   `hl7:"45" json:"dischargeDateTime,omitempty"`
	CurrentPatientBalance   float32     `hl7:"46" json:"currentPatientBalance,omitempty"`
	TotalCharges            float32     `hl7:"47" json:"totalCharges,omitempty"`
	TotalAdjustments        float32     `hl7:"48" json:"totalAdjustments,omitempty"`
	TotalPayments           float32     `hl7:"49" json:"totalPayments,omitempty"`
	AlternateVisitID        CX          `hl7:"50" json:"alternateVisitID,omitempty"`
	VisitIndicator1         string      `hl7:"51" json:"visitIndicator1,omitempty"`
	OtherHealthcareProvider []XCN       `hl7:"52" json:"otherHealthcareProvider,omitempty"`
}
