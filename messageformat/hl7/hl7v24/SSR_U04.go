package hl7v24

// HL7 v2.4 - SSR_U04 - Specimen status request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SSR_U04
type SSR_U04 struct {
	MSH                     MSH `hl7:"TAG=MSH" json:"MSH,omitempty"`
	EquipmentDetail         EQU `hl7:"TAG=EQU" json:"EquipmentDetail,omitempty"`
	SpecimenContainerDetail SAC `hl7:"TAG=SAC" json:"SpecimenContainerDetail,omitempty"`
	Role                    ROL `hl7:"TAG=ROL;ATR=optional" json:"Role,omitempty"`
}
