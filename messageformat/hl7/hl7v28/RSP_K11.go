package hl7v28

// RSP_K11_SEGMENT_PATTERN - Group struct
type RSP_K11_SEGMENT_PATTERN struct {
	AnyHl7Segment Hxx `hl7:"TAG=Hxx"`
}

// RSP_K11 - Segment pattern response in response to QBP^Q11
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/RSP_K11
type RSP_K11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	QueryParameterDefinition QPD `hl7:"TAG=QPD"`
	Segment_pattern RSP_K11_SEGMENT_PATTERN `hl7:"GROUP;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

