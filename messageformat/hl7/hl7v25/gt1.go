package hl7v25

import "time"

// GT1 - Guarantor
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/GT1
type GT1 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	GuarantorNumber []CX `hl7:"POS=3"`
	GuarantorName []XPN `hl7:"POS=4"`
	GuarantorSpouseName []XPN `hl7:"POS=5"`
	GuarantorAddress []XAD `hl7:"POS=6"`
	GuarantorPhNumHome []XTN `hl7:"POS=7"`
	GuarantorPhNumBusiness []XTN `hl7:"POS=8"`
	GuarantorDateTimeOfBirth time.Time `hl7:"POS=9"`
	GuarantorAdministrativeSex string `hl7:"POS=10"`
	GuarantorType string `hl7:"POS=11"`
	GuarantorRelationship CE `hl7:"POS=12"`
	GuarantorSsn string `hl7:"POS=13"`
	GuarantorDateBegin time.Time `hl7:"POS=14;ATR=date"`
	GuarantorDateEnd time.Time `hl7:"POS=15;ATR=date"`
	GuarantorPriority *int `hl7:"POS=16"`
	GuarantorEmployerName []XPN `hl7:"POS=17"`
	GuarantorEmployerAddress []XAD `hl7:"POS=18"`
	GuarantorEmployerPhoneNumber []XTN `hl7:"POS=19"`
	GuarantorEmployeeIDNumber []CX `hl7:"POS=20"`
	GuarantorEmploymentStatus string `hl7:"POS=21"`
	GuarantorOrganizationName []XON `hl7:"POS=22"`
	GuarantorBillingHoldFlag string `hl7:"POS=23"`
	GuarantorCreditRatingCode CE `hl7:"POS=24"`
	GuarantorDeathDateAndTime time.Time `hl7:"POS=25"`
	GuarantorDeathFlag string `hl7:"POS=26"`
	GuarantorChargeAdjustmentCode CE `hl7:"POS=27"`
	GuarantorHouseholdAnnualIncome CP `hl7:"POS=28"`
	GuarantorHouseholdSize *int `hl7:"POS=29"`
	GuarantorEmployerIDNumber []CX `hl7:"POS=30"`
	GuarantorMaritalStatusCode CE `hl7:"POS=31"`
	GuarantorHireEffectiveDate time.Time `hl7:"POS=32;ATR=date"`
	EmploymentStopDate time.Time `hl7:"POS=33;ATR=date"`
	LivingDependency string `hl7:"POS=34"`
	AmbulatoryStatus []string `hl7:"POS=35"`
	Citizenship []CE `hl7:"POS=36"`
	PrimaryLanguage CE `hl7:"POS=37"`
	LivingArrangement string `hl7:"POS=38"`
	PublicityCode CE `hl7:"POS=39"`
	ProtectionIndicator string `hl7:"POS=40"`
	StudentIndicator string `hl7:"POS=41"`
	Religion CE `hl7:"POS=42"`
	MothersMaidenName []XPN `hl7:"POS=43"`
	Nationality CE `hl7:"POS=44"`
	EthnicGroup []CE `hl7:"POS=45"`
	ContactPersonsName []XPN `hl7:"POS=46"`
	ContactPersonsTelephoneNumber []XTN `hl7:"POS=47"`
	ContactReason CE `hl7:"POS=48"`
	ContactRelationship string `hl7:"POS=49"`
	JobTitle string `hl7:"POS=50"`
	JobCodeClass JCC `hl7:"POS=51"`
	GuarantorEmployersOrganizationName []XON `hl7:"POS=52"`
	Handicap string `hl7:"POS=53"`
	JobStatus string `hl7:"POS=54"`
	GuarantorFinancialClass FC `hl7:"POS=55"`
	GuarantorRace []CE `hl7:"POS=56"`
	GuarantorBirthPlace string `hl7:"POS=57"`
	VipIndicator string `hl7:"POS=58"`
}

