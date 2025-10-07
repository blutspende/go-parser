package hl7v23

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PD1
type PD1 struct {
	LivingDependency    []string `hl7:"1" json:"livingDependency,omitempty"`
	LivingArrangement   string   `hl7:"2" json:"livingArrangement,omitempty"`
	PrimaryFacility     []XON    `hl7:"3" json:"primaryFacility,omitempty"`
	PrimaryCareProvider []XCN    `hl7:"4" json:"primaryCareProvider,omitempty"`
	StudentIndicator    string   `hl7:"5" json:"studentIndicator,omitempty"`
	Handicap            string   `hl7:"6" json:"handicap,omitempty"`
	LivingWill          string   `hl7:"7" json:"livingWill,omitempty"`
	OrganDonor          string   `hl7:"8" json:"organDonor,omitempty"`
	SeparateBill        string   `hl7:"9" json:"separateBill,omitempty"`
	DuplicatePatient    []CX     `hl7:"10" json:"duplicatePatient,omitempty"`
	PublicityIndicator  CE       `hl7:"11" json:"publicityIndicator,omitempty"`
	ProtectionIndicator string   `hl7:"12" json:"protectionIndicator,omitempty"`
}
