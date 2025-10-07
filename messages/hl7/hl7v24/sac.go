package hl7v24

import "time"

// HL7 v2.4 - SAC - Specimen and container detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SAC
type SAC struct {
	ExternalAccessionIdentifier       EI        `hl7:"1" json:"ExternalAccessionIdentifier,omitempty"`
	AccessionIdentifier               EI        `hl7:"2" json:"AccessionIdentifier,omitempty"`
	ContainerIdentifier               EI        `hl7:"3" json:"ContainerIdentifier,omitempty"`
	Primary                           EI        `hl7:"4" json:"Primary,omitempty"`
	EquipmentContainerIdentifier      EI        `hl7:"5" json:"EquipmentContainerIdentifier,omitempty"`
	SpecimenSource                    SPS       `hl7:"6" json:"SpecimenSource,omitempty"`
	RegistrationDateTime              time.Time `hl7:"7,longdate" json:"RegistrationDateTime,omitempty"`
	ContainerStatus                   CE        `hl7:"8" json:"ContainerStatus,omitempty"`
	CarrierType                       CE        `hl7:"9" json:"CarrierType,omitempty"`
	CarrierIdentifier                 EI        `hl7:"10" json:"CarrierIdentifier,omitempty"`
	PositionInCarrier                 NA        `hl7:"11" json:"PositionInCarrier,omitempty"`
	TrayTypeSAC                       CE        `hl7:"12" json:"TrayTypeSAC,omitempty"`
	TrayIdentifier                    EI        `hl7:"13" json:"TrayIdentifier,omitempty"`
	PositionInTray                    NA        `hl7:"14" json:"PositionInTray,omitempty"`
	Location                          []CE      `hl7:"15" json:"Location,omitempty"`
	ContainerHeight                   float32   `hl7:"16" json:"ContainerHeight,omitempty"`
	ContainerDiameter                 float32   `hl7:"17" json:"ContainerDiameter,omitempty"`
	BarrierDelta                      float32   `hl7:"18" json:"BarrierDelta,omitempty"`
	BottomDelta                       float32   `hl7:"19" json:"BottomDelta,omitempty"`
	ContainerHeightDiameterDeltaUnits CE        `hl7:"20" json:"ContainerHeightDiameterDeltaUnits,omitempty"`
	ContainerVolume                   float32   `hl7:"21" json:"ContainerVolume,omitempty"`
	AvailableVolume                   float32   `hl7:"22" json:"AvailableVolume,omitempty"`
	InitialSpecimenVolume             float32   `hl7:"23" json:"InitialSpecimenVolume,omitempty"`
	VolumeUnits                       CE        `hl7:"24" json:"VolumeUnits,omitempty"`
	SeparatorType                     CE        `hl7:"25" json:"SeparatorType,omitempty"`
	CapType                           CE        `hl7:"26" json:"CapType,omitempty"`
	Additive                          []CE      `hl7:"27" json:"Additive,omitempty"`
	SpecimenComponent                 CE        `hl7:"28" json:"SpecimenComponent,omitempty"`
	DilutionFactor                    SN        `hl7:"29" json:"DilutionFactor,omitempty"`
	Treatment                         CE        `hl7:"30" json:"Treatment,omitempty"`
	Temperature                       SN        `hl7:"31" json:"Temperature,omitempty"`
	HemolysisIndex                    float32   `hl7:"32" json:"HemolysisIndex,omitempty"`
	HomolysisIndexUnits               CE        `hl7:"33" json:"HomolysisIndexUnits,omitempty"`
	LipemiaIndex                      float32   `hl7:"34" json:"LipemiaIndex,omitempty"`
	LipemiaIndexUnits                 CE        `hl7:"35" json:"LipemiaIndexUnits,omitempty"`
	IcterusIndex                      float32   `hl7:"26" json:"IcterusIndex,omitempty"`
	IcterusIndexUnits                 CE        `hl7:"37" json:"IcterusIndexUnits,omitempty"`
	FibrinIndex                       float32   `hl7:"38" json:"FibrinIndex,omitempty"`
	FibrinIndexUnits                  CE        `hl7:"39" json:"FibrinIndexUnits,omitempty"`
	SystemInducedContaminants         []CE      `hl7:"40" json:"SystemInducedContaminants,omitempty"`
	DrugInterference                  []CE      `hl7:"41" json:"DrugInterference,omitempty"`
	ArtificialBlood                   CE        `hl7:"42" json:"ArtificialBlood,omitempty"`
	SpecialHandlingConsiderations     []CE      `hl7:"43" json:"SpecialHandlingConsiderations,omitempty"`
	OtherEnvironmentalFactors         []CE      `hl7:"44" json:"OtherEnvironmentalFactors,omitempty"`
}
