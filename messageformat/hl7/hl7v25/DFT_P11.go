package hl7v25

// DFT_P11_COMMON_ORDER - Group struct
type DFT_P11_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	TimingQuantity []DFT_P11_COMMON_ORDER_TIMING_QUANTITY `hl7:"GROUP;ATR=optional"`
	Order DFT_P11_COMMON_ORDER_ORDER `hl7:"GROUP;ATR=optional"`
	Observation []DFT_P11_COMMON_ORDER_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// DFT_P11_COMMON_ORDER_TIMING_QUANTITY - Group struct
type DFT_P11_COMMON_ORDER_TIMING_QUANTITY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// DFT_P11_COMMON_ORDER_ORDER - Group struct
type DFT_P11_COMMON_ORDER_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DFT_P11_COMMON_ORDER_OBSERVATION - Group struct
type DFT_P11_COMMON_ORDER_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DFT_P11_INSURANCE - Group struct
type DFT_P11_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification []IN3 `hl7:"TAG=IN3;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// DFT_P11_FINANCIAL - Group struct
type DFT_P11_FINANCIAL struct {
	FinancialTransaction FT1 `hl7:"TAG=FT1"`
	FinancialProcedure []DFT_P11_FINANCIAL_FINANCIAL_PROCEDURE `hl7:"GROUP;ATR=optional"`
	FinancialCommonOrder []DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER `hl7:"GROUP;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	FinancialInsurance []DFT_P11_FINANCIAL_FINANCIAL_INSURANCE `hl7:"GROUP;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_PROCEDURE - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_PROCEDURE struct {
	Procedures PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER struct {
	CommonOrder ORC `hl7:"TAG=ORC;ATR=optional"`
	FinancialTimingQuantity []DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_TIMING_QUANTITY `hl7:"GROUP;ATR=optional"`
	FinancialOrder DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_ORDER `hl7:"GROUP;ATR=optional"`
	FinancialObservation []DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_OBSERVATION `hl7:"GROUP;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_TIMING_QUANTITY - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_TIMING_QUANTITY struct {
	TimingQuantity TQ1 `hl7:"TAG=TQ1"`
	TimingQuantityRelationship []TQ2 `hl7:"TAG=TQ2;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_ORDER - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_ORDER struct {
	ObservationRequest OBR `hl7:"TAG=OBR"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_OBSERVATION - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_COMMON_ORDER_FINANCIAL_OBSERVATION struct {
	ObservationResult OBX `hl7:"TAG=OBX"`
	NotesAndComments []NTE `hl7:"TAG=NTE;ATR=optional"`
}

// DFT_P11_FINANCIAL_FINANCIAL_INSURANCE - Group struct
type DFT_P11_FINANCIAL_FINANCIAL_INSURANCE struct {
	Insurance IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformation IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertification []IN3 `hl7:"TAG=IN3;ATR=optional"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// DFT_P11 - Post Detail Financial Transactions - Expanded
// https://hl7-definition.caristix.com/v2/HL7v2.5/TriggerEvents/DFT_P11
type DFT_P11 struct {
	MessageHeader MSH `hl7:"TAG=MSH"`
	SoftwareSegment []SFT `hl7:"TAG=SFT;ATR=optional"`
	EventType EVN `hl7:"TAG=EVN"`
	PatientIdentification PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	Role1 []ROL `hl7:"TAG=ROL;ATR=optional"`
	PatientVisit PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformation PV2 `hl7:"TAG=PV2;ATR=optional"`
	Role2 []ROL `hl7:"TAG=ROL;ATR=optional"`
	Disability []DB1 `hl7:"TAG=DB1;ATR=optional"`
	CommonOrder []DFT_P11_COMMON_ORDER `hl7:"GROUP;ATR=optional"`
	Diagnosis []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroup DRG `hl7:"TAG=DRG;ATR=optional"`
	Guarantor []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []DFT_P11_INSURANCE `hl7:"GROUP;ATR=optional"`
	Accident ACC `hl7:"TAG=ACC;ATR=optional"`
	Financial []DFT_P11_FINANCIAL `hl7:"GROUP"`
}

