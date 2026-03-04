package hl7v271

// TQ2 - Timing/quantity Relationship
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/TQ2
type TQ2 struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	SequenceResultsFlag string `hl7:"POS=3"`
	RelatedPlacerNumber []EI `hl7:"POS=4"`
	RelatedFillerNumber []EI `hl7:"POS=5"`
	RelatedPlacerGroupNumber []EI `hl7:"POS=6"`
	SequenceConditionCode string `hl7:"POS=7"`
	CyclicEntryExitIndicator string `hl7:"POS=8"`
	SequenceConditionTimeInterval CQ `hl7:"POS=9"`
	CyclicGroupMaximumNumberOfRepeats *int `hl7:"POS=10"`
	SpecialServiceRequestRelationship string `hl7:"POS=11"`
}

