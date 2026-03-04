package hl7v24

// CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTI
type CTI struct {
	SponsorStudyID EI `hl7:"POS=2"`
	StudyPhaseIdentifier CE `hl7:"POS=3"`
	StudyScheduledTimePoint CE `hl7:"POS=4"`
}

