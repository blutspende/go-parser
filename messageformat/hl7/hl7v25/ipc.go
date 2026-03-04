package hl7v25

// IPC - Imaging Procedure Control Segment
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/IPC
type IPC struct {
	AccessionIdentifier EI `hl7:"POS=2"`
	RequestedProcedureID EI `hl7:"POS=3"`
	StudyInstanceUid EI `hl7:"POS=4"`
	ScheduledProcedureStepID EI `hl7:"POS=5"`
	Modality CE `hl7:"POS=6"`
	ProtocolCode []CE `hl7:"POS=7"`
	ScheduledStationName EI `hl7:"POS=8"`
	ScheduledProcedureStepLocation []CE `hl7:"POS=9"`
	ScheduledAeTitle string `hl7:"POS=10"`
}

