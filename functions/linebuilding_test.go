package functions

import (
	"testing"
	"time"

	"github.com/blutspende/go-parser/enums/notation"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/stretchr/testify/assert"
)

// Note: structures come from functions_test.go

// BuildLine tests - basics
func TestBuildLine_SimpleRecord(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{
		First:  "first",
		Second: "second",
		Third:  "third",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|second|third", result)
}
func TestBuildLine_UnorderedRecord(t *testing.T) {
	// Arrange
	source := UnorderedRecord{
		First:  "first",
		Second: "second",
		Third:  "third",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|second|third", result)
}
func TestBuildLine_EnumRecord(t *testing.T) {
	// Arrange
	source := EnumRecord{
		Enum: "enum",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|enum", result)
}
func TestBuildLine_MultitypeRecord(t *testing.T) {
	// Arrange
	source := MultitypeRecord{
		String:  "string",
		Int:     3,
		Float32: 3.14,
		Float64: 3.14159265,
		Time:    time.Date(2006, 1, 2, 15, 26, 37, 0, time.UTC),
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|string|3|3.140|3.142|20060102162637", result)
}

// BuildLine tests - float length and precision
func TestBuildLine_FloatLengthRounded(t *testing.T) {
	// Arrange
	source := FloatLengthRecord{
		Default:    3.14159265,
		Length0:    3.14159265,
		Length4:    3.14159265,
		LengthFull: 3.14159265,
	}
	config.RoundLastDecimal = true
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|3.142|3|3.1416|3.14159265", result)
	// Teardown
	teardown()
}
func TestBuildLine_FloatLengthTruncated(t *testing.T) {
	// Arrange
	source := FloatLengthRecord{
		Default:    3.14159265,
		Length0:    3.14159265,
		Length4:    3.14159265,
		LengthFull: 3.14159265,
	}
	config.RoundLastDecimal = false
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|3.141|3|3.1415|3.14159265", result)
	// Teardown
	teardown()
}
func TestBuildLine_FloatLengthDefault(t *testing.T) {
	// Arrange
	source := FloatLengthRecord{
		Default:    0,
		Length0:    0,
		Length4:    0,
		LengthFull: 0,
	}
	config.DefaultDecimalPrecision = 2
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|0.00|0|0.0000|0", result)
	// Teardown
	teardown()
}

// BuildLine tests - pointers and empty values
func TestBuildLine_MultitypeRecordEmpty(t *testing.T) {
	// Arrange
	source := MultitypeRecord{}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1||0|0.000|0.000|", result)
}
func TestBuildLine_MultitypeLengthRecordEmpty(t *testing.T) {
	// Arrange
	source := FloatLengthRecord{}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|0.000|0|0.0000|0", result)
}
func TestBuildLine_MultitypePointerRecord(t *testing.T) {
	// Arrange
	String := "string"
	Int := 3
	Float32 := float32(3.14)
	Float64 := 3.14159265
	Time := time.Date(2006, 1, 2, 15, 26, 37, 0, time.UTC)
	source := MultitypePointerRecord{
		String:  &String,
		Int:     &Int,
		Float32: &Float32,
		Float64: &Float64,
		Time:    &Time,
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|string|3|3.140|3.142|20060102162637", result)
}
func TestBuildLine_MultitypePointerRecordEmpty(t *testing.T) {
	// Arrange
	source := MultitypePointerRecord{}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|||||", result)
}

// BuildLine tests - missing data and short notation
func TestBuildLine_MissingData(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{
		First:  "first",
		Second: "",
		Third:  "third",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first||third", result)
}
func TestBuildLine_MissingDataAtEndStandardNotation(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{
		First:  "first",
		Second: "",
		Third:  "",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first||", result)
}
func TestBuildLine_MissingDataAtEndShortNotation(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{
		First:  "first",
		Second: "",
		Third:  "",
	}
	config.Notation = notation.Short
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first", result)
	// Teardown
	teardown()
}
func TestBuildLine_MissingDataAtEndWithComponentsShortNotation(t *testing.T) {
	// Arrange
	source := ComponentedRecord{
		First: "first",
	}
	config.Notation = notation.Short
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first", result)
	// Teardown
	teardown()
}
func TestBuildLine_EmptyRecord_HL7(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{}
	// Act
	result, err := BuildLine(source, "REC", 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "REC|", result)
}

// BuildLine tests - header and sequence
func TestBuildLine_HeaderRecord_ASTM(t *testing.T) {
	// Arrange
	source := HeaderRecord{
		First: "first",
	}
	// Act
	result, err := BuildLine(source, "H", 0, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "H|\\^&|first", result)
}
func TestBuildLine_HeaderRecord_HL7(t *testing.T) {
	// Arrange
	source := HeaderRecord{
		First: "first",
	}
	// Act
	result, err := BuildLine(source, "MSH", 0, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "MSH|^~\\&|first", result)
}
func TestBuildLine_HeaderRecordCustomDelimiters_ASTM(t *testing.T) {
	// Arrange
	source := HeaderRecord{
		First: "first",
	}
	config.Delimiters.Field = "/"
	config.Delimiters.Repeat = "!"
	config.Delimiters.Component = "*"
	config.Delimiters.Escape = "%"
	// Act
	result, err := BuildLine(source, "H", 0, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "H/!*%/first", result)
	// Teardown
	teardown()
	assert.Equal(t, "^", config.Delimiters.Component)
}
func TestBuildLine_HeaderRecordCustomDelimiters_HL7(t *testing.T) {
	// Arrange
	source := HeaderRecord{
		First: "first",
	}
	configHL7.Delimiters.Field = "/"
	configHL7.Delimiters.Component = "*"
	configHL7.Delimiters.Repeat = "!"
	configHL7.Delimiters.Escape = "%"
	configHL7.Delimiters.SubComponent = "?"
	// Act
	result, err := BuildLine(source, "MSH", 0, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "MSH/*!%?/first", result)
	// Teardown
	teardown()
}
func TestBuildLine_HeaderDelimiterChange_ASTM(t *testing.T) {
	// Arrange
	source := HeaderDelimiterChange{
		First: "first",
		Array: []string{"second1", "second2"},
		Comp1: "third1",
		Comp2: "third2",
	}
	config.Delimiters.Field = "/"
	config.Delimiters.Repeat = "!"
	config.Delimiters.Component = "*"
	config.Delimiters.Escape = "%"
	// Act
	result, err := BuildLine(source, "H", 0, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "H/!*%/first/second1!second2/third1*third2", result)
	// Teardown
	teardown()
}
func TestBuildLine_TagAndSequence_ASTM(t *testing.T) {
	// Arrange
	source := ThreeFieldRecord{
		First:  "first",
		Second: "second",
		Third:  "third",
	}
	// Act
	result, err := BuildLine(source, "D", 3, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "D|3|first|second|third", result)
}
func TestBuildLine_Sequence_HL7(t *testing.T) {
	// Arrange
	source := SequenceHl7{
		Data: "data",
	}
	// Act
	result, err := BuildLine(source, "SEC", 4, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "SEC||4|data", result)
}
func TestBuildLine_ReservedFieldRecord_HL7(t *testing.T) {
	// Arrange
	source := ReservedFieldRecordHL7{}
	// Act
	result, err := BuildLine(source, "T", 1, configHL7)
	// Assert
	assert.EqualError(t, err, errmsg.ErrLineBuildingReservedFieldPosReference.Error())
	assert.Equal(t, "", result)
}
func TestBuildLine_ReservedFieldRecord_ASTM(t *testing.T) {
	// Arrange
	source := ReservedFieldRecord{}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrLineBuildingReservedFieldPosReference.Error())
	assert.Equal(t, "", result)
}

// BuildLine tests - arrays, components, substructures and sub-substructures
func TestBuildLine_ArrayRecord(t *testing.T) {
	// Arrange
	source := ArrayRecord{
		First: "first",
		Array: []string{"second1", "second2", "second3"},
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|second1\\second2\\second3", result)
}
func TestBuildLine_ComponentedRecord(t *testing.T) {
	// Arrange
	source := ComponentedRecord{
		First:       "first",
		SecondComp1: "second1",
		SecondComp2: "second2",
		ThirdComp1:  "third1",
		ThirdComp2:  "third2",
		ThirdComp3:  "third3",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|second1^second2|third1^third2^third3", result)
}
func TestBuildLine_ComponentedRecordShortNotation(t *testing.T) {
	// Arrange
	source := ComponentedRecord{
		First:       "first",
		SecondComp1: "second1",
		SecondComp2: "second2",
		ThirdComp1:  "third1",
		ThirdComp2:  "third2",
		ThirdComp3:  "third3",
	}
	config.Notation = notation.Short
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|second1^second2|third1^third2^third3", result)
	// Teardown
	teardown()
}
func TestBuildLine_SubstructureRecord(t *testing.T) {
	// Arrange
	source := SubstructureRecord{
		First: "first",
		Second: SubstructureField{
			FirstComponent:  "firstComponent",
			SecondComponent: "secondComponent",
			ThirdComponent:  "thirdComponent",
		},
		Third: "third",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|firstComponent^secondComponent^thirdComponent|third", result)
}
func TestBuildLine_SubstructureRecordMissingData(t *testing.T) {
	// Arrange
	source := SubstructureRecord{
		Second: SubstructureField{
			FirstComponent: "firstComponent",
		},
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1||firstComponent^^|", result)
}
func TestBuildLine_SubstructureRecordMissingDataShortNotation(t *testing.T) {
	// Arrange
	source := SubstructureRecord{
		Second: SubstructureField{
			FirstComponent: "firstComponent",
		},
	}
	config.Notation = notation.Short
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1||firstComponent", result)
	// Teardown
	teardown()
}
func TestBuildLine_SparseSubstructureRecord(t *testing.T) {
	// Arrange
	source := SparseSubstructureRecord{
		First: "first",
		Second: SparseSubstructureField{
			Component1: "component1",
			Component3: "component3",
			Component6: "component6",
		},
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|component1^^component3^^^component6", result)
}
func TestBuildLine_SubstructureArrayRecord(t *testing.T) {
	// Arrange
	source := SubstructureArrayRecord{
		First: "first",
		Second: []SubstructureField{
			SubstructureField{
				FirstComponent:  "r1c1",
				SecondComponent: "r1c2",
				ThirdComponent:  "r1c3",
			},
			SubstructureField{
				FirstComponent:  "r2c1",
				SecondComponent: "r2c2",
				ThirdComponent:  "r2c3",
			},
		},
		Third: "third",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|r1c1^r1c2^r1c3\\r2c1^r2c2^r2c3|third", result)
}
func TestBuildLine_SubSub_ASTM(t *testing.T) {
	// Arrange
	source := SubSubRecord{}
	// Act
	_, err := BuildLine(source, "SUB", 1, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrLineBuildingMaximumRecursionDepthExceeded.Error())
}
func TestBuildLine_SubSub_HL7(t *testing.T) {
	// Arrange
	source := SubSubRecord{
		First: "field1",
		Second: SubField{
			First: "comp1",
			Second: SubSubField{
				First:  "sub1",
				Second: "sub2",
				Third:  "sub3",
			},
			Third: "comp3",
		},
		Third: "field3",
	}
	expected := "SUB||field1|comp1^sub1&sub2&sub3^comp3|field3"
	// Act
	result, err := BuildLine(source, "SUB", 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

// BuildLine tests - time and date
func TestBuildLine_TimeRecord(t *testing.T) {
	// Arrange
	source := TimeRecord{
		Time: time.Date(2006, 03, 04, 16, 44, 29, 0, time.UTC),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304174429", result)
}
func TestBuildLine_DateRecord(t *testing.T) {
	// Arrange
	source := DateRecord{
		Date: time.Date(2006, 03, 04, 0, 0, 0, 0, config.TimeLocation),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304", result)
}
func TestBuildLine_TimeRecord_DateInput(t *testing.T) {
	// Arrange
	source := TimeRecord{
		Time: time.Date(2006, 03, 04, 0, 0, 0, 0, config.TimeLocation),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304000000", result)
}
func TestBuildLine_DateRecord_TimeInput(t *testing.T) {
	// Arrange
	source := DateRecord{
		Date: time.Date(2006, 03, 04, 16, 44, 29, 0, time.UTC),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304", result)
}
func TestBuildLine_TimeRecord_LocalTime(t *testing.T) {
	// Arrange
	source := TimeRecord{
		Time: time.Date(2006, 03, 04, 16, 44, 29, 0, config.TimeLocation),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304164429", result)
}
func TestBuildLine_DateRecord_UTC(t *testing.T) {
	// Arrange
	source := DateRecord{
		Date: time.Date(2006, 03, 04, 0, 0, 0, 0, time.UTC),
	}
	// Act
	result, err := BuildLine(source, "R", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R|1|20060304", result)
}

// BuildLine tests - special and wrong annotations
func TestBuildLine_SparseFieldRecord(t *testing.T) {
	// Arrange
	source := SparseFieldRecord{
		Field3: "field3",
		Field5: "field5",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|field3||field5", result)
}
func TestBuildLine_WrongComponentOrder(t *testing.T) {
	// Arrange
	source := WrongComponentOrderRecord{
		First: "first",
		Comp2: "comp2",
		Comp1: "comp1",
		Comp3: "comp3",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|first|comp1^comp2^comp3", result)
}
func TestBuildLine_WrongComponentPlacement(t *testing.T) {
	// Arrange
	source := WrongComponentPlacementRecord{
		Field1: "field1",
		Comp1:  "comp1",
		Field2: "field2",
		Comp2:  "comp2",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|field1|comp1^comp2|field2", result)
}
func TestBuildLine_MultipleWrongComponentPlacement(t *testing.T) {
	// Arrange
	source := MultipleWrongComponentPlacementRecord{
		Field3: "field3",
		Comp41: "comp41",
		Field5: "field5",
		Comp62: "comp62",
		Comp42: "comp42",
		Field7: "field7",
		Comp61: "comp61",
		Field8: "field8",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|field3|comp41^comp42|field5|comp61^comp62|field7|field8", result)
}
func TestBuildLine_MissingAnnotation(t *testing.T) {
	// Arrange
	source := MissingAnnotationRecord{
		Field3:  "field3",
		Missing: "missing",
		Field4:  "field4",
	}
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|field3|field4", result)
}
func TestBuildLine_InvalidAttributeValue(t *testing.T) {
	// Arrange
	source := InvalidAttributeValueRecord{
		First: 3.14159,
	}
	// Act
	_, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrLineBuildingInvalidLengthAttributeValue.Error())
}

// BuildLine tests - escaped characters
func TestBuildLine_EscapedChars(t *testing.T) {
	// Arrange
	source := SimpleRecord{
		First: "esc^ape",
	}
	config.EscapeOutputStrings = true
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|esc&^ape", result)
	// Teardown
	teardown()
}
func TestBuildLine_EscapedCharsNoEscape(t *testing.T) {
	// Arrange
	source := SimpleRecord{
		First: "esc^ape",
	}
	config.EscapeOutputStrings = false
	// Act
	result, err := BuildLine(source, "T", 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "T|1|esc^ape", result)
	// Teardown
	teardown()
}
func TestBuildLine_EscapedChars_HL7(t *testing.T) {
	// Arrange
	source := SimpleRecord{
		First: "esc^ape",
	}
	configHL7.EscapeOutputStrings = true
	// Act
	result, err := BuildLine(source, "REC", 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "REC||esc\\S\\ape", result)
	// Teardown
	teardown()
}

// buildStringEscapeChars tests
func TestBuildStringEscapeChars_AllDelimiters_ASTM(t *testing.T) {
	// Arrange
	input := "esc|\\^&ape"
	// Act
	result := buildStringEscapeChars(input, config)
	// Assert
	assert.Equal(t, "esc&|&\\&^&&ape", result)
}
func TestBuildStringEscapeChars_Unicode_ASTM(t *testing.T) {
	// Arrange
	input := "^őáúäö|"
	// Act
	result := buildStringEscapeChars(input, config)
	// Assert
	assert.Equal(t, "&^őáúäö&|", result)
}
func TestBuildStringEscapeChars_AllDelimiters_HL7(t *testing.T) {
	// Arrange
	input := "esc\r|~^&\\ape"
	// Act
	result := buildStringEscapeChars(input, configHL7)
	// Assert
	assert.Equal(t, `esc\X0D\\F\\R\\S\\T\\E\ape`, result)
}
