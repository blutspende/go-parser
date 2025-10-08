package functions

import (
	"reflect"
	"testing"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/stretchr/testify/assert"
)

// Field annotation tests
func TestParseAstmFieldAnnotationString_SingleValue(t *testing.T) {
	// Arrange
	input := "4"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmFieldAnnotationString_Componented(t *testing.T) {
	// Arrange
	input := "4.1"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4.1", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 1, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmFieldAnnotationString_Attributed(t *testing.T) {
	// Arrange
	input := "4,required"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4,required", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
}
func TestParseAstmFieldAnnotationString_AttributedValue(t *testing.T) {
	// Arrange
	input := "4,length:2"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4,length:2", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "2", result.Attributes[constants.AttributeLength])
}
func TestParseAstmFieldAnnotationString_Complex(t *testing.T) {
	// Arrange
	input := "3.2,length:4"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3.2,length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 2, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseAstmFieldAnnotationString_InvalidAttribute(t *testing.T) {
	// Arrange
	input := "4.1,something"
	// Act
	_, err := parseFieldAnnotationString(input)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAttribute.Error())
}
func TestParseAstmFieldAnnotationString_InvalidAttributeFormat(t *testing.T) {
	// Arrange
	input := "4.1,length:2:3"
	// Act
	_, err := parseFieldAnnotationString(input)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAttributeFormat.Error())
}
func TestParseAstmFieldAnnotationString_InvalidAnnotationTooManyParts(t *testing.T) {
	// Arrange
	input := "2.1.2"
	// Act
	_, err := parseFieldAnnotationString(input)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAnnotation.Error())
}
func TestParseAstmFieldAnnotationString_InvalidAnnotationTooManyPartsWithAttribute(t *testing.T) {
	// Arrange
	input := "4.1.3,required"
	// Act
	_, err := parseFieldAnnotationString(input)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAnnotation.Error())
}
func TestParseAstmFieldAnnotationString_InvalidNumber(t *testing.T) {
	// Arrange
	input := "4.f"
	// Act
	_, err := parseFieldAnnotationString(input)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAnnotation.Error())
}
func TestParseAstmFieldAnnotationString_MultipleAttributes(t *testing.T) {
	// Arrange
	input := "4,required,longdate"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4,required,longdate", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
	assert.Contains(t, result.Attributes, constants.AttributeLongdate)
}
func TestParseAstmFieldAnnotationString_ValuedMultipleAttributes(t *testing.T) {
	// Arrange
	input := "4,length:3,required"
	// Act
	result, err := parseFieldAnnotationString(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4,length:3,required", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "3", result.Attributes[constants.AttributeLength])
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
}
func TestParseAstmFieldAnnotation_AnnotatedStruct(t *testing.T) {
	// Arrange
	var input AnnotatedLine
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3.2,length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 2, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseAstmFieldAnnotation_AnnotatedArrayStruct(t *testing.T) {
	// Arrange
	var input AnnotatedArrayLine
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3,length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseAstmFieldAnnotation_Substructure(t *testing.T) {
	// Arrange
	var input SubstructuredLine
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, true, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmFieldAnnotation_SubstructureArray(t *testing.T) {
	// Arrange
	var input SubstructuredLine
	field, _ := reflect.TypeOf(input).FieldByName("Array")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "4", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, true, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmFieldAnnotation_IllegalComponentArray(t *testing.T) {
	// Arrange
	var input IllegalComponentArray
	field, _ := reflect.TypeOf(input).FieldByName("ComponentArray")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingIllegalComponentArray.Error())
}
func TestParseAstmFieldAnnotation_IllegalComponentSubstructure(t *testing.T) {
	// Arrange
	var input IllegalComponentSubstructure
	field, _ := reflect.TypeOf(input).FieldByName("ComponentSubstructure")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingIllegalComponentSubstructure.Error())
}
func TestParseAstmFieldAnnotation_TimeLine(t *testing.T) {
	// Arrange
	var input TimeLine
	field, _ := reflect.TypeOf(input).FieldByName("Time")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmFieldAnnotation_InvalidFieldAttribute(t *testing.T) {
	// Arrange
	var input InvalidFieldAttribute
	field, _ := reflect.TypeOf(input).FieldByName("First")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAttribute.Error())
}

// Struct annotation tests
func TestParseAstmStructAnnotation_SingleLineStruct(t *testing.T) {
	// Arrange
	var input SingleLineStruct
	field, _ := reflect.TypeOf(input).FieldByName("Line")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "L", result.Raw)
	assert.Equal(t, false, result.IsComposite)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "L", result.StructName)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmStructAnnotation_AnnotatedArrayStruct(t *testing.T) {
	// Arrange
	var input AnnotatedArrayStruct
	field, _ := reflect.TypeOf(input).FieldByName("Lines")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "L,optional", result.Raw)
	assert.Equal(t, false, result.IsComposite)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "L", result.StructName)
	assert.Contains(t, result.Attributes, constants.AttributeOptional)
}
func TestParseAstmStructAnnotation_CompositeStruct(t *testing.T) {
	// Arrange
	var input CompositeStruct
	field, _ := reflect.TypeOf(input).FieldByName("Composite")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "", result.Raw)
	assert.Equal(t, true, result.IsComposite)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "", result.StructName)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmStructAnnotation_CompositeArrayStruct(t *testing.T) {
	// Arrange
	var input CompositeArrayStruct
	field, _ := reflect.TypeOf(input).FieldByName("Composite")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "", result.Raw)
	assert.Equal(t, true, result.IsComposite)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "", result.StructName)
	assert.Empty(t, result.Attributes)
}
func TestParseAstmStructAnnotation_InvalidStructAttribute(t *testing.T) {
	// Arrange
	var input InvalidStructAttribute
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAttribute.Error())
}
func TestParseAstmStructAnnotation_TooManyStructNameAttributeValues(t *testing.T) {
	// Arrange
	var input TooManyStructNameAttributeValues
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidAttributeFormat.Error())
}
func TestParseAstmStructAnnotation_SubnameAttribute(t *testing.T) {
	// Arrange
	var input SubnameAttribute
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "R,subname:SUBNAME", result.Raw)
	assert.Equal(t, false, result.IsComposite)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "R", result.StructName)
	assert.Contains(t, result.Attributes, constants.AttributeSubname)
	assert.Equal(t, "SUBNAME", result.Attributes[constants.AttributeSubname])
}

// ProcessStructReflection tests
func TestProcessStructReflection_SimpleRecord(t *testing.T) {
	// Arrange
	input := ThreeFieldRecord{}
	// Act
	types, values, length, err := ProcessStructReflection(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 3, length)
	assert.Len(t, types, 3)
	assert.Len(t, values, 3)
	assert.Equal(t, "First", types[0].Name)
}
func TestProcessStructReflection_CompositeRecordStruct(t *testing.T) {
	// Arrange
	input := CompositeRecordStruct{}
	// Act
	types, values, length, err := ProcessStructReflection(input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, length)
	assert.Len(t, types, 2)
	assert.Len(t, values, 2)
	assert.Equal(t, "Record1", types[0].Name)
}
func TestProcessStructReflection_SimpleRecordPointer(t *testing.T) {
	// Arrange
	input := ThreeFieldRecord{}
	// Act
	_, _, length, err := ProcessStructReflection(&input)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 3, length)
}

// HL7 specific tests
func TestParseAstmFieldAnnotation_LineHL7(t *testing.T) {
	// Arrange
	var input LineHL7
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "3", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
}
func TestParseAstmFieldAnnotation_LineHL7_WrongProtocol(t *testing.T) {
	// Arrange
	var input LineHL7
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingMissingAnnotation.Error())
	assert.Equal(t, "", result.Raw)
}
func TestParseAstmFieldAnnotation_LineHL7_InvalidProtocol(t *testing.T) {
	// Arrange
	var input LineHL7
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	configHL7.Protocol = "X"
	// Act
	result, err := ParseFieldAnnotation(field, configHL7)
	// Assert
	assert.EqualError(t, err, errmsg.ErrAnnotationParsingInvalidProtocol.Error())
	assert.Equal(t, "", result.Raw)
	// Teardown
	teardown()
}
func TestParseAstmStructAnnotation_SingleLineStructHL7(t *testing.T) {
	// Arrange
	var input SingleLineStructHL7
	field, _ := reflect.TypeOf(input).FieldByName("Line")
	// Act
	result, err := ParseStructAnnotation(field, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "LIN", result.Raw)
	assert.Equal(t, false, result.IsComposite)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "LIN", result.StructName)
	assert.Empty(t, result.Attributes)
}
