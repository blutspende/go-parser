package hl7v23

// ADR_A19_QUERY_RESPONSE - Group struct
type ADR_A19_QUERY_RESPONSE struct {
	EventType EVN `hl7:"TAG=EVN;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Procedure []ADR_A19_QUERY_RESPONSE_PROCEDURE `hl7:"GROUP;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []ADR_A19_QUERY_RESPONSE_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82Data UB1 `hl7:"TAG=UB1;ATR=optional"`
	Ub92Data UB2 `hl7:"TAG=UB2;ATR=optional"`
}

// ADR_A19_QUERY_RESPONSE_PROCEDURE - Group struct
type ADR_A19_QUERY_RESPONSE_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// ADR_A19_QUERY_RESPONSE_INSURANCE - Group struct
type ADR_A19_QUERY_RESPONSE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ADR_A19 - Patient query
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ADR_A19
type ADR_A19 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	MessageAcknowledgementSegment MSA `hl7:"TAG=MSA"`
	ErrorSegment ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinitionSegment QRD `hl7:"TAG=QRD"`
	QueryFilterSegment QRF `hl7:"TAG=QRF;ATR=optional"`
	QueryResponse []ADR_A19_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointerSegment DSC `hl7:"TAG=DSC;ATR=optional"`
}

