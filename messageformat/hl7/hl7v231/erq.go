package hl7v231

// ERQ - Event replay query segment
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/Segments/ERQ
type ERQ struct {
	QueryTag string `hl7:"POS=2"`
	EventIdentifier CE `hl7:"POS=3"`
	InputParameterList []QIP `hl7:"POS=4"`
}

