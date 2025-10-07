package hl7v24

// HL7 v2.4 - CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTI
type CTI struct {
	SponsorStudyID          EI `hl7:"1" json:"SponsorStudyID,omitempty"`
	StudyPhaseIdentifier    CE `hl7:"2" json:"StudyPhaseIdentifier,omitempty"`
	StudyScheduledTimePoint CE `hl7:"3" json:"StudyScheduledTimePoint,omitempty"`
}
