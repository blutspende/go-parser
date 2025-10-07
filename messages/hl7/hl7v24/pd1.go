package hl7v24

import "time"

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PD1
type PD1 struct {
	LivingDependency                  []string  `hl7:"1" json:"LivingDependency,omitempty"`
	LivingArrangement                 string    `hl7:"2" json:"LivingArrangement,omitempty"`
	PrimaryFacility                   []XON     `hl7:"3" json:"PrimaryFacility,omitempty"`
	PrimaryCarePrvoider               []XCN     `hl7:"4" json:"PrimaryCarePrvoider,omitempty"`
	StudentIndicator                  string    `hl7:"5" json:"StudentIndicator,omitempty"`
	Handicap                          string    `hl7:"6" json:"Handicap,omitempty"`
	LivingWill                        string    `hl7:"7" json:"LivingWill,omitempty"`
	OrganDonor                        string    `hl7:"8" json:"OrganDonor,omitempty"`
	SeparateBill                      string    `hl7:"9" json:"SeparateBill,omitempty"`
	DuplicatePatient                  []CX      `hl7:"10" json:"DuplicatePatient,omitempty"`
	PublicityIndicator                CE        `hl7:"11" json:"PublicityIndicator,omitempty"`
	ProtectionIndicator               string    `hl7:"12" json:"ProtectionIndicator,omitempty"`
	ProtectionIndicatorEffectiveDate  time.Time `hl7:"13" json:"ProtectionIndicatorEffectiveDate,omitempty"`
	PlaceOfWorship                    []XON     `hl7:"14" json:"PlaceOfWorship,omitempty"`
	AdvanceDirectiveCode              []CE      `hl7:"15" json:"AdvanceDirectiveCode,omitempty"`
	ImmunizationRegistryStatus        string    `hl7:"16" json:"ImmunizationRegistryStatus,omitempty"`
	ImmunizationRegistryEffectiveDate time.Time `hl7:"17" json:"ImmunizationRegistryEffectiveDate,omitempty"`
	PublicityCodeEffectiveDate        time.Time `hl7:"18" json:"PublicityCodeEffectiveDate,omitempty"`
	MilitaryBranch                    string    `hl7:"19" json:"MilitaryBranch,omitempty"`
	MilitaryRank                      string    `hl7:"10" json:"MilitaryRank,omitempty"`
	MilitaryStatus                    string    `hl7:"21" json:"MilitaryStatus,omitempty"`
}
