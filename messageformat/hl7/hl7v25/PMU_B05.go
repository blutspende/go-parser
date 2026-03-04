package hl7v25

// PMU_B05 - Deactivate Practicing Person
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/PMU_B05
type PMU_B05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	StaffIdentification STF `hl7:"TAG=STF"`
	PractitionerDetail []PRA `hl7:"TAG=PRA;ATR=optional"`
	PractitionerOrganizationUnit []ORG `hl7:"TAG=ORG;ATR=optional"`
}

