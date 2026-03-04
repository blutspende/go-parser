package hl7v231

// DFT_P03_FINANCIAL - Group struct
type DFT_P03_FINANCIAL struct {
	FinancialTransactionSegment FT1 `hl7:"TAG=FT1"`
	FinancialProcedure []DFT_P03_FINANCIAL_FINANCIAL_PROCEDURE `hl7:"GROUP;ATR=optional"`
}

// DFT_P03_FINANCIAL_FINANCIAL_PROCEDURE - Group struct
type DFT_P03_FINANCIAL_FINANCIAL_PROCEDURE struct {
	ProceduresSegment PR1 `hl7:"TAG=PR1"`
	Role []ROL `hl7:"TAG=ROL;ATR=optional"`
}

// DFT_P03_INSURANCE - Group struct
type DFT_P03_INSURANCE struct {
	InsuranceSegment IN1 `hl7:"TAG=IN1"`
	InsuranceAdditionalInformationSegment IN2 `hl7:"TAG=IN2;ATR=optional"`
	InsuranceAdditionalInformationCertificationSegment []IN3 `hl7:"TAG=IN3;ATR=optional"`
}

// DFT_P03 - Post detail financial transaction
// https://hl7-definition.caristix.com/v2/HL7v2.3.1/TriggerEvents/DFT_P03
type DFT_P03 struct {
	MessageHeaderSegment MSH `hl7:"TAG=MSH"`
	EventTypeSegment EVN `hl7:"TAG=EVN"`
	PatientIdentificationSegment PID `hl7:"TAG=PID"`
	PatientAdditionalDemographic PD1 `hl7:"TAG=PD1;ATR=optional"`
	PatientVisitSegment PV1 `hl7:"TAG=PV1;ATR=optional"`
	PatientVisitAdditionalInformationSegment PV2 `hl7:"TAG=PV2;ATR=optional"`
	DisabilitySegment []DB1 `hl7:"TAG=DB1;ATR=optional"`
	ObservationResultSegment []OBX `hl7:"TAG=OBX;ATR=optional"`
	Financial []DFT_P03_FINANCIAL `hl7:"GROUP"`
	DiagnosisSegment []DG1 `hl7:"TAG=DG1;ATR=optional"`
	DiagnosisRelatedGroupSegment DRG `hl7:"TAG=DRG;ATR=optional"`
	GuarantorSegment []GT1 `hl7:"TAG=GT1;ATR=optional"`
	Insurance []DFT_P03_INSURANCE `hl7:"GROUP;ATR=optional"`
	AccidentSegment ACC `hl7:"TAG=ACC;ATR=optional"`
}

