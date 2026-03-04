package hl7v25

// OM4 - Observations that Require Specimens
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/OM4
type OM4 struct {
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=2"`
	DerivedSpecimen string `hl7:"POS=3"`
	ContainerDescription string `hl7:"POS=4"`
	ContainerVolume *int `hl7:"POS=5"`
	ContainerUnits CE `hl7:"POS=6"`
	Specimen CE `hl7:"POS=7"`
	Additive CWE `hl7:"POS=8"`
	Preparation string `hl7:"POS=9"`
	SpecialHandlingRequirements string `hl7:"POS=10"`
	NormalCollectionVolume CQ `hl7:"POS=11"`
	MinimumCollectionVolume CQ `hl7:"POS=12"`
	SpecimenRequirements string `hl7:"POS=13"`
	SpecimenPriorities []string `hl7:"POS=14"`
	SpecimenRetentionTime CQ `hl7:"POS=15"`
}

