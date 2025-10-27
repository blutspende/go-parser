package hl7v23

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PID
type PID struct {
	ID                      int       `hl7:"POS=2" json:"id,omitempty"`
	ExternalID              []CX      `hl7:"POS=3" json:"externalID,omitempty"`
	InternalID              []CX      `hl7:"POS=4" json:"internalID,omitempty"`
	AlternateID             []CX      `hl7:"POS=5" json:"alternateID,omitempty"`
	Name                    XPN       `hl7:"POS=6" json:"name,omitempty"`
	MothersMaidenName       XPN       `hl7:"POS=7" json:"mothersMaidenName,omitempty"`
	DOB                     time.Time `hl7:"POS=8;ATR=date" json:"dob,omitempty"`
	Sex                     string    `hl7:"POS=9" json:"sex,omitempty"`
	Alias                   []XPN     `hl7:"POS=10" json:"alias,omitempty"`
	Race                    string    `hl7:"POS=11" json:"race,omitempty"`
	Address                 []XAD     `hl7:"POS=12" json:"address,omitempty"`
	CountryCode             string    `hl7:"POS=13" json:"countryCode,omitempty"`
	PhoneNumber             []XTN     `hl7:"POS=14" json:"phoneNumber,omitempty"`
	PhoneNumberBusiness     []XTN     `hl7:"POS=15" json:"phoneNumberBusiness,omitempty"`
	PrimaryLanguage         CE        `hl7:"POS=16" json:"primaryLanguage,omitempty"`
	MaritalStatus           string    `hl7:"POS=17" json:"maritalStatus,omitempty"`
	Religion                string    `hl7:"POS=18" json:"religion,omitempty"`
	PatientAccountNumber    CX        `hl7:"POS=19" json:"patientAccountNumber,omitempty"`
	SSNNumber               string    `hl7:"POS=20" json:"ssnNumber,omitempty"`
	DriversLicenseNumber    DLN       `hl7:"POS=21" json:"driversLicenseNumber,omitempty"`
	MothersIdentifier       CX        `hl7:"POS=22" json:"mothersIdentifier,omitempty"`
	EthnicGroup             string    `hl7:"POS=23" json:"ethnicGroup,omitempty"`
	BirthPlace              string    `hl7:"POS=24" json:"birthPlace,omitempty"`
	MultipleBirthIndicator  string    `hl7:"POS=25" json:"multipleBirthIndicator,omitempty"`
	BirthOrder              *int      `hl7:"POS=26" json:"birthOrder,omitempty"`
	Citizenship             []string  `hl7:"POS=27" json:"citizenship,omitempty"`
	VeteransMilitaryStatus  CE        `hl7:"POS=28" json:"veteransMilitaryStatus,omitempty"`
	NationalityCode         CE        `hl7:"POS=29" json:"nationalityCode,omitempty"`
	PatientDeathDateAndTime time.Time `hl7:"POS=30" json:"patientDeathDateAndTime,omitempty"`
	PatientDeathIndicator   string    `hl7:"POS=31" json:"patientDeathIndicator,omitempty"`
}
