package hl7v28

// OSM_R26_SHIPMENT - Group struct
type OSM_R26_SHIPMENT struct {
	Shipment SHP `hl7:"TAG=SHP"`
	ParticipationInformation []PRT `hl7:"TAG=PRT"`
	Shipping_observation []OSM_R26_SHIPMENT_SHIPPING_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Package []OSM_R26_SHIPMENT_PACKAGE `hl7:"GROUP"`
}

// OSM_R26_SHIPMENT_SHIPPING_OBSERVATION - Group struct
type OSM_R26_SHIPMENT_SHIPPING_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE - Group struct
type OSM_R26_SHIPMENT_PACKAGE struct {
	ShipmentPackage PAC `hl7:"TAG=PAC"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Specimen []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN `hl7:"GROUP;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN struct {
	Specimen SPM `hl7:"TAG=SPM"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	Specimen_observation []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SPECIMEN_OBSERVATION `hl7:"GROUP;ATR=optional"`
	Container []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER `hl7:"GROUP;ATR=optional"`
	Subject_person_animal_identification OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION `hl7:"GROUP;ATR=optional"`
	Subject_population_location_identification OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION `hl7:"GROUP;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SPECIMEN_OBSERVATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SPECIMEN_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER struct {
	SpecimenContainerDetail SAC `hl7:"TAG=SAC"`
	Container_observation []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_CONTAINER_CONTAINER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION struct {
	PatientIdentification PID `hl7:"TAG=PID"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
	AccessRestriction []ARV `hl7:"TAG=ARV;ATR=optional"`
	Patient_observation []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION_PATIENT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION_PATIENT_OBSERVATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_PERSON_ANIMAL_IDENTIFICATION_PATIENT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION struct {
	PatientVisit PV1 `hl7:"TAG=PV1"`
	ParticipationInformation1 []PRT `hl7:"TAG=PRT;ATR=optional"`
	Patient_visit_observation []OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION_PATIENT_VISIT_OBSERVATION `hl7:"GROUP;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID;ATR=optional"`
	ParticipationInformation2 []PRT `hl7:"TAG=PRT;ATR=optional"`
	NextOfKinAssociatedParties []NK1 `hl7:"TAG=NK1;ATR=optional"`
}

// OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION_PATIENT_VISIT_OBSERVATION - Group struct
type OSM_R26_SHIPMENT_PACKAGE_SPECIMEN_SUBJECT_POPULATION_LOCATION_IDENTIFICATION_PATIENT_VISIT_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	ParticipationInformation []PRT `hl7:"TAG=PRT;ATR=optional"`
}

// OSM_R26 - Unsolicited Specimen Shipment Manifest Message
// https://hl7-definition.caristix.com/v2/HL7v2.8/TriggerEvents/OSM_R26
type OSM_R26 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	UserAuthenticationCredentialSegment UAC `hl7:"TAG=UAC;ATR=optional"`
	Shipment []OSM_R26_SHIPMENT `hl7:"GROUP"`
}

