package hl7v26

import "time"

// SCD - Anti-Microbial Cycle Data
// https://hl7-definition.caristix.com/v2/HL7v2.6/Segments/SCD
type SCD struct {
	CycleStartTime string `hl7:"POS=2"` // TODO - handle TM type parsing
	CycleCount *int `hl7:"POS=3"`
	TempMax CQ `hl7:"POS=4"`
	TempMin CQ `hl7:"POS=5"`
	LoadNumber *int `hl7:"POS=6"`
	ConditionTime CQ `hl7:"POS=7"`
	SterilizeTime CQ `hl7:"POS=8"`
	ExhaustTime CQ `hl7:"POS=9"`
	TotalCycleTime CQ `hl7:"POS=10"`
	DeviceStatus CWE `hl7:"POS=11"`
	CycleStartDateTime time.Time `hl7:"POS=12"`
	DryTime CQ `hl7:"POS=13"`
	LeakRate CQ `hl7:"POS=14"`
	ControlTemperature CQ `hl7:"POS=15"`
	SterilizerTemperature CQ `hl7:"POS=16"`
	CycleCompleteTime string `hl7:"POS=17"` // TODO - handle TM type parsing
	UnderTemperature CQ `hl7:"POS=18"`
	OverTemperature CQ `hl7:"POS=19"`
	AbortCycle CNE `hl7:"POS=20"`
	Alarm CNE `hl7:"POS=21"`
	LongInChargePhase CNE `hl7:"POS=22"`
	LongInExhaustPhase CNE `hl7:"POS=23"`
	LongInFastExhaustPhase CNE `hl7:"POS=24"`
	Reset CNE `hl7:"POS=25"`
	OperatorUnload XCN `hl7:"POS=26"`
	DoorOpen CNE `hl7:"POS=27"`
	ReadingFailure CNE `hl7:"POS=28"`
	CycleType CWE `hl7:"POS=29"`
	ThermalRinseTime CQ `hl7:"POS=30"`
	WashTime CQ `hl7:"POS=31"`
	InjectionRate CQ `hl7:"POS=32"`
	ProcedureCode CNE `hl7:"POS=33"`
	PatientIdentifierList []CX `hl7:"POS=34"`
	AttendingDoctor XCN `hl7:"POS=35"`
	DilutionFactor SN `hl7:"POS=36"`
	FillTime CQ `hl7:"POS=37"`
	InletTemperature CQ `hl7:"POS=38"`
}

