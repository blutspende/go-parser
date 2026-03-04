package hl7v251

// LRL - Location Relationship
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/LRL
type LRL struct {
	PrimaryKeyValueLrl PL `hl7:"POS=2"`
	SegmentActionCode string `hl7:"POS=3"`
	SegmentUniqueKey EI `hl7:"POS=4"`
	LocationRelationshipID CE `hl7:"POS=5"`
	OrganizationalLocationRelationshipValue []XON `hl7:"POS=6"`
	PatientLocationRelationshipValue PL `hl7:"POS=7"`
}

