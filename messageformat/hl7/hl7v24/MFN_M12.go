package hl7v24

// MFN_M12_MF_OBS_ATTRIBUTES - Group struct
type MFN_M12_MF_OBS_ATTRIBUTES struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	AdditionalBasicAttributes OM7 `hl7:"TAG=OM7;ATR=optional"`
}

// MFN_M12 - Master files notification - Service/text/observation master file
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/MFN_M12
type MFN_M12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	MfObsAttributes []MFN_M12_MF_OBS_ATTRIBUTES `hl7:"GROUP"`
}

