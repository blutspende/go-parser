package hl7v24

// HL7 v2.4 - OML_O21 - Laboratory order
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OML_O21
type OML_O21 struct {
	MSH              MSH   `hl7:"MSH" json:"MSH,omitempty"`
	NotesAndComments []NTE `hl7:"NTE,optional" json:"NotesAndComments,omitempty"`
	Patient          struct {
		PatientIdentification         PID   `hl7:"PID" json:"PatientIdentification,omitempty"`
		PatientAdditionalDemographics PD1   `hl7:"PD1,optional" json:"PatientAdditionalDemographics,omitempty"`
		NotesAndComments              []NTE `hl7:"NTE,optional" json:"NotesAndComments,omitempty"`
		Visit                         struct {
			PatientVisit          PV1 `hl7:"PV1" json:"PatientVisit,omitempty"`
			AdditionalInformation PV2 `hl7:"PV2,optional" json:"AdditionalInformation,omitempty"`
		}
		Insurance []struct {
			Insurance                          IN1 `hl7:"IN1" json:"Insurance,omitempty"`
			AdditionalInformation              IN2 `hl7:"IN2, optional" json:"AdditionalInformation,omitempty"`
			AdditionalInformationCertification IN3 `hl7:"IN3, optional" json:"AdditionalInformationCertification,omitempty"`
		}
		Guarantor                 GT1   `hl7:"GT1, optional" json:"Guarantor,omitempty"`
		PatientAllergyInformation []AL1 `hl7:"AL1, optional" json:"PatientAllergyInformation,omitempty"`
	}
	OrderGeneral []struct {
		Container struct {
			SpecimenAndContainerDetail SAC   `hl7:"SAC" json:"SpecimenAndContainerDetail,omitempty"`
			ObservationResult          []OBX `hl7:"OBX, optional" json:"ObservationResult,omitempty"`
		}
		Order []struct {
			CommonOrder        ORC `hl7:"ORC" json:"CommonOrder,omitempty"`
			ObservationRequest struct {
				ObservationRequest OBR `hl7:"OBR" json:"ObservationRequest,omitempty"`
				Container          []struct {
					TestCodeDetail   TCD   `hl7:"TCD, optional" json:"TestCodeDetail,omitempty"`
					NotesAndComments []NTE `hl7:"NTE, optional" json:"NotesAndComments,omitempty"`
					Diagnosis        []DG1 `hl7:"DG1, optional" json:"Diagnosis,omitempty"`
				}
				Observation []struct {
					ObservationResult OBX   `hl7:"OBX" json:"ObservationResult,omitempty"`
					TestCodeDetail    TCD   `hl7:"TCD, optional" json:"TestCodeDetail,omitempty"`
					NotesAndComments  []NTE `hl7:"NTE, optional" json:"NotesAndComments,omitempty"`
				}
				PriorResult []struct {
					PatientPrior struct {
						PatientIdentification         PID `hl7:"PID" json:"PatientIdentification,omitempty"`
						PatientAdditionalDemographics PD1 `hl7:"PD1,optional" json:"PatientAdditionalDemographics,omitempty"`
					}
					PatientVisitPrior struct {
						PatientVisit          PV1 `hl7:"PV1" json:"PatientVisit,omitempty"`
						AdditionalInformation PV2 `hl7:"PV2,optional" json:"AdditionalInformation,omitempty"`
					}
					PatientAllergyInformation []AL1 `hl7:"AL1, optional" json:"PatientAllergyInformation,omitempty"`
					OrderPrior                []struct {
						CommonOrder        ORC   `hl7:"ORC, optional" json:"CommonOrder,omitempty"`
						ObservationRequest OBR   `hl7:"OBR" json:"ObservationRequest,omitempty"`
						NotesAndComments   []NTE `hl7:"NTE, optional" json:"NotesAndComments,omitempty"`
						ObservationPrior   []struct {
							ObservationResult OBX   `hl7:"OBX" json:"ObservationResult,omitempty"`
							NotesAndComments  []NTE `hl7:"NTE, optional" json:"NotesAndComments,omitempty"`
						}
					}
				}
			}
			FinancialTransaction        []FT1 `hl7:"FT1, optional" json:"FinancialTransaction,omitempty"`
			ClinicalTrialIdentification []CTI `hl7:"CTI, optional" json:"ClinicalTrialIdentification,omitempty"`
			Billing                     BLG   `hl7:"BLG, optional" json:"Billing,omitempty"`
		}
	}
}
