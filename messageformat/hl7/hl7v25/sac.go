package hl7v25

import "time"

// SAC - Specimen Container detail
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/SAC
type SAC struct {
	ExternalAccessionIdentifier EI `hl7:"POS=2"`
	AccessionIdentifier EI `hl7:"POS=3"`
	ContainerIdentifier EI `hl7:"POS=4"`
	PrimaryParentContainerIdentifier EI `hl7:"POS=5"`
	EquipmentContainerIdentifier EI `hl7:"POS=6"`
	SpecimenSource SPS `hl7:"POS=7"`
	RegistrationDateTime time.Time `hl7:"POS=8"`
	ContainerStatus CE `hl7:"POS=9"`
	CarrierType CE `hl7:"POS=10"`
	CarrierIdentifier EI `hl7:"POS=11"`
	PositionInCarrier NA `hl7:"POS=12"`
	TrayTypeSac CE `hl7:"POS=13"`
	TrayIdentifier EI `hl7:"POS=14"`
	PositionInTray NA `hl7:"POS=15"`
	Location []CE `hl7:"POS=16"`
	ContainerHeight *int `hl7:"POS=17"`
	ContainerDiameter *int `hl7:"POS=18"`
	BarrierDelta *int `hl7:"POS=19"`
	BottomDelta *int `hl7:"POS=20"`
	ContainerHeightDiameterDeltaUnits CE `hl7:"POS=21"`
	ContainerVolume *int `hl7:"POS=22"`
	AvailableSpecimenVolume *int `hl7:"POS=23"`
	InitialSpecimenVolume *int `hl7:"POS=24"`
	VolumeUnits CE `hl7:"POS=25"`
	SeparatorType CE `hl7:"POS=26"`
	CapType CE `hl7:"POS=27"`
	Additive []CWE `hl7:"POS=28"`
	SpecimenComponent CE `hl7:"POS=29"`
	DilutionFactor SN `hl7:"POS=30"`
	Treatment CE `hl7:"POS=31"`
	Temperature SN `hl7:"POS=32"`
	HemolysisIndex *int `hl7:"POS=33"`
	HemolysisIndexUnits CE `hl7:"POS=34"`
	LipemiaIndex *int `hl7:"POS=35"`
	LipemiaIndexUnits CE `hl7:"POS=36"`
	IcterusIndex *int `hl7:"POS=37"`
	IcterusIndexUnits CE `hl7:"POS=38"`
	FibrinIndex *int `hl7:"POS=39"`
	FibrinIndexUnits CE `hl7:"POS=40"`
	SystemInducedContaminants []CE `hl7:"POS=41"`
	DrugInterference []CE `hl7:"POS=42"`
	ArtificialBlood CE `hl7:"POS=43"`
	SpecialHandlingCode []CWE `hl7:"POS=44"`
	OtherEnvironmentalFactors []CE `hl7:"POS=45"`
}

