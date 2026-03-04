package hl7v22

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/PID
type PID struct {
	SetID                       int           `hl7:"POS=2;ATR=sequence"`
	ExternalID                  CK            `hl7:"POS=3"`
	InternalID                  []CM_PAT_ID   `hl7:"POS=4"`
	AlternatePatientID          string        `hl7:"POS=5"`
	PatientName                 PN            `hl7:"POS=6"`
	MothersMaidenName           string        `hl7:"POS=7"`
	DateOfBirth                 time.Time     `hl7:"POS=8"`
	Sex                         string        `hl7:"POS=9"`
	PatientAlias                []PN          `hl7:"POS=10"`
	Race                        string        `hl7:"POS=11"`
	PatientAddress              []AD          `hl7:"POS=12"`
	CountyCode                  string        `hl7:"POS=13"`
	PhoneNumberHome             []string      `hl7:"POS=14"`
	PhoneNumberBusiness         []string      `hl7:"POS=15"`
	LanguagePatient             string        `hl7:"POS=16"`
	MaritalStatus               string        `hl7:"POS=17"`
	Religion                    string        `hl7:"POS=18"`
	PatientAccountNumber        CK            `hl7:"POS=19"`
	SocialSecurityNumberPatient string        `hl7:"POS=20"`
	DriversLicenseNumberPatient CM_LICENSE_NO `hl7:"POS=21"`
	MothersIdentifier           CK            `hl7:"POS=22"`
	EthnicGroup                 string        `hl7:"POS=23"`
	BirthPlace                  string        `hl7:"POS=24"`
	MultipleBirthIndicator      string        `hl7:"POS=25"`
	BirthOrder                  *int          `hl7:"POS=26"`
	Citizenship                 []string      `hl7:"POS=27"`
	VeteransMilitaryStatus      CE            `hl7:"POS=28"`
}
