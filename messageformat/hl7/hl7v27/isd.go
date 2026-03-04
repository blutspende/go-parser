package hl7v27

// ISD - Interaction Status Detail
// https://hl7-definition.caristix.com/v2/HL7v2.7/Segments/ISD
type ISD struct {
	ReferenceInteractionNumber *int `hl7:"POS=2"`
	InteractionTypeIdentifier CWE `hl7:"POS=3"`
	InteractionActiveState CWE `hl7:"POS=4"`
}

