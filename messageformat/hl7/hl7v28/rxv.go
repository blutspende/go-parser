package hl7v28

// RXV - Pharmacy/Treatment Infusion
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/RXV
type RXV struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BolusType string `hl7:"POS=3"`
	BolusDoseAmount *int `hl7:"POS=4"`
	BolusDoseAmountUnits CWE `hl7:"POS=5"`
	BolusDoseVolume *int `hl7:"POS=6"`
	BolusDoseVolumeUnits CWE `hl7:"POS=7"`
	PcaType string `hl7:"POS=8"`
	PcaDoseAmount *int `hl7:"POS=9"`
	PcaDoseAmountUnits CWE `hl7:"POS=10"`
	PcaDoseAmountVolume *int `hl7:"POS=11"`
	PcaDoseAmountVolumeUnits CWE `hl7:"POS=12"`
	MaxDoseAmount *int `hl7:"POS=13"`
	MaxDoseAmountUnits CWE `hl7:"POS=14"`
	MaxDoseAmountVolume *int `hl7:"POS=15"`
	MaxDoseAmountVolumeUnits CWE `hl7:"POS=16"`
	MaxDosePerTime CQ `hl7:"POS=17"`
	LockoutInterval CQ `hl7:"POS=18"`
	SyringeManufacturer CWE `hl7:"POS=19"`
	SyringeModelNumber CWE `hl7:"POS=20"`
	SyringeSize *int `hl7:"POS=21"`
	SyringeSizeUnits CWE `hl7:"POS=22"`
	ActionCode string `hl7:"POS=23"`
}

