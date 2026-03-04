package hl7v251

// QRY_Q30 - Pharmacy/Treatment Dose Information Query
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/QRY_Q30
type QRY_Q30 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

