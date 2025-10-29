package functions

import (
	"testing"
	"time"

	"github.com/blutspende/go-parser/errmsg"
	"github.com/stretchr/testify/assert"
)

// Note: structures come from functions_test.go

// ParseLine tests - basics
func TestParseLine_SimpleRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|second|third"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "second", target.Second)
	assert.Equal(t, "third", target.Third)
}
func TestParseLine_UnorderedRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|second|third"
	target := UnorderedRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "second", target.Second)
	assert.Equal(t, "third", target.Third)
}
func TestParseLine_MultitypeRecord(t *testing.T) {
	// Arrange
	input := "T|1|string|3|3.14|3.14159265|20060102"
	target := MultitypeRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "string", target.String)
	assert.Equal(t, 3, target.Int)
	assert.Equal(t, float32(3.14), target.Float32)
	assert.Equal(t, 3.14159265, target.Float64)
	expectedShortTime := time.Date(2006, 1, 2, 0, 0, 0, 0, config.TimeLocation)
	assert.Equal(t, expectedShortTime, target.Time)
}
func TestParseLine_MultitypePointerRecord(t *testing.T) {
	// Arrange
	input := "T|1|string|3|3.14|3.14159265|20060102"
	target := MultitypePointerRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "string", *target.String)
	assert.Equal(t, 3, *target.Int)
	assert.Equal(t, float32(3.14), *target.Float32)
	assert.Equal(t, 3.14159265, *target.Float64)
	expectedShortTime := time.Date(2006, 1, 2, 0, 0, 0, 0, config.TimeLocation)
	assert.Equal(t, expectedShortTime, *target.Time)
}
func TestParseLine_ComponentedRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|second1^second2|third1^third2^third3"
	target := ComponentedRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "second1", target.SecondComp1)
	assert.Equal(t, "second2", target.SecondComp2)
	assert.Equal(t, "third1", target.ThirdComp1)
	assert.Equal(t, "third2", target.ThirdComp2)
	assert.Equal(t, "third3", target.ThirdComp3)
}
func TestParseLine_EnumRecord(t *testing.T) {
	// Arrange
	input := "T|1|enum"
	target := EnumRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, EnumString("enum"), target.Enum)
}

// ParseLine tests - header
func TestParseLine_Header_ASTM(t *testing.T) {
	// Arrange
	input := `H|\^&|first`
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("H"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
}
func TestParseLine_HeaderDelimiterChange_ASTM(t *testing.T) {
	// Arrange
	input := "H/!*%/first/second1!second2/third1*third2"
	target := HeaderDelimiterChange{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("H"), 0, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Len(t, target.Array, 2)
	assert.Equal(t, "second1", target.Array[0])
	assert.Equal(t, "second2", target.Array[1])
	assert.Equal(t, "third1", target.Comp1)
	assert.Equal(t, "third2", target.Comp2)
	// Teardown
	teardown()
}
func TestParseLine_HeaderDelimiterParsing_ASTM(t *testing.T) {
	// Arrange
	input := "H/!*%"
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("H"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "/", config.Delimiters.Field)
	assert.Equal(t, "!", config.Delimiters.Repeat)
	assert.Equal(t, "*", config.Delimiters.Component)
	assert.Equal(t, "%", config.Delimiters.Escape)
	// Teardown
	teardown()
}
func TestParseLine_HeaderTooShort_ASTM(t *testing.T) {
	// Arrange
	input := `H|\^`
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("H"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingHeaderTooShort)
}
func TestParseLine_Header_HL7(t *testing.T) {
	// Arrange
	input := `MSH|^~\&|first`
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("MSH"), 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
}
func TestParseLine_HeaderDelimiterParsing_HL7(t *testing.T) {
	// Arrange
	input := "MSH/!*%?"
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("MSH"), 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "/", configHL7.Delimiters.Field)
	assert.Equal(t, "!", configHL7.Delimiters.Component)
	assert.Equal(t, "*", configHL7.Delimiters.Repeat)
	assert.Equal(t, "%", configHL7.Delimiters.Escape)
	assert.Equal(t, "?", configHL7.Delimiters.SubComponent)
	// Teardown
	teardown()
}
func TestParseLine_HeaderTooShort_HL7(t *testing.T) {
	// Arrange
	input := `MSH|^~\`
	target := HeaderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("MSH"), 1, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingHeaderTooShort)
}

// ParseLine tests - missing data and errors
func TestParseLine_MissingData(t *testing.T) {
	// Arrange
	input := "T|1|first||third"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "", target.Second)
	assert.Equal(t, "third", target.Third)
}
func TestParseLine_MissingComponent(t *testing.T) {
	// Arrange
	input := "T|1|first^second"
	target := RequiredComponentRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "second", target.Second)
	assert.Equal(t, "", target.Third)
}
func TestParseLine_MissingRequiredComponent(t *testing.T) {
	// Arrange
	input := "T|1|first"
	target := RequiredComponentRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingInputComponentsMissing)
}
func TestParseLine_MissingDataAtTheEnd(t *testing.T) {
	// Arrange
	input := "T|1|first"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "", target.Second)
	assert.Equal(t, "", target.Third)
}

func TestParseLine_RecordTypeNameMismatch(t *testing.T) {
	// Arrange
	input := "W|1|first|second|third"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingLineTagMismatch)
}
func TestParseLine_EmptyInput(t *testing.T) {
	// Arrange
	input := ""
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingEmptyInput)
}
func TestParseLine_MandatoryFieldsMissing(t *testing.T) {
	// Arrange
	input := "T"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingMandatoryInputFieldsMissing)
}
func TestParseLine_MissingRequiredField(t *testing.T) {
	// Arrange
	input := "T|1|first||third"
	target := RequiredFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingRequiredInputFieldMissing)
}
func TestParseLine_NotEnoughInputFields(t *testing.T) {
	// Arrange
	input := "T|1|first"
	target := RequiredFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingRequiredInputFieldMissing)
}
func TestParseLine_SequenceNumberMismatch(t *testing.T) {
	// Arrange
	input := "T|2|first|second|third"
	target := ThreeFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingSequenceNumberMismatch)
}
func TestParseLine_SequenceNumberMismatchWithoutEnforcing(t *testing.T) {
	// Arrange
	input := "T|2|first|second|third"
	target := ThreeFieldRecord{}
	config.EnforceSequenceNumberCheck = false
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	// Teardown
	teardown()
}
func TestParseLine_ReservedFieldRecord(t *testing.T) {
	// Arrange
	input := "T|1"
	target := ReservedFieldRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingReservedFieldPosReference)
}
func TestParseLine_MissingAnnotation(t *testing.T) {
	// Arrange
	input := "T|1|field3|field4"
	target := MissingAnnotationRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "field3", target.Field3)
	assert.Equal(t, "field4", target.Field4)
}

// ParseLine tests - arrays, substructures and sub-substructures
func TestParseLine_ArrayRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|second1\\second2\\second3"
	target := ArrayRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Len(t, target.Array, 3)
	assert.Equal(t, "second1", target.Array[0])
	assert.Equal(t, "second2", target.Array[1])
	assert.Equal(t, "second3", target.Array[2])
}
func TestParseLine_SubstructureRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|firstComponent^secondComponent^thirdComponent|third"
	target := SubstructureRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "firstComponent", target.Second.FirstComponent)
	assert.Equal(t, "secondComponent", target.Second.SecondComponent)
	assert.Equal(t, "thirdComponent", target.Second.ThirdComponent)
	assert.Equal(t, "third", target.Third)
}
func TestParseLine_SubstructureArrayRecord(t *testing.T) {
	// Arrange
	input := "T|1|first|r1c1^r1c2^r1c3\\r2c1^r2c2^r2c3|third"
	target := SubstructureArrayRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Len(t, target.Second, 2)
	assert.Equal(t, "r1c1", target.Second[0].FirstComponent)
	assert.Equal(t, "r1c2", target.Second[0].SecondComponent)
	assert.Equal(t, "r1c3", target.Second[0].ThirdComponent)
	assert.Equal(t, "r2c1", target.Second[1].FirstComponent)
	assert.Equal(t, "r2c2", target.Second[1].SecondComponent)
	assert.Equal(t, "r2c3", target.Second[1].ThirdComponent)
	assert.Equal(t, "third", target.Third)
}
func TestParseLine_WrongComponentOrder(t *testing.T) {
	// Arrange
	input := "T|1|first|comp1^comp2^comp3"
	target := WrongComponentOrderRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "first", target.First)
	assert.Equal(t, "comp1", target.Comp1)
	assert.Equal(t, "comp2", target.Comp2)
	assert.Equal(t, "comp3", target.Comp3)
}
func TestParseLine_WrongComponentPlacement(t *testing.T) {
	// Arrange
	input := "T|1|field1|comp1^comp2|field2"
	target := WrongComponentPlacementRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "field1", target.Field1)
	assert.Equal(t, "comp1", target.Comp1)
	assert.Equal(t, "field2", target.Field2)
	assert.Equal(t, "comp2", target.Comp2)
}
func TestParseLine_SubSub_ASTM(t *testing.T) {
	// Arrange
	input := "T|1|field1|comp1^sub1?sub2?sub3^comp3|field3"
	target := SubSubRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingMaximumRecursionDepthExceeded)
}
func TestParseLine_SubSub_HL7(t *testing.T) {
	// Arrange
	input := "REC|1|field1|comp1^sub1&sub2&sub3^comp3|field3"
	target := SubSubRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("REC"), 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "field1", target.First)
	assert.Equal(t, "comp1", target.Second.First)
	assert.Equal(t, "sub1", target.Second.Second.First)
	assert.Equal(t, "sub2", target.Second.Second.Second)
	assert.Equal(t, "sub3", target.Second.Second.Third)
	assert.Equal(t, "comp3", target.Second.Third)
	assert.Equal(t, "field3", target.Third)
}

// ParseLine tests - time and date
func TestParseLine_TimeRecord(t *testing.T) {
	// Arrange
	input := "T|1|20060306164429"
	target := TimeRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 06, 15, 44, 29, 0, time.UTC)
	assert.Equal(t, expected, target.Time)
}
func TestParseLine_DateRecord(t *testing.T) {
	// Arrange
	input := "T|1|20060306"
	target := DateRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 06, 0, 0, 0, 0, config.TimeLocation)
	assert.Equal(t, expected, target.Date)
}
func TestParseLine_TimeRecord_NotKeep(t *testing.T) {
	// Arrange
	input := "T|1|20060306164429"
	target := TimeRecord{}
	config.KeepShortDateTimeZone = false
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 06, 15, 44, 29, 0, time.UTC)
	assert.Equal(t, expected, target.Time)
	// Teardown
	teardown()
}
func TestParseLine_DateRecord_NotKeep(t *testing.T) {
	// Arrange
	input := "T|1|20060304"
	target := DateRecord{}
	config.KeepShortDateTimeZone = false
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 03, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, expected, target.Date)
	// Teardown
	teardown()
}
func TestParseLine_TimeRecord_DateInput(t *testing.T) {
	// Arrange
	input := "T|1|20060306"
	target := TimeRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 06, 0, 0, 0, 0, config.TimeLocation)
	assert.Equal(t, expected, target.Time)
}
func TestParseLine_DateRecord_TimeInput(t *testing.T) {
	// Arrange
	input := "T|1|20060306164429"
	target := DateRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("T"), 1, config)
	// Assert
	assert.Nil(t, err)
	expected := time.Date(2006, 03, 06, 15, 44, 29, 0, time.UTC)
	assert.Equal(t, expected, target.Date)
}

// ParseLine tests - escape, reserved and sequence
func TestParseLine_Escape_ASTM(t *testing.T) {
	// Arrange
	input := `R|1|esc&|ape`
	target := SimpleRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("R"), 1, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "esc|ape", target.Field)
}
func TestParseLine_Escape_HL7(t *testing.T) {
	// Arrange
	input := `REC||esc\F\ape`
	target := SimpleRecord{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("REC"), 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "esc|ape", target.Field)
}
func TestParseLine_ReservedFieldRecordHl7(t *testing.T) {
	// Arrange
	input := "REC|two|three"
	target := ReservedFieldRecordHL7{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("REC"), 1, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingReservedFieldPosReference)
}
func TestParseLine_NotReservedFieldRecordHl7(t *testing.T) {
	// Arrange
	input := "REC|two|three"
	target := NotReservedFieldRecordHL7{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("REC"), 1, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "two", target.Two)
	assert.Equal(t, "three", target.Three)
}
func TestParseLine_SequenceHl7(t *testing.T) {
	// Arrange
	input := "REC||2|data"
	target := SequenceHl7{}
	// Act
	err := ParseLine(input, &target, createStructAnnotation("REC"), 2, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, target.Sequence)
	assert.Equal(t, "data", target.Data)
}

// splitStringWithEscape tests
func TestSplitStringWithEscape_NoEscape(t *testing.T) {
	// Arrange
	input := "no&|split"
	// Act
	result := splitStringWithEscape(input, config.Delimiters.Field, config)
	// Assert
	assert.Len(t, result, 1)
	assert.Equal(t, "no&|split", result[0])
}
func TestSplitStringWithEscape_Mixed(t *testing.T) {
	// Arrange
	input := "no&^split^second&&^third"
	// Act
	result := splitStringWithEscape(input, config.Delimiters.Component, config)
	// Assert
	assert.Len(t, result, 3)
	assert.Equal(t, "no&^split", result[0])
	assert.Equal(t, "second&&", result[1])
	assert.Equal(t, "third", result[2])
}
func TestSplitStringWithEscape_EmptyFields(t *testing.T) {
	// Arrange
	input := "first||third"
	// Act
	result := splitStringWithEscape(input, config.Delimiters.Field, config)
	// Assert
	assert.Len(t, result, 3)
	assert.Equal(t, "first", result[0])
	assert.Equal(t, "", result[1])
	assert.Equal(t, "third", result[2])
}
func TestSplitStringWithEscape_EmptyInput(t *testing.T) {
	// Arrange
	input := ""
	// Act
	result := splitStringWithEscape(input, config.Delimiters.Field, config)
	// Assert
	assert.Len(t, result, 0)
}
func TestSplitStringWithEscape_Unicode(t *testing.T) {
	// Arrange
	input := "first|őáúäö&||third"
	// Act
	result := splitStringWithEscape(input, config.Delimiters.Field, config)
	// Assert
	assert.Len(t, result, 3)
	assert.Equal(t, "first", result[0])
	assert.Equal(t, "őáúäö&|", result[1])
	assert.Equal(t, "third", result[2])
}

// filterStringEscapeChars tests
func TestFilterEscapeChars_Delimiters(t *testing.T) {
	// Arrange
	input := "escaped&| and&^ and&&"
	// Act
	result := filterStringEscapeChars(input, config.Delimiters.Escape)
	// Assert
	assert.Equal(t, "escaped| and^ and&", result)
}
func TestFilterEscapeChars_Multiple(t *testing.T) {
	// Arrange
	input := "esc&&&|ape"
	// Act
	result := filterStringEscapeChars(input, config.Delimiters.Escape)
	// Assert
	assert.Equal(t, "esc&|ape", result)
}
func TestFilterEscapeChars_AtTheEnd(t *testing.T) {
	// Arrange
	// Note: this should be invalid, but for simplicity we allow it by escaping the nothing
	input := "escape&"
	// Act
	result := filterStringEscapeChars(input, config.Delimiters.Escape)
	// Assert
	assert.Equal(t, "escape", result)
}
func TestFilterEscapeChars_Empty(t *testing.T) {
	// Arrange
	input := ""
	// Act
	result := filterStringEscapeChars(input, config.Delimiters.Escape)
	// Assert
	assert.Equal(t, "", result)
}
func TestFilterEscapeChars_Unicode(t *testing.T) {
	// Arrange
	input := "őáúäö&|"
	// Act
	result := filterStringEscapeChars(input, config.Delimiters.Escape)
	// Assert
	assert.Equal(t, "őáúäö|", result)
}

// replaceHL7Escapes tests
func TestReplaceHL7Escapes(t *testing.T) {
	// Arrange
	input := `abc\F\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "abc|def", result)
}
func TestReplaceHL7Escapes_AtTheEnd(t *testing.T) {
	// Arrange
	input := `abcd\S\`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "abcd^", result)
}
func TestReplaceHL7Escapes_Unicode(t *testing.T) {
	// Arrange
	input := `őáú\S\äö`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "őáú^äö", result)
}
func TestReplaceHL7Escapes_MissingTerminator(t *testing.T) {
	// Arrange
	input := `abc\Fdef`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingUnterminatedEscapeSequence)
	assert.Equal(t, "", result)
}
func TestReplaceHL7Escapes_UnknownSequence(t *testing.T) {
	// Arrange
	input := `abc\U\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingUnknownEscapeSequence)
	assert.Equal(t, "", result)
}
func TestReplaceHL7Escapes_HexChar(t *testing.T) {
	// Arrange
	input := `abc\X0A\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "abc\ndef", result)
}
func TestReplaceHL7Escapes_InvalidHex(t *testing.T) {
	// Arrange
	input := `abc\X0A3\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingUnknownEscapeSequence)
	assert.Equal(t, "", result)
}
func TestReplaceHL7Escapes_InvalidLowercaseHex(t *testing.T) {
	// Arrange
	input := `abc\X0a\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.ErrorIs(t, err, errmsg.ErrLineParsingUnknownEscapeSequence)
	assert.Equal(t, "", result)
}
func TestReplaceHL7Escapes_Br(t *testing.T) {
	// Arrange
	input := `abc\.br\def`
	// Act
	result, err := replaceHL7Escapes(input, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "abc\rdef", result)
}
