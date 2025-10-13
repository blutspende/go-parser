package hl7v24

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PID
type PID struct {
	ID                       int       `hl7:"2" json:"ID,omitempty"`
	ExternalID               []CX      `hl7:"3" json:"ExternalID,omitempty"`
	InternalID               []CX      `hl7:"4" json:"InternalID,omitempty"`
	AlternateID              []CX      `hl7:"5" json:"AlternateID,omitempty"`
	Name                     []XPN     `hl7:"6" json:"Name,omitempty"`
	MothersMaidenName        []XPN     `hl7:"7" json:"MothersMaidenName,omitempty"`
	DOB                      time.Time `hl7:"8" json:"DOB,omitempty"`
	Sex                      string    `hl7:"9" json:"Sex,omitempty"`
	Alias                    []XPN     `hl7:"10" json:"Alias,omitempty"`
	Race                     []CE      `hl7:"11" json:"Race,omitempty"`
	Address                  []XAD     `hl7:"12" json:"Address,omitempty"`
	CountryCode              string    `hl7:"13" json:"CountryCode,omitempty"`
	PhoneNumber              []XTN     `hl7:"14" json:"PhoneNumber,omitempty"`
	PhoneNumberBusiness      []XTN     `hl7:"15" json:"PhoneNumberBusiness,omitempty"`
	PrimaryLanguage          CE        `hl7:"16" json:"PrimaryLanguage,omitempty"`
	MaritalStatus            CE        `hl7:"17" json:"MaritalStatus,omitempty"`
	Religion                 CE        `hl7:"18" json:"Religion,omitempty"`
	PatientAccountNumber     CX        `hl7:"19" json:"PatientAccountNumber,omitempty"`
	SSNNumber                string    `hl7:"20" json:"SSNNumber,omitempty"`
	DriversLicenseNumber     DLN       `hl7:"21" json:"DriversLicenseNumber,omitempty"`
	MothersIdentifier        []CX      `hl7:"22" json:"MothersIdentifier,omitempty"`
	EthnicGroup              []string  `hl7:"23" json:"EthnicGroup,omitempty"`
	BirthPlace               string    `hl7:"24" json:"BirthPlace,omitempty"`
	MultipleBirthIndicator   string    `hl7:"25" json:"MultipleBirthIndicator,omitempty"`
	BirthOrder               int       `hl7:"26" json:"BirthOrder,omitempty"`
	Citizenship              []string  `hl7:"27" json:"Citizenship,omitempty"`
	VeteransMilitaryStatus   CE        `hl7:"28" json:"VeteransMilitaryStatus,omitempty"`
	NationalityCode          CE        `hl7:"29" json:"NationalityCode,omitempty"`
	PatientDeathDateAndTime  time.Time `hl7:"30" json:"PatientDeathDateAndTime,omitempty"`
	PatientDeathIndicator    string    `hl7:"31" json:"PatientDeathIndicator,omitempty"`
	IdentityUnknownIndicator string    `hl7:"32" json:"IdentityUnknownIndicator,omitempty"`
	IdentityReliabilityCode  string    `hl7:"33" json:"IdentityReliabilityCode,omitempty"`
	LastUpdateDateTime       time.Time `hl7:"34" json:"LastUpdateDateTime,omitempty"`
	LastUpdateFacility       HD        `hl7:"35" json:"LastUpdateFacility,omitempty"`
	SpeciesCode              CE        `hl7:"36" json:"SpeciesCode,omitempty"`
	BreedCode                CE        `hl7:"37" json:"BreedCode,omitempty"`
	Strain                   string    `hl7:"38" json:"Strain,omitempty"`
	ProductionClassCode      []CE      `hl7:"39" json:"ProductionClassCode,omitempty"`
}
