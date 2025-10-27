package hl7v24

// HL7 v2.4 - OUL_R21 - Unsolicited laboratory observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OUL_R21
type OUL_R21 struct {
	MSH              MSH   `hl7:"TAG=MSH" json:"MSH,omitempty"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
	Patient          struct {
		PatientIdentification PID   `hl7:"TAG=PID;ATR=optional" json:"PatientIdentification,omitempty"`
		PatientDemographics   PD1   `hl7:"TAG=PD1;ATR=optional" json:"PatientDemographics,omitempty"`
		NotesAndComments      []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
	} `hl7:"GROUP"`
	Visit struct {
		PatientVisit          PV1 `hl7:"TAG=PV1;ATR=optional" json:"PatientVisit,omitempty"`
		AdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional" json:"AdditionalInformation,omitempty"`
	} `hl7:"GROUP"`
	OrderObservation []struct {
		Container struct {
			SpecimenAndContainerDetail SAC   `hl7:"TAG=SAC;ATR=optional" json:"SpecimenAndContainerDetail,omitempty"`
			SubstanceIdentifier        SID   `hl7:"TAG=SID;ATR=optional" json:"SubstanceIdentifier,omitempty"`
			ObservationResult          []OBX `hl7:"TAG=OBX;ATR=optional" json:"ObservationResult,omitempty"`
		} `hl7:"GROUP"`
		CommonOrder        ORC   `hl7:"TAG=ORC;ATR=optional" json:"CommonOrder,omitempty"`
		ObservationRequest OBR   `hl7:"TAG=OBR;ATR=optional" json:"ObservationRequest,omitempty"`
		NotesAndComments   []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
		Observation        []struct {
			Observation         OBX   `hl7:"TAG=OBX;ATR=optional" json:"Observation,omitempty"`
			TestCodeDetail      TCD   `hl7:"TAG=TCD;ATR=optional" json:"TestCodeDetail,omitempty"`
			SubstanceIdentifier SID   `hl7:"TAG=SID;ATR=optional" json:"SubstanceIdentifier,omitempty"`
			NotesAndComments    []NTE `hl7:"TAG=NTE;ATR=optional" json:"NotesAndComments,omitempty"`
		} `hl7:"GROUP"`
		ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional" json:"ClinicalTrialIdentification,omitempty"`
	} `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=MSH" json:"ContinuationPointer,omitempty"`
}
