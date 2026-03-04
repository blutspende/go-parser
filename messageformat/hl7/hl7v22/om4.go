package hl7v22

// OM4 - Observation That Require Specimens
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/OM4
type OM4 struct {
	SegmentTypeID string `hl7:"POS=2"`
	SequenceNumberTestObservationMasterFile *int `hl7:"POS=3"`
	DerivedSpecimen string `hl7:"POS=4"`
	ContainerDescription string `hl7:"POS=5"`
	ContainerVolume *int `hl7:"POS=6"`
	ContainerUnits CE `hl7:"POS=7"`
	Specimen CE `hl7:"POS=8"`
	Additive CE `hl7:"POS=9"`
	Preparation string `hl7:"POS=10"`
	SpecialHandlingRequirements string `hl7:"POS=11"`
	NormalCollectionVolume CQ `hl7:"POS=12"`
	MinimumCollectionVolume CQ `hl7:"POS=13"`
	SpecimenRequirements string `hl7:"POS=14"`
	SpecimenPriorities []string `hl7:"POS=15"`
	SpecimenRetentionTime CQ `hl7:"POS=16"`
}

