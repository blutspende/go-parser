package hl7v271

// RDY_Z80_SEGMENT_PATTERN - Group struct
type RDY_Z80_SEGMENT_PATTERN struct {
	AnyHl7Segment Hxx `hl7:"TAG=Hxx"`
}

// RDY_Z80 - Dispense Information (Response)
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/RDY_Z80
type RDY_Z80 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Segment_pattern RDY_Z80_SEGMENT_PATTERN `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

