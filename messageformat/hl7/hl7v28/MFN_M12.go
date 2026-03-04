package hl7v28

// MFN_M12_MF_OBS_ATTRIBUTES - Group struct
type MFN_M12_MF_OBS_ATTRIBUTES struct {
	MasterFileEntry MFE `hl7:"TAG=MFE"`
	GeneralSegment OM1 `hl7:"TAG=OM1"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Mf_obs_other_attributes MFN_M12_MF_OBS_ATTRIBUTES_MF_OBS_OTHER_ATTRIBUTES `hl7:"GROUP;ATR=optional"`
}

// MFN_M12_MF_OBS_ATTRIBUTES_MF_OBS_OTHER_ATTRIBUTES - Group struct
type MFN_M12_MF_OBS_ATTRIBUTES_MF_OBS_OTHER_ATTRIBUTES struct {
	AdditionalBasicAttributes OM7 `hl7:"TAG=OM7"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// MFN_M12 - Master file notification message
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/MFN_M12
type MFN_M12 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	MasterFileIdentification MFI `hl7:"TAG=MFI"`
	Mf_obs_attributes []MFN_M12_MF_OBS_ATTRIBUTES `hl7:"GROUP"`
}

