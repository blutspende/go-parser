package hl7v251

// CDM - Charge Description Master
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/Segments/CDM
type CDM struct {
	PrimaryKeyValueCdm CE `hl7:"POS=2"`
	ChargeCodeAlias []CE `hl7:"POS=3"`
	ChargeDescriptionShort string `hl7:"POS=4"`
	ChargeDescriptionLong string `hl7:"POS=5"`
	DescriptionOverrideIndicator string `hl7:"POS=6"`
	ExplodingCharges []CE `hl7:"POS=7"`
	ProcedureCode []CE `hl7:"POS=8"`
	ActiveInactiveFlag string `hl7:"POS=9"`
	InventoryNumber []CE `hl7:"POS=10"`
	ResourceLoad *int `hl7:"POS=11"`
	ContractNumber []CX `hl7:"POS=12"`
	ContractOrganization []XON `hl7:"POS=13"`
	RoomFeeIndicator string `hl7:"POS=14"`
}

