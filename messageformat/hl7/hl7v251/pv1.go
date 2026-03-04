package hl7v251

import "time"

// PV1 - Patient Visit
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/PV1
type PV1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PatientClass string `hl7:"POS=3"`
	AssignedPatientLocation PL `hl7:"POS=4"`
	AdmissionType string `hl7:"POS=5"`
	PreAdmitNumber CX `hl7:"POS=6"`
	PriorPatientLocation PL `hl7:"POS=7"`
	AttendingDoctor []XCN `hl7:"POS=8"`
	ReferringDoctor []XCN `hl7:"POS=9"`
	ConsultingDoctor []XCN `hl7:"POS=10"`
	HospitalService string `hl7:"POS=11"`
	TemporaryLocation PL `hl7:"POS=12"`
	PreAdmitTestIndicator string `hl7:"POS=13"`
	ReAdmissionIndicator string `hl7:"POS=14"`
	AdmitSource string `hl7:"POS=15"`
	AmbulatoryStatus []string `hl7:"POS=16"`
	VipIndicator string `hl7:"POS=17"`
	AdmittingDoctor []XCN `hl7:"POS=18"`
	PatientType string `hl7:"POS=19"`
	VisitNumber CX `hl7:"POS=20"`
	FinancialClass []FC `hl7:"POS=21"`
	ChargePriceIndicator string `hl7:"POS=22"`
	CourtesyCode string `hl7:"POS=23"`
	CreditRating string `hl7:"POS=24"`
	ContractCode []string `hl7:"POS=25"`
	ContractEffectiveDate []time.Time `hl7:"POS=26;ATR=date"`
	ContractAmount []*int `hl7:"POS=27"`
	ContractPeriod []*int `hl7:"POS=28"`
	InterestCode string `hl7:"POS=29"`
	TransferToBadDebtCode string `hl7:"POS=30"`
	TransferToBadDebtDate time.Time `hl7:"POS=31;ATR=date"`
	BadDebtAgencyCode string `hl7:"POS=32"`
	BadDebtTransferAmount *int `hl7:"POS=33"`
	BadDebtRecoveryAmount *int `hl7:"POS=34"`
	DeleteAccountIndicator string `hl7:"POS=35"`
	DeleteAccountDate time.Time `hl7:"POS=36;ATR=date"`
	DischargeDisposition string `hl7:"POS=37"`
	DischargedToLocation DLD `hl7:"POS=38"`
	DietType CE `hl7:"POS=39"`
	ServicingFacility string `hl7:"POS=40"`
	BedStatus string `hl7:"POS=41"`
	AccountStatus string `hl7:"POS=42"`
	PendingLocation PL `hl7:"POS=43"`
	PriorTemporaryLocation PL `hl7:"POS=44"`
	AdmitDateTime time.Time `hl7:"POS=45"`
	DischargeDateTime []time.Time `hl7:"POS=46"`
	CurrentPatientBalance *int `hl7:"POS=47"`
	TotalCharges *int `hl7:"POS=48"`
	TotalAdjustments *int `hl7:"POS=49"`
	TotalPayments *int `hl7:"POS=50"`
	AlternateVisitID CX `hl7:"POS=51"`
	VisitIndicator string `hl7:"POS=52"`
	OtherHealthcareProvider []XCN `hl7:"POS=53"`
}

