package hl7v27

// OML_O39_PATIENT - Group struct
type OML_O39_PATIENT struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
	Patient_visit OML_O39_PATIENT_PATIENT_VISIT `hl7:"GROUP;ATR=optional"`
	Insurance []OML_O39_PATIENT_INSURANCE `hl7:"GROUP;ATR=optional"`
	Guarantor GT1 `hl7:"TAG=GT1;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
}

// OML_O39_PATIENT_PATIENT_VISIT - Group struct
type OML_O39_PATIENT_PATIENT_VISIT struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O39_PATIENT_INSURANCE - Group struct
type OML_O39_PATIENT_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// OML_O39_ORDER - Group struct
type OML_O39_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Timing []OML_O39_ORDER_TIMING `hl7:"GROUP;ATR=optional"`
	Observation_request OML_O39_ORDER_OBSERVATION_REQUEST `hl7:"GROUP;ATR=optional"`
	FinancialTransaction []FT1 `hl7:"TAG=FT1;ATR=optional"`
	ClinicalTrialIdentification []CTI `hl7:"TAG=CTI;ATR=optional"`
	Billing BLG `hl7:"TAG=BLG;ATR=optional"`
}

// OML_O39_ORDER_TIMING - Group struct
type OML_O39_ORDER_TIMING struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	ContactData CTD `hl7:"TAG=CTD;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Observation []OML_O39_ORDER_OBSERVATION_REQUEST_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Specimen_shipment []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT `hl7:"GROUP;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_OBSERVATION - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	TestCodeDetail TCD `hl7:"TAG=TCD;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT struct {
	Shipment SHP `hl7:"TAG=SHP"`
	Shipment_observation []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_SHIPMENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Package []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE `hl7:"GROUP"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_SHIPMENT_OBSERVATION - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_SHIPMENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE struct {
	ShipmentPackage PAC `hl7:"TAG=PAC"`
	Specimen_in_package []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE `hl7:"GROUP;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE struct {
	Specimen SPM `hl7:"TAG=SPM"`
	Specimen_observation []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Specimen_container_in_package []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE `hl7:"GROUP;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_OBSERVATION - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Container_observation []OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE_CONTAINER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE_CONTAINER_OBSERVATION - Group struct
type OML_O39_ORDER_OBSERVATION_REQUEST_SPECIMEN_SHIPMENT_PACKAGE_SPECIMEN_IN_PACKAGE_SPECIMEN_CONTAINER_IN_PACKAGE_CONTAINER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OML_O39 - Specimen shipment centric laboratory order
// https://hl7-definition.caristix.com/v2/HL7v2.7/TriggerEvents/OML_O39
type OML_O39 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
	Patient OML_O39_PATIENT `hl7:"GROUP;ATR=optional"`
	Order []OML_O39_ORDER `hl7:"GROUP"`
}

