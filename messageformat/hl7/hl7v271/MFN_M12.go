package hl7v271

// MFN_M12_MF_OBS_ATTRIBUTES - Group struct
type MFN_M12_MF_OBS_ATTRIBUTES struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	AdditionalBasicAttributes OM7 `hl7:"TAG=OM7;ATR=optional"`
}

// MFN_M12 - Master file notification message
// https://hl7-definition.caristix.com/v2/HL7v2.7.1/TriggerEvents/MFN_M12
type MFN_M12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_obs_attributes []MFN_M12_MF_OBS_ATTRIBUTES `hl7:"GROUP"`
}

