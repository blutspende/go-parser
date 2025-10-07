package hl7v23

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PID
type PID struct {
	ID                      int       `hl7:"1" json:"id,omitempty"`
	ExternalID              []CX      `hl7:"2" json:"externalID,omitempty"`
	InternalID              []CX      `hl7:"3" json:"internalID,omitempty"`
	AlternateID             []CX      `hl7:"4" json:"alternateID,omitempty"`
	Name                    XPN       `hl7:"5" json:"name,omitempty"`
	MothersMaidenName       XPN       `hl7:"6" json:"mothersMaidenName,omitempty"`
	DOB                     time.Time `hl7:"7" json:"dob,omitempty"`
	Sex                     string    `hl7:"8" json:"sex,omitempty"`
	Alias                   []XPN     `hl7:"9" json:"alias,omitempty"`
	Race                    string    `hl7:"10" json:"race,omitempty"`
	Address                 []XAD     `hl7:"11" json:"address,omitempty"`
	CountryCode             string    `hl7:"12" json:"countryCode,omitempty"`
	PhoneNumber             []XTN     `hl7:"13" json:"phoneNumber,omitempty"`
	PhoneNumberBusiness     []XTN     `hl7:"14" json:"phoneNumberBusiness,omitempty"`
	PrimaryLanguage         CE        `hl7:"15" json:"primaryLanguage,omitempty"`
	MaritalStatus           string    `hl7:"16" json:"maritalStatus,omitempty"`
	Religion                string    `hl7:"17" json:"religion,omitempty"`
	PatientAccountNumber    CX        `hl7:"18" json:"patientAccountNumber,omitempty"`
	SSNNumber               string    `hl7:"19" json:"ssnNumber,omitempty"`
	DriversLicenseNumber    DLN       `hl7:"20" json:"driversLicenseNumber,omitempty"`
	MothersIdentifier       CX        `hl7:"21" json:"mothersIdentifier,omitempty"`
	EthnicGroup             string    `hl7:"22" json:"ethnicGroup,omitempty"`
	BirthPlace              string    `hl7:"23" json:"birthPlace,omitempty"`
	MultipleBirthIndicator  string    `hl7:"24" json:"multipleBirthIndicator,omitempty"`
	BirthOrder              int       `hl7:"25" json:"birthOrder,omitempty"`
	Citizenship             []string  `hl7:"26" json:"citizenship,omitempty"`
	VeteransMilitaryStatus  CE        `hl7:"27" json:"veteransMilitaryStatus,omitempty"`
	NationalityCode         CE        `hl7:"28" json:"nationalityCode,omitempty"`
	PatientDeathDateAndTime time.Time `hl7:"29" json:"patientDeathDateAndTime,omitempty"`
	PatientDeathIndicator   string    `hl7:"30" json:"patientDeathIndicator,omitempty"`
}
