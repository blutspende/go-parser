package hl7v24

import "time"

// NK1 - Next of kin / associated parties
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NK1
type NK1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	NkName []XPN `hl7:"POS=3"`
	Relationship CE `hl7:"POS=4"`
	Address []XAD `hl7:"POS=5"`
	PhoneNumber []XTN `hl7:"POS=6"`
	BusinessPhoneNumber []XTN `hl7:"POS=7"`
	ContactRole CE `hl7:"POS=8"`
	StartDate time.Time `hl7:"POS=9;ATR=date"`
	EndDate time.Time `hl7:"POS=10;ATR=date"`
	NextOfKinAssociatedPartiesJobTitle string `hl7:"POS=11"`
	NextOfKinAssociatedPartiesJobCodeClass JCC `hl7:"POS=12"`
	NextOfKinAssociatedPartiesEmployeeNumber CX `hl7:"POS=13"`
	OrganizationNameNk1 []XON `hl7:"POS=14"`
	MaritalStatus CE `hl7:"POS=15"`
	AdministrativeSex string `hl7:"POS=16"`
	DateTimeOfBirth time.Time `hl7:"POS=17"`
	LivingDependency []string `hl7:"POS=18"`
	AmbulatoryStatus []string `hl7:"POS=19"`
	Citizenship []CE `hl7:"POS=20"`
	PrimaryLanguage CE `hl7:"POS=21"`
	LivingArrangement string `hl7:"POS=22"`
	PublicityCode CE `hl7:"POS=23"`
	ProtectionIndicator string `hl7:"POS=24"`
	StudentIndicator string `hl7:"POS=25"`
	Religion CE `hl7:"POS=26"`
	MothersMaidenName []XPN `hl7:"POS=27"`
	Nationality CE `hl7:"POS=28"`
	EthnicGroup []CE `hl7:"POS=29"`
	ContactReason []CE `hl7:"POS=30"`
	ContactPersonsName []XPN `hl7:"POS=31"`
	ContactPersonsTelephoneNumber []XTN `hl7:"POS=32"`
	ContactPersonsAddress []XAD `hl7:"POS=33"`
	NextOfKinAssociatedPartysIdentifiers []CX `hl7:"POS=34"`
	JobStatus string `hl7:"POS=35"`
	Race []CE `hl7:"POS=36"`
	Handicap string `hl7:"POS=37"`
	ContactPersonSocialSecurityNumber string `hl7:"POS=38"`
}

