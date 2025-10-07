package hl7v24

// HL7 v2.4 - ORL_O22 - General laboratory order response message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MSH                    MSH   `hl7:"MSH" json:"MSH,omitempty"`
	MessageAcknowledgement MSA   `hl7:"MSA" json:"MessageAcknowledgement,omitempty"`
	Error                  ERR   `hl7:"ERR,optional" json:"Error,omitempty"`
	NotesAndComments       []NTE `hl7:"NTE,optional" json:"NotesAndComments,omitempty"`
	Response               struct {
		Patient struct {
			PatientIdentification PID `hl7:"PID,optional" json:"PatientIdentification,omitempty"`
		}
		GeneralOrder []struct {
			Container struct {
				SpecimenAndContainerDetail SAC   `hl7:"SAC,optional" json:"SpecimenAndContainerDetail,omitempty"`
				ObservationResult          []OBX `hl7:"OBX,optional" json:"ObservationResult,omitempty"`
			}
			Order []struct {
				CommonOrder        ORC `hl7:"ORC,optional" json:"CommonOrder,omitempty"`
				ObservationRequest struct {
					TheRequest                 OBR   `hl7:"OBR,optional" json:"TheRequest,omitempty"`
					SpecimenAndContainerDetail []SAC `hl7:"SAC,optional" json:"SpecimenAndContainerDetail,omitempty"`
				}
			}
		}
	}
}
