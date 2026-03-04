package hl7v231

// CTI - Clinical trial identification segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/CTI
type CTI struct {
	SponsorStudyID EI `hl7:"POS=2"`
	StudyPhaseIdentifier CE `hl7:"POS=3"`
	StudyScheduledTimePoint CE `hl7:"POS=4"`
}

