package hl7v28

// REL - Clinical Relationship Segment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/REL
type REL struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	RelationshipType CWE `hl7:"POS=3"`
	ThisRelationshipInstanceIdentifier EI `hl7:"POS=4"`
	SourceInformationInstanceIdentifier EI `hl7:"POS=5"`
	TargetInformationInstanceIdentifier EI `hl7:"POS=6"`
	AssertingEntityInstanceID EI `hl7:"POS=7"`
	AssertingPerson XCN `hl7:"POS=8"`
	AssertingOrganization XON `hl7:"POS=9"`
	AssertorAddress XAD `hl7:"POS=10"`
	AssertorContact XTN `hl7:"POS=11"`
	AssertionDateRange DR `hl7:"POS=12"`
	NegationIndicator string `hl7:"POS=13"`
	CertaintyOfRelationship CWE `hl7:"POS=14"`
	PriorityNo *int `hl7:"POS=15"`
	PrioritySequenceNoRelPreferenceForConsideration *int `hl7:"POS=16"`
	SeparabilityIndicator string `hl7:"POS=17"`
}

