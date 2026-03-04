package hl7v25

// ERR - Error
// https://hl7-definition.caristix.com/v2/HL7v2.5/Segments/ERR
type ERR struct {
	ErrorCodeAndLocation []ELD `hl7:"POS=2"`
	ErrorLocation []ERL `hl7:"POS=3"`
	Hl7ErrorCode CWE `hl7:"POS=4"`
	Severity string `hl7:"POS=5"`
	ApplicationErrorCode CWE `hl7:"POS=6"`
	ApplicationErrorParameter []string `hl7:"POS=7"`
	DiagnosticInformation string `hl7:"POS=8"`
	UserMessage string `hl7:"POS=9"`
	InformPersonIndicator []string `hl7:"POS=10"`
	OverrideType CWE `hl7:"POS=11"`
	OverrideReasonCode []CWE `hl7:"POS=12"`
	HelpDeskContactPoint []XTN `hl7:"POS=13"`
}

