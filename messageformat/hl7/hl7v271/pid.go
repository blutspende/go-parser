package hl7v271

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/PID
type PID struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	PatientIdentifierList []CX `hl7:"POS=4"`
	PatientName []XPN `hl7:"POS=6"`
	MothersMaidenName []XPN `hl7:"POS=7"`
	DateTimeOfBirth time.Time `hl7:"POS=8"`
	AdministrativeSex CWE `hl7:"POS=9"`
	Race []CWE `hl7:"POS=11"`
	PatientAddress []XAD `hl7:"POS=12"`
	PhoneNumberHome []XTN `hl7:"POS=14"`
	PhoneNumberBusiness []XTN `hl7:"POS=15"`
	PrimaryLanguage CWE `hl7:"POS=16"`
	MaritalStatus CWE `hl7:"POS=17"`
	Religion CWE `hl7:"POS=18"`
	PatientAccountNumber CX `hl7:"POS=19"`
	MothersIdentifier []CX `hl7:"POS=22"`
	EthnicGroup []CWE `hl7:"POS=23"`
	BirthPlace string `hl7:"POS=24"`
	MultipleBirthIndicator string `hl7:"POS=25"`
	BirthOrder *int `hl7:"POS=26"`
	Citizenship []CWE `hl7:"POS=27"`
	VeteransMilitaryStatus CWE `hl7:"POS=28"`
	PatientDeathDateAndTime time.Time `hl7:"POS=30"`
	PatientDeathIndicator string `hl7:"POS=31"`
	IdentityUnknownIndicator string `hl7:"POS=32"`
	IdentityReliabilityCode []CWE `hl7:"POS=33"`
	LastUpdateDateTime time.Time `hl7:"POS=34"`
	LastUpdateFacility HD `hl7:"POS=35"`
	SpeciesCode CWE `hl7:"POS=36"`
	BreedCode CWE `hl7:"POS=37"`
	Strain string `hl7:"POS=38"`
	ProductionClassCode []CWE `hl7:"POS=39"`
	TribalCitizenship []CWE `hl7:"POS=40"`
	PatientTelecommunicationInformation []XTN `hl7:"POS=41"`
}

