package hl7v24

import "time"

// HL7 v2.4 - GT1 - Guarantor
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/GT1
type GT1 struct {
	SetId                         string    `hl7:"1" json:"SetId,omitempty"`
	Number                        []CX      `hl7:"2" json:"Number,omitempty"`
	Name                          []XPN     `hl7:"3" json:"Name,omitempty"`
	SpouseName                    []XPN     `hl7:"4" json:"SpouseName,omitempty"`
	Address                       []XAD     `hl7:"5" json:"Address,omitempty"`
	PhoneNumberHome               []XTN     `hl7:"6" json:"PhoneNumberHome,omitempty"`
	PhoneNumberBusiness           []XTN     `hl7:"7" json:"PhoneNumberBusiness,omitempty"`
	DateTimeOfBirth               time.Time `hl7:"8,longdate" json:"DateTimeOfBirth,omitempty"`
	Sex                           string    `hl7:"9" json:"Sex,omitempty"`
	Type                          string    `hl7:"10" json:"Type,omitempty"`
	Relationship                  CE        `hl7:"11" json:"Relationship,omitempty"`
	SSN                           string    `hl7:"12" json:"SSN,omitempty"`
	DateBegin                     time.Time `hl7:"13,shortdate" json:"DateBegin,omitempty"`
	DateEnd                       time.Time `hl7:"14,shortdate" json:"DateEnd,omitempty"`
	Priority                      float32   `hl7:"15" json:"Priority,omitempty"`
	EmployerName                  []XPN     `hl7:"16" json:"EmployerName,omitempty"`
	EmployerAddress               []XAD     `hl7:"17" json:"EmployerAddress,omitempty"`
	EmployerPhoneNumber           []XTN     `hl7:"18" json:"EmployerPhoneNumber,omitempty"`
	EmployeeIDNumber              []CX      `hl7:"19" json:"EmployeeIDNumber,omitempty"`
	EmploymentStatus              string    `hl7:"20" json:"EmploymentStatus,omitempty"`
	OrganizationName              []XON     `hl7:"21" json:"OrganizationName,omitempty"`
	BillingHoldFlag               string    `hl7:"22" json:"BillingHoldFlag,omitempty"`
	CreditRatingCode              CE        `hl7:"23" json:"CreditRatingCode,omitempty"`
	DeathDateAndTime              time.Time `hl7:"24,longdate" json:"DeathDateAndTime,omitempty"`
	DeathFlag                     string    `hl7:"25" json:"DeathFlag,omitempty"`
	ChargeAdjustmentCode          CE        `hl7:"26" json:"ChargeAdjustmentCode,omitempty"`
	HouseholdAnnualIncome         CP        `hl7:"27" json:"HouseholdAnnualIncome,omitempty"`
	HouseholdSize                 float32   `hl7:"28" json:"HouseholdSize,omitempty"`
	EmployerIDNumber              []CX      `hl7:"29" json:"EmployerIDNumber,omitempty"`
	MaritalStatusCode             CE        `hl7:"30" json:"MaritalStatusCode,omitempty"`
	HireEffectiveDate             time.Time `hl7:"31,shortdate" json:"HireEffectiveDate,omitempty"`
	EmploymentStopDAte            time.Time `hl7:"32,shortdate" json:"EmploymentStopDAte,omitempty"`
	LivingDependency              string    `hl7:"33" json:"LivingDependency,omitempty"`
	AmbulatoryStatusCode          []string  `hl7:"34" json:"AmbulatoryStatusCode,omitempty"`
	Citizenship                   []CE      `hl7:"35" json:"Citizenship,omitempty"`
	PrimaryLanguage               CE        `hl7:"36" json:"PrimaryLanguage,omitempty"`
	LivingArrangement             string    `hl7:"37" json:"LivingArrangement,omitempty"`
	PublicityIndicator            CE        `hl7:"38" json:"PublicityIndicator,omitempty"`
	ProtectionIndicator           string    `hl7:"39" json:"ProtectionIndicator,omitempty"`
	StudentIndicator              string    `hl7:"40" json:"StudentIndicator,omitempty"`
	Religion                      CE        `hl7:"41" json:"Religion,omitempty"`
	MothersMaidenName             []XPN     `hl7:"42" json:"MothersMaidenName,omitempty"`
	Nationality                   CE        `hl7:"43" json:"Nationality,omitempty"`
	EthnicGroup                   []CE      `hl7:"44" json:"EthnicGroup,omitempty"`
	ContactPersonsName            []XPN     `hl7:"45" json:"ContactPersonsName,omitempty"`
	ContactPersonsTelephoneNumber []XTN     `hl7:"46" json:"ContactPersonsTelephoneNumber,omitempty"`
	ContactReason                 CE        `hl7:"47" json:"ContactReason,omitempty"`
	ContactRelationshipCode       string    `hl7:"48" json:"ContactRelationshipCode,omitempty"`
	JobTitle                      string    `hl7:"49" json:"JobTitle,omitempty"`
	JobCode                       JCC       `hl7:"50" json:"JobCode,omitempty"`
	EmployersOrganizationName     []XON     `hl7:"51" json:"EmployersOrganizationName,omitempty"`
	Handicap                      string    `hl7:"52" json:"Handicap,omitempty"`
	JobStatus                     string    `hl7:"53" json:"JobStatus,omitempty"`
	FinancialClass                FC        `hl7:"54" json:"FinancialClass,omitempty"`
	GuarantorRace                 []CE      `hl7:"55" json:"GuarantorRace,omitempty"`
}
