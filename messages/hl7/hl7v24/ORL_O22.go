package hl7v24

// HL7 v2.4 - ORL_O22 - General laboratory order response message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MSH                    MSH   `hl7:"TAG=MSH" json:"MSH,omitempty"`
	MessageAcknowledgement MSA   `hl7:"TAG=MSA" json:"MessageAcknowledgement,omitempty"`
	Error                  ERR   `hl7:"TAG=ERR;ATR=optional" json:"Error,omitempty"`
	NotesAndComments       []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
	Response               struct {
		Patient struct {
			PatientIdentification PID `hl7:"TAG=PID;ATR=optional" json:"PatientIdentification,omitempty"`
		} `hl7:"GROUP"`
		GeneralOrder []struct {
			Container struct {
				SpecimenAndContainerDetail SAC   `hl7:"TAG=SAC;ATR=optional" json:"SpecimenAndContainerDetail,omitempty"`
				ObservationResult          []OBX `hl7:"TAG=OBX;ATR=optional" json:"ObservationResult,omitempty"`
			} `hl7:"GROUP"`
			Order []struct {
				CommonOrder        ORC `hl7:"TAG=ORC;ATR=optional" json:"CommonOrder,omitempty"`
				ObservationRequest struct {
					TheRequest                 OBR   `hl7:"TAG=OBR;ATR=optional" json:"TheRequest,omitempty"`
					SpecimenAndContainerDetail []SAC `hl7:"TAG=SAC;ATR=optional" json:"SpecimenAndContainerDetail,omitempty"`
				} `hl7:"GROUP"`
			} `hl7:"GROUP"`
		} `hl7:"GROUP"`
	} `hl7:"GROUP"`
}
