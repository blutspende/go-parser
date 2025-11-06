package hl7v24

import "time"

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PD1
type PD1 struct {
	LivingDependency                  []string  `hl7:"POS=2" json:"LivingDependency,omitempty"`
	LivingArrangement                 string    `hl7:"POS=3" json:"LivingArrangement,omitempty"`
	PrimaryFacility                   []XON     `hl7:"POS=4" json:"PrimaryFacility,omitempty"`
	PrimaryCarePrvoider               []XCN     `hl7:"POS=5" json:"PrimaryCarePrvoider,omitempty"`
	StudentIndicator                  string    `hl7:"POS=6" json:"StudentIndicator,omitempty"`
	Handicap                          string    `hl7:"POS=7" json:"Handicap,omitempty"`
	LivingWill                        string    `hl7:"POS=8" json:"LivingWill,omitempty"`
	OrganDonor                        string    `hl7:"POS=9" json:"OrganDonor,omitempty"`
	SeparateBill                      string    `hl7:"POS=10" json:"SeparateBill,omitempty"`
	DuplicatePatient                  []CX      `hl7:"POS=11" json:"DuplicatePatient,omitempty"`
	PublicityIndicator                CE        `hl7:"POS=12" json:"PublicityIndicator,omitempty"`
	ProtectionIndicator               string    `hl7:"POS=13" json:"ProtectionIndicator,omitempty"`
	ProtectionIndicatorEffectiveDate  time.Time `hl7:"POS=14" json:"ProtectionIndicatorEffectiveDate,omitempty"`
	PlaceOfWorship                    []XON     `hl7:"POS=15" json:"PlaceOfWorship,omitempty"`
	AdvanceDirectiveCode              []CE      `hl7:"POS=16" json:"AdvanceDirectiveCode,omitempty"`
	ImmunizationRegistryStatus        string    `hl7:"POS=17" json:"ImmunizationRegistryStatus,omitempty"`
	ImmunizationRegistryEffectiveDate time.Time `hl7:"POS=18" json:"ImmunizationRegistryEffectiveDate,omitempty"`
	PublicityCodeEffectiveDate        time.Time `hl7:"POS=19" json:"PublicityCodeEffectiveDate,omitempty"`
	MilitaryBranch                    string    `hl7:"POS=20" json:"MilitaryBranch,omitempty"`
	MilitaryRank                      string    `hl7:"POS=11" json:"MilitaryRank,omitempty"`
	MilitaryStatus                    string    `hl7:"POS=22" json:"MilitaryStatus,omitempty"`
}
