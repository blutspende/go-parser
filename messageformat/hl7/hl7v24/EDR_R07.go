package hl7v24

// EDR_R07 - Enhanced display response
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/EDR_R07
type EDR_R07 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryAcknowledgment QAK `hl7:"TAG=QAK"`
	DisplayData []DSP `hl7:"TAG=DSP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

