package hl7v25

import "time"

// STF - Staff Identification
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/STF
type STF struct {
	PrimaryKeyValueStf CE `hl7:"POS=2"`
	StaffIdentifierList []CX `hl7:"POS=3"`
	StaffName []XPN `hl7:"POS=4"`
	StaffType []string `hl7:"POS=5"`
	AdministrativeSex string `hl7:"POS=6"`
	DateTimeOfBirth time.Time `hl7:"POS=7"`
	ActiveInactiveFlag string `hl7:"POS=8"`
	Department []CE `hl7:"POS=9"`
	HospitalServiceStf []CE `hl7:"POS=10"`
	Phone []XTN `hl7:"POS=11"`
	OfficeHomeAddressBirthplace []XAD `hl7:"POS=12"`
	InstitutionActivationDate []DIN `hl7:"POS=13"`
	InstitutionInactivationDate []DIN `hl7:"POS=14"`
	BackupPersonID []CE `hl7:"POS=15"`
	EMailAddress []string `hl7:"POS=16"`
	PreferredMethodOfContact CE `hl7:"POS=17"`
	MaritalStatus CE `hl7:"POS=18"`
	JobTitle string `hl7:"POS=19"`
	JobCodeClass JCC `hl7:"POS=20"`
	EmploymentStatusCode CE `hl7:"POS=21"`
	AdditionalInsuredOnAuto string `hl7:"POS=22"`
	DriversLicenseNumberStaff DLN `hl7:"POS=23"`
	CopyAutoIns string `hl7:"POS=24"`
	AutoInsExpires time.Time `hl7:"POS=25;ATR=date"`
	DateLastDmvReview time.Time `hl7:"POS=26;ATR=date"`
	DateNextDmvReview time.Time `hl7:"POS=27;ATR=date"`
	Race CE `hl7:"POS=28"`
	EthnicGroup CE `hl7:"POS=29"`
	ReActivationApprovalIndicator string `hl7:"POS=30"`
	Citizenship []CWE `hl7:"POS=31"`
	DeathDateAndTime time.Time `hl7:"POS=32"`
	DeathIndicator string `hl7:"POS=33"`
	InstitutionRelationshipTypeCode CWE `hl7:"POS=34"`
	InstitutionRelationshipPeriod DR `hl7:"POS=35"`
	ExpectedReturnDate time.Time `hl7:"POS=36;ATR=date"`
	CostCenterCode []CWE `hl7:"POS=37"`
	GenericClassificationIndicator string `hl7:"POS=38"`
	InactiveReasonCode CWE `hl7:"POS=39"`
}

