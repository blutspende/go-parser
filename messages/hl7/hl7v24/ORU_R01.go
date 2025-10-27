package hl7v24

// HL7 v2.4 - ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MSH MSH `hl7:"TAG=MSH" json:"MSH,omitempty"`

	PatientResult []struct {
		Patient struct {
			PatientIdentification PID   `hl7:"TAG=PID;ATR=optional" json:"PatientIdentification,omitempty"`
			PatientDemographics   PD1   `hl7:"TAG=PD1;ATR=optional" json:"PatientDemographics,omitempty"`
			NextOfKin             NK1   `hl7:"TAG=PD1;ATR=optional" json:"NextOfKin,omitempty"`
			NotesAndComments      []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
			Visit                 struct {
				PatientVisit          PV1 `hl7:"TAG=PV1;ATR=optional" json:"PatientVisit,omitempty"`
				AdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional" json:"AdditionalInformation,omitempty"`
			} `hl7:"GROUP"`
		} `hl7:"GROUP"`

		OrderObservation []struct {
			CommonOrder        ORC   `hl7:"TAG=ORC;ATR=optional" json:"CommonOrder,omitempty"`
			ObservationRequest OBR   `hl7:"TAG=OBR;ATR=optional" json:"ObservationRequest,omitempty"`
			NotesAndComments   []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
			ContactData        []CTD `hl7:"TAG=CTD;ATR=optional" json:"ContactData,omitempty"`

			Observation []struct {
				Observation                 OBX   `hl7:"TAG=OBX;ATR=optional" json:"Observation,omitempty"`
				ObservationNotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional" json:"ObservationNotesAndComments,omitempty"`
			} `hl7:"GROUP"`
			FinancialTransaction        []FT1 `hl7:"TAG=FT1;ATR=optional" json:"FinancialTransaction,omitempty"`
			ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional" json:"ClinicalTrialIdentification,omitempty"`
		} `hl7:"GROUP"`
	} `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=MSH" json:"ContinuationPointer,omitempty"`
}
