package hl7v24

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PID
type PID struct {
	ID                       int       `hl7:"1" json:"ID,omitempty"`
	ExternalID               []CX      `hl7:"2" json:"ExternalID,omitempty"`
	InternalID               []CX      `hl7:"3" json:"InternalID,omitempty"`
	AlternateID              []CX      `hl7:"4" json:"AlternateID,omitempty"`
	Name                     []XPN     `hl7:"5" json:"Name,omitempty"`
	MothersMaidenName        []XPN     `hl7:"6" json:"MothersMaidenName,omitempty"`
	DOB                      time.Time `hl7:"7" json:"DOB,omitempty"`
	Sex                      string    `hl7:"8" json:"Sex,omitempty"`
	Alias                    []XPN     `hl7:"9" json:"Alias,omitempty"`
	Race                     []CE      `hl7:"10" json:"Race,omitempty"`
	Address                  []XAD     `hl7:"11" json:"Address,omitempty"`
	CountryCode              string    `hl7:"12" json:"CountryCode,omitempty"`
	PhoneNumber              []XTN     `hl7:"13" json:"PhoneNumber,omitempty"`
	PhoneNumberBusiness      []XTN     `hl7:"14" json:"PhoneNumberBusiness,omitempty"`
	PrimaryLanguage          CE        `hl7:"15" json:"PrimaryLanguage,omitempty"`
	MaritalStatus            CE        `hl7:"16" json:"MaritalStatus,omitempty"`
	Religion                 CE        `hl7:"17" json:"Religion,omitempty"`
	PatientAccountNumber     CX        `hl7:"18" json:"PatientAccountNumber,omitempty"`
	SSNNumber                string    `hl7:"19" json:"SSNNumber,omitempty"`
	DriversLicenseNumber     DLN       `hl7:"20" json:"DriversLicenseNumber,omitempty"`
	MothersIdentifier        []CX      `hl7:"21" json:"MothersIdentifier,omitempty"`
	EthnicGroup              []string  `hl7:"22" json:"EthnicGroup,omitempty"`
	BirthPlace               string    `hl7:"23" json:"BirthPlace,omitempty"`
	MultipleBirthIndicator   string    `hl7:"24" json:"MultipleBirthIndicator,omitempty"`
	BirthOrder               int       `hl7:"25" json:"BirthOrder,omitempty"`
	Citizenship              []string  `hl7:"26" json:"Citizenship,omitempty"`
	VeteransMilitaryStatus   CE        `hl7:"27" json:"VeteransMilitaryStatus,omitempty"`
	NationalityCode          CE        `hl7:"28" json:"NationalityCode,omitempty"`
	PatientDeathDateAndTime  time.Time `hl7:"29" json:"PatientDeathDateAndTime,omitempty"`
	PatientDeathIndicator    string    `hl7:"30" json:"PatientDeathIndicator,omitempty"`
	IdentityUnknownIndicator string    `hl7:"31" json:"IdentityUnknownIndicator,omitempty"`
	IdentityReliabilityCode  string    `hl7:"32" json:"IdentityReliabilityCode,omitempty"`
	LastUpdateDateTime       time.Time `hl7:"33" json:"LastUpdateDateTime,omitempty"`
	LastUpdateFacility       HD        `hl7:"34" json:"LastUpdateFacility,omitempty"`
	SpeciesCode              CE        `hl7:"35" json:"SpeciesCode,omitempty"`
	BreedCode                CE        `hl7:"36" json:"BreedCode,omitempty"`
	Strain                   string    `hl7:"37" json:"Strain,omitempty"`
	ProductionClassCode      []CE      `hl7:"38" json:"ProductionClassCode,omitempty"`
}
