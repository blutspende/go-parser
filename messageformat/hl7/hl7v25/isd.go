package hl7v25

// ISD - Interaction Status Detail
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/ISD
type ISD struct {
	ReferenceInteractionNumberUniqueIdentifier *int `hl7:"POS=2"`
	InteractionTypeIdentifier CE `hl7:"POS=3"`
	InteractionActiveState CE `hl7:"POS=4"`
}

