package hl7v23

// CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/CTI
type CTI struct {
	SponsorStudyID          EI `hl7:"1" json:"sponsorStudyID,omitempty"`
	StudyPhaseIdentifier    CE `hl7:"2" json:"studyPhaseIdentifier,omitempty"`
	StudyScheduledTimePoint CE `hl7:"3" json:"studyScheduledTimePoint,omitempty"`
}
