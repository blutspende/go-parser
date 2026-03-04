package hl7v27

// UDM_Q05 - Unsolicited display update message
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/UDM_Q05
type UDM_Q05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

