package hl7v24

// HL7 v2.4 - OML_O21 - Laboratory order
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OML_O21
type OML_O21 struct {
	MSH              MSH   `hl7:"TAG=MSH" json:"MSH,omitempty"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
	Patient          struct {
		PatientIdentification         PID   `hl7:"TAG=PID" json:"PatientIdentification,omitempty"`
		PatientAdditionalDemographics PD1   `hl7:"TAG=PD1;ATR=optional" json:"PatientAdditionalDemographics,omitempty"`
		NotesAndComments              []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
		Visit                         struct {
			PatientVisit          PV1 `hl7:"TAG=PV1" json:"PatientVisit,omitempty"`
			AdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional" json:"AdditionalInformation,omitempty"`
		} `hl7:"GROUP"`
		Insurance []struct {
			Insurance                          IN1 `hl7:"TAG=IN1" json:"Insurance,omitempty"`
			AdditionalInformation              IN2 `hl7:"TAG=IN2;ATR= optional" json:"AdditionalInformation,omitempty"`
			AdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR= optional" json:"AdditionalInformationCertification,omitempty"`
		} `hl7:"GROUP"`
		Guarantor                 GT1   `hl7:"TAG=GT1;ATR= optional" json:"Guarantor,omitempty"`
		PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR= optional" json:"PatientAllergyInformation,omitempty"`
	} `hl7:"GROUP"`
	OrderGeneral []struct {
		Container struct {
			SpecimenAndContainerDetail SAC   `hl7:"TAG=SAC" json:"SpecimenAndContainerDetail,omitempty"`
			ObservationResult          []OBX `hl7:"TAG=OBX;ATR= optional" json:"ObservationResult,omitempty"`
		} `hl7:"GROUP"`
		Order []struct {
			CommonOrder        ORC `hl7:"TAG=ORC" json:"CommonOrder,omitempty"`
			ObservationRequest struct {
				ObservationRequest OBR `hl7:"TAG=OBR" json:"ObservationRequest,omitempty"`
				Container          []struct {
					TestCodeDetail   TCD   `hl7:"TAG=TCD;ATR= optional" json:"TestCodeDetail,omitempty"`
					NotesAndComments []NTE `hl7:"TAG=NTE;ATR= optional" json:"NotesAndComments,omitempty"`
					Diagnosis        []DG1 `hl7:"TAG=DG1;ATR= optional" json:"Diagnosis,omitempty"`
				} `hl7:"GROUP"`
				Observation []struct {
					ObservationResult OBX   `hl7:"TAG=OBX" json:"ObservationResult,omitempty"`
					TestCodeDetail    TCD   `hl7:"TAG=TCD;ATR= optional" json:"TestCodeDetail,omitempty"`
					NotesAndComments  []NTE `hl7:"TAG=NTE;ATR= optional" json:"NotesAndComments,omitempty"`
				} `hl7:"GROUP"`
				PriorResult []struct {
					PatientPrior struct {
						PatientIdentification         PID `hl7:"TAG=PID" json:"PatientIdentification,omitempty"`
						PatientAdditionalDemographics PD1 `hl7:"TAG=PD1;ATR=optional" json:"PatientAdditionalDemographics,omitempty"`
					} `hl7:"GROUP"`
					PatientVisitPrior struct {
						PatientVisit          PV1 `hl7:"TAG=PV1" json:"PatientVisit,omitempty"`
						AdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional" json:"AdditionalInformation,omitempty"`
					} `hl7:"GROUP"`
					PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR= optional" json:"PatientAllergyInformation,omitempty"`
					OrderPrior                []struct {
						CommonOrder        ORC   `hl7:"TAG=ORC;ATR= optional" json:"CommonOrder,omitempty"`
						ObservationRequest OBR   `hl7:"TAG=OBR" json:"ObservationRequest,omitempty"`
						NotesAndComments   []NTE `hl7:"TAG=NTE;ATR= optional" json:"NotesAndComments,omitempty"`
						ObservationPrior   []struct {
							ObservationResult OBX   `hl7:"TAG=OBX" json:"ObservationResult,omitempty"`
							NotesAndComments  []NTE `hl7:"TAG=NTE;ATR= optional" json:"NotesAndComments,omitempty"`
						} `hl7:"GROUP"`
					} `hl7:"GROUP"`
				} `hl7:"GROUP"`
			} `hl7:"GROUP"`
			FinancialTransaction        []FT1 `hl7:"TAG=FT1;ATR= optional" json:"FinancialTransaction,omitempty"`
			ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR= optional" json:"ClinicalTrialIdentification,omitempty"`
			Billing                     BLG   `hl7:"TAG=BLG;ATR= optional" json:"Billing,omitempty"`
		} `hl7:"GROUP"`
	} `hl7:"GROUP"`
}
