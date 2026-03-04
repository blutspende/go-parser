package hl7v22

// ADR_A19_QUERY_RESPONSE - Group struct
type ADR_A19_QUERY_RESPONSE struct {
	EventType EVN `hl7:"TAG=EVN;ATR=optional"`
	PatientIdentification PID `hl7:"TAG=PID"`
	NextOfKin []NK1 `hl7:"TAG=NK1;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	ObservationResult []OBX `hl7:"TAG=OBX;ATR=optional"`
	PatientAllergyInformation []AL1 `hl7:"TAG=AL1;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	Procedures []PR1 `hl7:"TAG=PR1;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []ADR_A19_QUERY_RESPONSE_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Ub82Data UB1 `hl7:"TAG=UB1;ATR=optional"`
	Ub92Data UB2 `hl7:"TAG=UB2;ATR=optional"`
}

// ADR_A19_QUERY_RESPONSE_INSURANCE - Group struct
type ADR_A19_QUERY_RESPONSE_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInfo IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInfoCertification IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// ADR_A19 - Patient query
// https://hl7-definition.caristix.com/v2/HL7v2.2/TriggerEvents/ADR_A19
type ADR_A19 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	MessageAcknowledgment MSA `hl7:"TAG=MSA"`
	Error ERR `hl7:"TAG=ERR;ATR=optional"`
	QueryDefinition QRD `hl7:"TAG=QRD"`
	QueryResponse []ADR_A19_QUERY_RESPONSE `hl7:"GROUP"`
	ContinuationPointer DSC `hl7:"TAG=DSC;ATR=optional"`
}

