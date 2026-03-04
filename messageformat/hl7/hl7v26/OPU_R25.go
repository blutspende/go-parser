package hl7v26

// OPU_R25_ACCESSION_DETAIL - Group struct
type OPU_R25_ACCESSION_DETAIL struct {
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1"`
	Patient OPU_R25_ACCESSION_DETAIL_PATIENT `hl7:"GROUP;ATR=optional"`
	Specimen []OPU_R25_ACCESSION_DETAIL_SPECIMEN `hl7:"GROUP"`
}

// OPU_R25_ACCESSION_DETAIL_PATIENT - Group struct
type OPU_R25_ACCESSION_DETAIL_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientObservation []OPU_R25_ACCESSION_DETAIL_PATIENT_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OPU_R25_ACCESSION_DETAIL_PATIENT_PATIENT_OBSERVATION - Group struct
type OPU_R25_ACCESSION_DETAIL_PATIENT_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	SpecimenObservation []OPU_R25_ACCESSION_DETAIL_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Container []OPU_R25_ACCESSION_DETAIL_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
	Order []OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER `hl7:"GROUP"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN_CONTAINER - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	InventoryDetail INV `hl7:"TAG=INV;ATR=optional"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
	TimingQty []OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_TIMING_QTY `hl7:"GROUP;ATR=optional"`
	Result []OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_RESULT `hl7:"GROUP"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_TIMING_QTY - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_TIMING_QTY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_RESULT - Group struct
type OPU_R25_ACCESSION_DETAIL_SPECIMEN_ORDER_RESULT struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OPU_R25 - Unsolicited Population/Location-Based Laboratory Observation 
// https://hl7-definition.caristix.com/v2/HL7v2.6/TriggerEvents/OPU_R25
type OPU_R25 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredential UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments1 NTE `hl7:"TAG=NTE;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	NotesAndComments2 []NTE `hl7:"TAG=NTE;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL"`
	AccessionDetail []OPU_R25_ACCESSION_DETAIL `hl7:"GROUP"`
}

