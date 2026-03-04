package hl7v24

// ISD - Interaction Status Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ISD
type ISD struct {
	ReferenceInteractionNumber *int `hl7:"POS=2"`
	InteractionTypeIdentifier CE `hl7:"POS=3"`
	InteractionActiveState CE `hl7:"POS=4"`
}

