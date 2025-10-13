package hl7v24

import "time"

// HL7 v2.4 - GT1 - Guarantor
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/GT1
type GT1 struct {
	SetId                         string    `hl7:"2" json:"SetId,omitempty"`
	Number                        []CX      `hl7:"3" json:"Number,omitempty"`
	Name                          []XPN     `hl7:"4" json:"Name,omitempty"`
	SpouseName                    []XPN     `hl7:"5" json:"SpouseName,omitempty"`
	Address                       []XAD     `hl7:"6" json:"Address,omitempty"`
	PhoneNumberHome               []XTN     `hl7:"7" json:"PhoneNumberHome,omitempty"`
	PhoneNumberBusiness           []XTN     `hl7:"8" json:"PhoneNumberBusiness,omitempty"`
	DateTimeOfBirth               time.Time `hl7:"9,longdate" json:"DateTimeOfBirth,omitempty"`
	Sex                           string    `hl7:"10" json:"Sex,omitempty"`
	Type                          string    `hl7:"11" json:"Type,omitempty"`
	Relationship                  CE        `hl7:"12" json:"Relationship,omitempty"`
	SSN                           string    `hl7:"13" json:"SSN,omitempty"`
	DateBegin                     time.Time `hl7:"14" json:"DateBegin,omitempty"`
	DateEnd                       time.Time `hl7:"15" json:"DateEnd,omitempty"`
	Priority                      float32   `hl7:"16" json:"Priority,omitempty"`
	EmployerName                  []XPN     `hl7:"17" json:"EmployerName,omitempty"`
	EmployerAddress               []XAD     `hl7:"18" json:"EmployerAddress,omitempty"`
	EmployerPhoneNumber           []XTN     `hl7:"19" json:"EmployerPhoneNumber,omitempty"`
	EmployeeIDNumber              []CX      `hl7:"20" json:"EmployeeIDNumber,omitempty"`
	EmploymentStatus              string    `hl7:"21" json:"EmploymentStatus,omitempty"`
	OrganizationName              []XON     `hl7:"22" json:"OrganizationName,omitempty"`
	BillingHoldFlag               string    `hl7:"23" json:"BillingHoldFlag,omitempty"`
	CreditRatingCode              CE        `hl7:"24" json:"CreditRatingCode,omitempty"`
	DeathDateAndTime              time.Time `hl7:"25,longdate" json:"DeathDateAndTime,omitempty"`
	DeathFlag                     string    `hl7:"26" json:"DeathFlag,omitempty"`
	ChargeAdjustmentCode          CE        `hl7:"27" json:"ChargeAdjustmentCode,omitempty"`
	HouseholdAnnualIncome         CP        `hl7:"28" json:"HouseholdAnnualIncome,omitempty"`
	HouseholdSize                 float32   `hl7:"29" json:"HouseholdSize,omitempty"`
	EmployerIDNumber              []CX      `hl7:"30" json:"EmployerIDNumber,omitempty"`
	MaritalStatusCode             CE        `hl7:"31" json:"MaritalStatusCode,omitempty"`
	HireEffectiveDate             time.Time `hl7:"32" json:"HireEffectiveDate,omitempty"`
	EmploymentStopDAte            time.Time `hl7:"33" json:"EmploymentStopDAte,omitempty"`
	LivingDependency              string    `hl7:"34" json:"LivingDependency,omitempty"`
	AmbulatoryStatusCode          []string  `hl7:"35" json:"AmbulatoryStatusCode,omitempty"`
	Citizenship                   []CE      `hl7:"36" json:"Citizenship,omitempty"`
	PrimaryLanguage               CE        `hl7:"37" json:"PrimaryLanguage,omitempty"`
	LivingArrangement             string    `hl7:"38" json:"LivingArrangement,omitempty"`
	PublicityIndicator            CE        `hl7:"39" json:"PublicityIndicator,omitempty"`
	ProtectionIndicator           string    `hl7:"40" json:"ProtectionIndicator,omitempty"`
	StudentIndicator              string    `hl7:"41" json:"StudentIndicator,omitempty"`
	Religion                      CE        `hl7:"42" json:"Religion,omitempty"`
	MothersMaidenName             []XPN     `hl7:"43" json:"MothersMaidenName,omitempty"`
	Nationality                   CE        `hl7:"44" json:"Nationality,omitempty"`
	EthnicGroup                   []CE      `hl7:"45" json:"EthnicGroup,omitempty"`
	ContactPersonsName            []XPN     `hl7:"46" json:"ContactPersonsName,omitempty"`
	ContactPersonsTelephoneNumber []XTN     `hl7:"47" json:"ContactPersonsTelephoneNumber,omitempty"`
	ContactReason                 CE        `hl7:"48" json:"ContactReason,omitempty"`
	ContactRelationshipCode       string    `hl7:"49" json:"ContactRelationshipCode,omitempty"`
	JobTitle                      string    `hl7:"50" json:"JobTitle,omitempty"`
	JobCode                       JCC       `hl7:"51" json:"JobCode,omitempty"`
	EmployersOrganizationName     []XON     `hl7:"52" json:"EmployersOrganizationName,omitempty"`
	Handicap                      string    `hl7:"53" json:"Handicap,omitempty"`
	JobStatus                     string    `hl7:"54" json:"JobStatus,omitempty"`
	FinancialClass                FC        `hl7:"55" json:"FinancialClass,omitempty"`
	GuarantorRace                 []CE      `hl7:"56" json:"GuarantorRace,omitempty"`
}
