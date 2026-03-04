package hl7v251

// RCI_I05_PROVIDER - Group struct
type RCI_I05_PROVIDER struct {
	ProviderData PRD `hl7:"TAG=PRD"`
	ContactData []CTD `hl7:"TAG=CTD;ATR=optional"`
}

// RCI_I05_OBSERVATION - Group struct
type RCI_I05_OBSERVATION struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Results []RCI_I05_OBSERVATION_RESULTS `hl7:"GROUP;ATR=optional"`
}

// RCI_I05_OBSERVATION_RESULTS - Group struct
type RCI_I05_OBSERVATION_RESULTS struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// RCI_I05 - Request for patient clinical information acknowledgement
// https://hl7-definition.caristix.com/v2/HL7v2.5.1/TriggerEvents/RCI_I05
type RCI_I05 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	OriginalStyleQueryDefinition QRD `hl7:"TAG=QRD"`
	OriginalStyleQueryFilter QRF `hl7:"TAG=QRF;ATR=optional"`
	Provider []RCI_I05_PROVIDER `hl7:"GROUP"`
	PatientIdentification PID `hl7:"TAG=PID"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup []DRG `hl7:"TAG=DRG;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Observation []RCI_I05_OBSERVATION `hl7:"GROUP;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

