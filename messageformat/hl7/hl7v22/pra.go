package hl7v22

// PRA - Practitioner Detail
// https://hl7-definition.caristix.com/v2/HL7v2.2/Segments/PRA
type PRA struct {
	PraPrimaryKeyValue string `hl7:"POS=2"`
	PractitionerGroup []CE `hl7:"POS=3"`
	PractitionerCategory []string `hl7:"POS=4"`
	ProviderBilling string `hl7:"POS=5"`
	Specialty []CM_SPD `hl7:"POS=6"`
	PractitionerIDNumbers []CM_PLN `hl7:"POS=7"`
	Privileges []CM_PIP `hl7:"POS=8"`
}

