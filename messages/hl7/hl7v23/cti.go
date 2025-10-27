package hl7v23

// CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/CTI
type CTI struct {
	SponsorStudyID          EI `hl7:"POS=2" json:"sponsorStudyID,omitempty"`
	StudyPhaseIdentifier    CE `hl7:"POS=3" json:"studyPhaseIdentifier,omitempty"`
	StudyScheduledTimePoint CE `hl7:"POS=4" json:"studyScheduledTimePoint,omitempty"`
}
