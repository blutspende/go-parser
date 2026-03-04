package hl7v271

// CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/Segments/CTI
type CTI struct {
	SponsorStudyID EI `hl7:"POS=2"`
	StudyPhaseIdentifier CWE `hl7:"POS=3"`
	StudyScheduledTimePoint CWE `hl7:"POS=4"`
}

