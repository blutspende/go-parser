package hl7v24

import "time"

// HL7 v2.4 - SAC - Specimen and container detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SAC
type SAC struct {
	ExternalAccessionIdentifier       EI        `hl7:"POS=2" json:"ExternalAccessionIdentifier,omitempty"`
	AccessionIdentifier               EI        `hl7:"POS=3" json:"AccessionIdentifier,omitempty"`
	ContainerIdentifier               EI        `hl7:"POS=4" json:"ContainerIdentifier,omitempty"`
	Primary                           EI        `hl7:"POS=5" json:"Primary,omitempty"`
	EquipmentContainerIdentifier      EI        `hl7:"POS=6" json:"EquipmentContainerIdentifier,omitempty"`
	SpecimenSource                    SPS       `hl7:"POS=7" json:"SpecimenSource,omitempty"`
	RegistrationDateTime              time.Time `hl7:"POS=8" json:"RegistrationDateTime,omitempty"`
	ContainerStatus                   CE        `hl7:"POS=9" json:"ContainerStatus,omitempty"`
	CarrierType                       CE        `hl7:"POS=10" json:"CarrierType,omitempty"`
	CarrierIdentifier                 EI        `hl7:"POS=11" json:"CarrierIdentifier,omitempty"`
	PositionInCarrier                 NA        `hl7:"POS=12" json:"PositionInCarrier,omitempty"`
	TrayTypeSAC                       CE        `hl7:"POS=13" json:"TrayTypeSAC,omitempty"`
	TrayIdentifier                    EI        `hl7:"POS=14" json:"TrayIdentifier,omitempty"`
	PositionInTray                    NA        `hl7:"POS=15" json:"PositionInTray,omitempty"`
	Location                          []CE      `hl7:"POS=16" json:"Location,omitempty"`
	ContainerHeight                   *float32  `hl7:"POS=17" json:"ContainerHeight,omitempty"`
	ContainerDiameter                 *float32  `hl7:"POS=18" json:"ContainerDiameter,omitempty"`
	BarrierDelta                      *float32  `hl7:"POS=19" json:"BarrierDelta,omitempty"`
	BottomDelta                       *float32  `hl7:"POS=20" json:"BottomDelta,omitempty"`
	ContainerHeightDiameterDeltaUnits CE        `hl7:"POS=21" json:"ContainerHeightDiameterDeltaUnits,omitempty"`
	ContainerVolume                   *float32  `hl7:"POS=22" json:"ContainerVolume,omitempty"`
	AvailableVolume                   *float32  `hl7:"POS=23" json:"AvailableVolume,omitempty"`
	InitialSpecimenVolume             *float32  `hl7:"POS=24" json:"InitialSpecimenVolume,omitempty"`
	VolumeUnits                       CE        `hl7:"POS=25" json:"VolumeUnits,omitempty"`
	SeparatorType                     CE        `hl7:"POS=26" json:"SeparatorType,omitempty"`
	CapType                           CE        `hl7:"POS=27" json:"CapType,omitempty"`
	Additive                          []CE      `hl7:"POS=28" json:"Additive,omitempty"`
	SpecimenComponent                 CE        `hl7:"POS=29" json:"SpecimenComponent,omitempty"`
	DilutionFactor                    SN        `hl7:"POS=30" json:"DilutionFactor,omitempty"`
	Treatment                         CE        `hl7:"POS=31" json:"Treatment,omitempty"`
	Temperature                       SN        `hl7:"POS=32" json:"Temperature,omitempty"`
	HemolysisIndex                    *float32  `hl7:"POS=33" json:"HemolysisIndex,omitempty"`
	HomolysisIndexUnits               CE        `hl7:"POS=34" json:"HomolysisIndexUnits,omitempty"`
	LipemiaIndex                      *float32  `hl7:"POS=35" json:"LipemiaIndex,omitempty"`
	LipemiaIndexUnits                 CE        `hl7:"POS=36" json:"LipemiaIndexUnits,omitempty"`
	IcterusIndex                      *float32  `hl7:"POS=27" json:"IcterusIndex,omitempty"`
	IcterusIndexUnits                 CE        `hl7:"POS=38" json:"IcterusIndexUnits,omitempty"`
	FibrinIndex                       *float32  `hl7:"POS=39" json:"FibrinIndex,omitempty"`
	FibrinIndexUnits                  CE        `hl7:"POS=40" json:"FibrinIndexUnits,omitempty"`
	SystemInducedContaminants         []CE      `hl7:"POS=41" json:"SystemInducedContaminants,omitempty"`
	DrugInterference                  []CE      `hl7:"POS=42" json:"DrugInterference,omitempty"`
	ArtificialBlood                   CE        `hl7:"POS=43" json:"ArtificialBlood,omitempty"`
	SpecialHandlingConsiderations     []CE      `hl7:"POS=44" json:"SpecialHandlingConsiderations,omitempty"`
	OtherEnvironmentalFactors         []CE      `hl7:"POS=45" json:"OtherEnvironmentalFactors,omitempty"`
}
