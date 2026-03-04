package hl7v231

import "time"

// STF - Staff identification segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/STF
type STF struct {
	PrimaryKeyValueStf CE `hl7:"POS=2"`
	StaffIDCode []CX `hl7:"POS=3"`
	StaffName []XPN `hl7:"POS=4"`
	StaffType []string `hl7:"POS=5"`
	Sex string `hl7:"POS=6"`
	DateTimeOfBirth time.Time `hl7:"POS=7"`
	ActiveInactiveFlag string `hl7:"POS=8"`
	Department []CE `hl7:"POS=9"`
	HospitalService []CE `hl7:"POS=10"`
	Phone []XTN `hl7:"POS=11"`
	OfficeHomeAddress []XAD `hl7:"POS=12"`
	InstitutionActivationDate []DIN `hl7:"POS=13"`
	InstitutionInactivationDate []DIN `hl7:"POS=14"`
	BackupPersonID []CE `hl7:"POS=15"`
	EMailAddress []string `hl7:"POS=16"`
	PreferredMethodOfContact CE `hl7:"POS=17"`
	MaritalStatus CE `hl7:"POS=18"`
	JobTitle string `hl7:"POS=19"`
	JobCodeClass JCC `hl7:"POS=20"`
	EmploymentStatus string `hl7:"POS=21"`
	AdditionalInsuredOnAuto string `hl7:"POS=22"`
	DriversLicenseNumberStaff DLN `hl7:"POS=23"`
	CopyAutoIns string `hl7:"POS=24"`
	AutoInsExpires time.Time `hl7:"POS=25;ATR=date"`
	DateLastDmvReview time.Time `hl7:"POS=26;ATR=date"`
	DateNextDmvReview time.Time `hl7:"POS=27;ATR=date"`
}

