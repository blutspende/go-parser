package hl7v28

// CDM - Charge Description Master
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/CDM
type CDM struct {
	PrimaryKeyValueCdm CWE `hl7:"POS=2"`
	ChargeCodeAlias []CWE `hl7:"POS=3"`
	ChargeDescriptionShort string `hl7:"POS=4"`
	ChargeDescriptionLong string `hl7:"POS=5"`
	DescriptionOverrideIndicator CWE `hl7:"POS=6"`
	ExplodingCharges []CWE `hl7:"POS=7"`
	ProcedureCode []CNE `hl7:"POS=8"`
	ActiveInactiveFlag string `hl7:"POS=9"`
	InventoryNumber []CWE `hl7:"POS=10"`
	ResourceLoad *int `hl7:"POS=11"`
	ContractNumber []CX `hl7:"POS=12"`
	ContractOrganization []XON `hl7:"POS=13"`
	RoomFeeIndicator string `hl7:"POS=14"`
}

