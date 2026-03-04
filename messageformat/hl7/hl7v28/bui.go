package hl7v28

// BUI - Blood Unit information Segment
// https://hl7-definition.caristix.com/v2/HL7v2.8/Segments/BUI
type BUI struct {
	SetID int `hl7:"POS=2;ATR=sequence"`
	BloodUnitIdentifier EI `hl7:"POS=3"`
	BloodUnitType CWE `hl7:"POS=4"`
	BloodUnitWeight *int `hl7:"POS=5"`
	WeightUnits CNE `hl7:"POS=6"`
	BloodUnitVolume *int `hl7:"POS=7"`
	VolumeUnits CNE `hl7:"POS=8"`
	ContainerCatalogNumber string `hl7:"POS=9"`
	ContainerLotNumber string `hl7:"POS=10"`
	ContainerManufacturer XON `hl7:"POS=11"`
	TransportTemperature NR `hl7:"POS=12"`
	TransportTemperatureUnits []CNE `hl7:"POS=13"`
}

