package hl7v231

import "time"

// PID - Patient identification segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/PID
type PID struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PatientID CX `hl7:"POS=3"`
	PatientIdentifierList []CX `hl7:"POS=4"`
	AlternatePatientIDPid []CX `hl7:"POS=5"`
	PatientName []XPN `hl7:"POS=6"`
	MothersMaidenName []XPN `hl7:"POS=7"`
	DateTimeOfBirth time.Time `hl7:"POS=8"`
	Sex string `hl7:"POS=9"`
	PatientAlias []XPN `hl7:"POS=10"`
	Race []CE `hl7:"POS=11"`
	PatientAddress []XAD `hl7:"POS=12"`
	CountyCode string `hl7:"POS=13"`
	PhoneNumberHome []XTN `hl7:"POS=14"`
	PhoneNumberBusiness []XTN `hl7:"POS=15"`
	PrimaryLanguage CE `hl7:"POS=16"`
	MaritalStatus CE `hl7:"POS=17"`
	Religion CE `hl7:"POS=18"`
	PatientAccountNumber CX `hl7:"POS=19"`
	SsnNumberPatient string `hl7:"POS=20"`
	DriversLicenseNumberPatient DLN `hl7:"POS=21"`
	MothersIdentifier []CX `hl7:"POS=22"`
	EthnicGroup []CE `hl7:"POS=23"`
	BirthPlace string `hl7:"POS=24"`
	MultipleBirthIndicator string `hl7:"POS=25"`
	BirthOrder *int `hl7:"POS=26"`
	Citizenship []CE `hl7:"POS=27"`
	VeteransMilitaryStatus CE `hl7:"POS=28"`
	Nationality CE `hl7:"POS=29"`
	PatientDeathDateAndTime time.Time `hl7:"POS=30"`
	PatientDeathIndicator string `hl7:"POS=31"`
}

