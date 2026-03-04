package hl7v27

import "time"

// IN2 - Insurance Additional Information
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/IN2
type IN2 struct {
	InsuredsEmployeeID []CX `hl7:"POS=2"`
	InsuredsSocialSecurityNumber string `hl7:"POS=3"`
	InsuredsEmployersNameAndID []XCN `hl7:"POS=4"`
	EmployerInformationData CWE `hl7:"POS=5"`
	MailClaimParty []CWE `hl7:"POS=6"`
	MedicareHealthInsCardNumber string `hl7:"POS=7"`
	MedicaidCaseName []XPN `hl7:"POS=8"`
	MedicaidCaseNumber string `hl7:"POS=9"`
	MilitarySponsorName []XPN `hl7:"POS=10"`
	MilitaryIDNumber string `hl7:"POS=11"`
	DependentOfMilitaryRecipient CWE `hl7:"POS=12"`
	MilitaryOrganization string `hl7:"POS=13"`
	MilitaryStation string `hl7:"POS=14"`
	MilitaryService CWE `hl7:"POS=15"`
	MilitaryRankGrade CWE `hl7:"POS=16"`
	MilitaryStatus CWE `hl7:"POS=17"`
	MilitaryRetireDate time.Time `hl7:"POS=18;ATR=date"`
	MilitaryNonAvailCertOnFile string `hl7:"POS=19"`
	BabyCoverage string `hl7:"POS=20"`
	CombineBabyBill string `hl7:"POS=21"`
	BloodDeductible string `hl7:"POS=22"`
	SpecialCoverageApprovalName []XPN `hl7:"POS=23"`
	SpecialCoverageApprovalTitle string `hl7:"POS=24"`
	NonCoveredInsuranceCode []CWE `hl7:"POS=25"`
	PayorID []CX `hl7:"POS=26"`
	PayorSubscriberID []CX `hl7:"POS=27"`
	EligibilitySource CWE `hl7:"POS=28"`
	RoomCoverageTypeAmount []RMC `hl7:"POS=29"`
	PolicyTypeAmount []PTA `hl7:"POS=30"`
	DailyDeductible DDI `hl7:"POS=31"`
	LivingDependency CWE `hl7:"POS=32"`
	AmbulatoryStatus []CWE `hl7:"POS=33"`
	Citizenship []CWE `hl7:"POS=34"`
	PrimaryLanguage CWE `hl7:"POS=35"`
	LivingArrangement CWE `hl7:"POS=36"`
	PublicityCode CWE `hl7:"POS=37"`
	ProtectionIndicator string `hl7:"POS=38"`
	StudentIndicator CWE `hl7:"POS=39"`
	Religion CWE `hl7:"POS=40"`
	MothersMaidenName []XPN `hl7:"POS=41"`
	Nationality CWE `hl7:"POS=42"`
	EthnicGroup []CWE `hl7:"POS=43"`
	MaritalStatus []CWE `hl7:"POS=44"`
	InsuredsEmploymentStartDate time.Time `hl7:"POS=45;ATR=date"`
	EmploymentStopDate time.Time `hl7:"POS=46;ATR=date"`
	JobTitle string `hl7:"POS=47"`
	JobCodeClass JCC `hl7:"POS=48"`
	JobStatus CWE `hl7:"POS=49"`
	EmployerContactPersonName []XPN `hl7:"POS=50"`
	EmployerContactPersonPhoneNumber []XTN `hl7:"POS=51"`
	EmployerContactReason CWE `hl7:"POS=52"`
	InsuredsContactPersonsName []XPN `hl7:"POS=53"`
	InsuredsContactPersonPhoneNumber []XTN `hl7:"POS=54"`
	InsuredsContactPersonReason []CWE `hl7:"POS=55"`
	RelationshipToThePatientStartDate time.Time `hl7:"POS=56;ATR=date"`
	RelationshipToThePatientStopDate []time.Time `hl7:"POS=57;ATR=date"`
	InsuranceCoContactReason CWE `hl7:"POS=58"`
	InsuranceCoContactPhoneNumber []XTN `hl7:"POS=59"`
	PolicyScope CWE `hl7:"POS=60"`
	PolicySource CWE `hl7:"POS=61"`
	PatientMemberNumber CX `hl7:"POS=62"`
	GuarantorsRelationshipToInsured CWE `hl7:"POS=63"`
	InsuredsPhoneNumberHome []XTN `hl7:"POS=64"`
	InsuredsEmployerPhoneNumber []XTN `hl7:"POS=65"`
	MilitaryHandicappedProgram CWE `hl7:"POS=66"`
	SuspendFlag string `hl7:"POS=67"`
	CopayLimitFlag string `hl7:"POS=68"`
	StopLossLimitFlag string `hl7:"POS=69"`
	InsuredOrganizationNameAndID []XON `hl7:"POS=70"`
	InsuredEmployerOrganizationNameAndID []XON `hl7:"POS=71"`
	Race []CWE `hl7:"POS=72"`
	PatientsRelationshipToInsured CWE `hl7:"POS=73"`
}

