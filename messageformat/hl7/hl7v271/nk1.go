package hl7v271

import "time"

// NK1 - Next Of Kin / Associated Parties
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/NK1
type NK1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	Name []XPN `hl7:"POS=3"`
	Relationship CWE `hl7:"POS=4"`
	Address []XAD `hl7:"POS=5"`
	PhoneNumber []XTN `hl7:"POS=6"`
	BusinessPhoneNumber []XTN `hl7:"POS=7"`
	ContactRole CWE `hl7:"POS=8"`
	StartDate time.Time `hl7:"POS=9;ATR=date"`
	EndDate time.Time `hl7:"POS=10;ATR=date"`
	NextOfKinAssociatedPartiesJobTitle string `hl7:"POS=11"`
	NextOfKinAssociatedPartiesJobCodeClass JCC `hl7:"POS=12"`
	NextOfKinAssociatedPartiesEmployeeNumber CX `hl7:"POS=13"`
	OrganizationNameNk1 []XON `hl7:"POS=14"`
	MaritalStatus CWE `hl7:"POS=15"`
	AdministrativeSex CWE `hl7:"POS=16"`
	DateTimeOfBirth time.Time `hl7:"POS=17"`
	LivingDependency []CWE `hl7:"POS=18"`
	AmbulatoryStatus []CWE `hl7:"POS=19"`
	Citizenship []CWE `hl7:"POS=20"`
	PrimaryLanguage CWE `hl7:"POS=21"`
	LivingArrangement CWE `hl7:"POS=22"`
	PublicityCode CWE `hl7:"POS=23"`
	ProtectionIndicator string `hl7:"POS=24"`
	StudentIndicator CWE `hl7:"POS=25"`
	Religion CWE `hl7:"POS=26"`
	MothersMaidenName []XPN `hl7:"POS=27"`
	Nationality CWE `hl7:"POS=28"`
	EthnicGroup []CWE `hl7:"POS=29"`
	ContactReason []CWE `hl7:"POS=30"`
	ContactPersonsName []XPN `hl7:"POS=31"`
	ContactPersonsTelephoneNumber []XTN `hl7:"POS=32"`
	ContactPersonsAddress []XAD `hl7:"POS=33"`
	NextOfKinAssociatedPartysIdentifiers []CX `hl7:"POS=34"`
	JobStatus CWE `hl7:"POS=35"`
	Race []CWE `hl7:"POS=36"`
	Handicap CWE `hl7:"POS=37"`
	ContactPersonSocialSecurityNumber string `hl7:"POS=38"`
	NextOfKinBirthPlace string `hl7:"POS=39"`
	VipIndicator CWE `hl7:"POS=40"`
	NextOfKinTelecommunicationInformation XTN `hl7:"POS=41"`
	ContactPersonsTelecommunicationInformation XTN `hl7:"POS=42"`
}

