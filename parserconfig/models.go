package parserconfig

import (
	"time"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/timezone"
	"github.com/blutspende/go-parser/enums/lineseparator"
	"github.com/blutspende/go-parser/enums/notation"
)

// Configuration struct for the whole process
type Configuration struct {
	Protocol                   Protocol
	Encoding                   encoding.Encoding
	LineSeparator              string
	AutoDetectLineSeparator    bool
	TimeZone                   timezone.TimeZone
	EnforceSequenceNumberCheck bool
	Notation                   string
	DefaultDecimalPrecision    int
	RoundLastDecimal           bool
	KeepShortDateTimeZone      bool
	EscapeOutputStrings        bool
	EnforceMessageCompleteness bool
	DropEmptyOutputRecords     bool
	Delimiters                 Delimiters
	TimeLocation               *time.Location
}

var DefaultConfigurationASTM = Configuration{
	Protocol:                   ASTM,
	Encoding:                   encoding.UTF8,
	LineSeparator:              lineseparator.LF,
	AutoDetectLineSeparator:    true,
	TimeZone:                   timezone.EuropeBerlin,
	EnforceSequenceNumberCheck: true,
	Notation:                   notation.Standard,
	DefaultDecimalPrecision:    3,
	RoundLastDecimal:           true,
	KeepShortDateTimeZone:      true,
	EscapeOutputStrings:        false,
	EnforceMessageCompleteness: true,
	DropEmptyOutputRecords:     false,
	Delimiters:                 DefaultDelimitersASTM,
	TimeLocation:               nil,
}

var DefaultConfigurationHL7 = Configuration{
	Protocol:                   HL7,
	Encoding:                   encoding.UTF8,
	LineSeparator:              lineseparator.CR,
	AutoDetectLineSeparator:    true,
	TimeZone:                   timezone.EuropeBerlin,
	EnforceSequenceNumberCheck: true,
	Notation:                   notation.Short,
	DefaultDecimalPrecision:    3,
	RoundLastDecimal:           true,
	KeepShortDateTimeZone:      true,
	EscapeOutputStrings:        false,
	EnforceMessageCompleteness: false,
	DropEmptyOutputRecords:     true,
	Delimiters:                 DefaultDelimitersHL7,
	TimeLocation:               nil,
}

type Protocol string

const ASTM Protocol = "ASTM"
const HL7 Protocol = "HL7"

// Delimiters used in ASTM parsing
type Delimiters struct {
	Field        string
	Repeat       string
	Component    string
	SubComponent string
	Escape       string
}

var DefaultDelimitersASTM = Delimiters{
	Field:     `|`,
	Repeat:    `\`,
	Component: `^`,
	Escape:    `&`,
}
var DefaultDelimitersHL7 = Delimiters{
	Field:        `|`,
	Component:    `^`,
	Repeat:       `~`,
	Escape:       `\`,
	SubComponent: `&`,
}
