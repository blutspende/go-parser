package hl7v22

import "time"

// IN2 - Insurance Additional Info
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/IN2
type IN2 struct {
	InsuredsEmployeeID string `hl7:"POS=2"`
	InsuredsSocialSecurityNumber *int `hl7:"POS=3"`
	InsuredsEmployerName CN_PERSON `hl7:"POS=4"`
	EmployerInformationData string `hl7:"POS=5"`
	MailClaimParty string `hl7:"POS=6"`
	MedicareHealthInsuranceCardNumber *int `hl7:"POS=7"`
	MedicaidCaseName PN `hl7:"POS=8"`
	MedicaidCaseNumber *int `hl7:"POS=9"`
	ChampusSponsorName PN `hl7:"POS=10"`
	ChampusIDNumber *int `hl7:"POS=11"`
	DependentOfChampusRecipient string `hl7:"POS=12"`
	ChampusOrganization string `hl7:"POS=13"`
	ChampusStation string `hl7:"POS=14"`
	ChampusService string `hl7:"POS=15"`
	ChampusRankGrade string `hl7:"POS=16"`
	ChampusStatus string `hl7:"POS=17"`
	ChampusRetireDate time.Time `hl7:"POS=18;ATR=date"`
	ChampusNonAvailabilityCertificationOnFile string `hl7:"POS=19"`
	BabyCoverage string `hl7:"POS=20"`
	CombineBabyBill string `hl7:"POS=21"`
	BloodDeductible *int `hl7:"POS=22"`
	SpecialCoverageApprovalName PN `hl7:"POS=23"`
	SpecialCoverageApprovalTitle string `hl7:"POS=24"`
	NonCoveredInsuranceCode []string `hl7:"POS=25"`
	PayorID string `hl7:"POS=26"`
	PayorSubscriberID string `hl7:"POS=27"`
	EligibilitySource string `hl7:"POS=28"`
	RoomCoverageTypeAmount []CM_RMC `hl7:"POS=29"`
	PolicyTypeAmount []CM_PTA `hl7:"POS=30"`
	DailyDeductible CM_DDI `hl7:"POS=31"`
}

