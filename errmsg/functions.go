package errmsg

import "errors"

// Lining
var (
	ErrLineProcessingEmptyInput       = errors.New("empty input")
	ErrLineProcessingInvalidLinebreak = errors.New("invalid line breaking")
	ErrLineProcessingNoLineSeparator  = errors.New("separator has to be provided if auto-detect is disabled")
)

// AnnotationParsing
var (
	ErrAnnotationParsingInvalidProtocol    = errors.New("invalid protocol")
	ErrAnnotationParsingMissingAnnotation  = errors.New("annotation missing")
	ErrAnnotationParsingInvalidElement     = errors.New("invalid annotation element")
	ErrAnnotationParsingInvalidElementKey  = errors.New("invalid annotation element key")
	ErrAnnotationParsingIllegal            = errors.New("illegal annotation combination")
	ErrAnnotationParsingInvalidInputStruct = errors.New("invalid input struct type")
)

// LineParsing
var (
	ErrLineParsingEmptyInput                    = errors.New("empty input")
	ErrLineParsingHeaderTooShort                = errors.New("header too short")
	ErrLineParsingMandatoryInputFieldsMissing   = errors.New("mandatory input fields missing")
	ErrLineParsingSequenceNumberMismatch        = errors.New("sequence number mismatch")
	ErrLineParsingRequiredInputFieldMissing     = errors.New("required input field missing")
	ErrLineParsingInputComponentsMissing        = errors.New("input components missing")
	ErrLineParsingNonSettableField              = errors.New("field is not settable")
	ErrLineParsingDataParsingError              = errors.New("data parsing error")
	ErrLineParsingInvalidDateFormat             = errors.New("invalid date format")
	ErrLineParsingUnsupportedDataType           = errors.New("unsupported data type")
	ErrLineParsingReservedFieldPosReference     = errors.New("field position 1 and 2 are reserved")
	ErrLineParsingUnterminatedEscapeSequence    = errors.New("unterminated escape sequence")
	ErrLineParsingUnknownEscapeSequence         = errors.New("unknown escape sequence")
	ErrLineParsingMaximumRecursionDepthExceeded = errors.New("line parsing maximum recursion depth exceeded")
	ErrLineParsingInvalidRecursionDepth         = errors.New("invalid recursion depth")
)

// StructureParsing
var (
	ErrStructureParsingMaxDepthReached    = errors.New("max depth reached")
	ErrStructureParsingInputLinesDepleted = errors.New("input lines depleted")
	ErrStructureParsingStructureMismatch  = errors.New("structure mismatch")
)

// LineBuilding
var (
	ErrLineBuildingInvalidDateFormat             = errors.New("invalid date format")
	ErrLineBuildingUsupportedDataType            = errors.New("unsupported data type")
	ErrLineBuildingReservedFieldPosReference     = errors.New("field position 1 and 2 are reserved")
	ErrLineBuildingInvalidLengthAttributeValue   = errors.New("invalid length attribute value")
	ErrLineBuildingMaximumRecursionDepthExceeded = errors.New("line building maximum recursion depth exceeded")
	ErrLineBuildingInvalidRecursionDepth         = errors.New("invalid recursion depth")
)
