package functions

import (
	"reflect"
	"testing"
	"time"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errdef"
	"github.com/stretchr/testify/assert"
)

// ParseFieldAnnotation tests
func TestParseFieldAnnotation_Position(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `astm:"POS=4"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseFieldAnnotation_Component(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `astm:"POS=4.1"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4.1", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 1, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseFieldAnnotation_Attribute(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `astm:"POS=4;ATR=required"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4;ATR=required", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
}
func TestParseFieldAnnotation_AttributeValue(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=4;ATR=length:2"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4;ATR=length:2", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "2", result.Attributes[constants.AttributeLength])
}
func TestParseFieldAnnotation_Complex(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=3.2;ATR=length:4"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3.2;ATR=length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 2, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseFieldAnnotation_InvalidAttribute(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `astm:"POS=4.1;ATR=something"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseFieldAnnotation_InvalidAttributeFormat(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=4.1;ATR=length:2:3"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
}
func TestParseFieldAnnotation_TooManyPositionParts(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=2.1.2"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
}
func TestParseFieldAnnotation_TooManyPositionPartsWithAttribute(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=2.1.2;ATR=required"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
}
func TestParseFieldAnnotation_InvalidNumber(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=4.f"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
}
func TestParseFieldAnnotation_MultipleAttributes(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=4;ATR=required,date"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4;ATR=required,date", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
	assert.Contains(t, result.Attributes, constants.AttributeDate)
}
func TestParseFieldAnnotation_ValuedMultipleAttributes(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `astm:"POS=4;ATR=length:3,required"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4;ATR=length:3,required", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, 0, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "3", result.Attributes[constants.AttributeLength])
	assert.Contains(t, result.Attributes, constants.AttributeRequired)
}
func TestParseFieldAnnotation_AnnotatedStruct(t *testing.T) {
	// Arrange
	var input ComplexRecord
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3.2;ATR=length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 2, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseFieldAnnotation_Array(t *testing.T) {
	// Arrange
	type inputType struct {
		Field []float32 `astm:"POS=3;ATR=length:4"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3;ATR=length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseFieldAnnotation_Substructure(t *testing.T) {
	// Arrange
	type inputType struct {
		Field Substructure `astm:"POS=3"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, true, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseFieldAnnotation_SubstructureArray(t *testing.T) {
	// Arrange
	type inputType struct {
		Field []Substructure `astm:"POS=4"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=4", result.Raw)
	assert.Equal(t, 4, result.FieldPos)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, true, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseFieldAnnotation_IllegalComponentArray(t *testing.T) {
	// Arrange
	type inputType struct {
		Field []string `astm:"POS=3.1"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingIllegal)
}
func TestParseFieldAnnotation_IllegalComponentSubstructure(t *testing.T) {
	// Arrange
	type inputType struct {
		Field []Substructure `astm:"POS=3.1"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingIllegal)
}
func TestParseFieldAnnotation_Time(t *testing.T) {
	// Arrange
	type inputType struct {
		Field time.Time `astm:"POS=3"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, false, result.IsComponent)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Empty(t, result.Attributes)
}
func TestParseFieldAnnotation_ComplexHL7(t *testing.T) {
	// Arrange
	type inputType struct {
		Field float32 `hl7:"POS=3.2;ATR=length:4"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	result, err := ParseFieldAnnotation(field, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "POS=3.2;ATR=length:4", result.Raw)
	assert.Equal(t, 3, result.FieldPos)
	assert.Equal(t, true, result.IsComponent)
	assert.Equal(t, 2, result.ComponentPos)
	assert.Equal(t, false, result.IsSubstructure)
	assert.Contains(t, result.Attributes, constants.AttributeLength)
	assert.Equal(t, "4", result.Attributes[constants.AttributeLength])
}
func TestParseFieldAnnotation_InvalidProtocol(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `astm:"POS=3"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	config.Protocol = "X"
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidProtocol)
	// Teardown
	teardown()
}
func TestParseFieldAnnotation_WrongProtocol(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string `hl7:"POS=3"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingWrongProtocol)
}
func TestParseFieldAnnotation_MissingAnnotation(t *testing.T) {
	// Arrange
	type inputType struct {
		Field string
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Field")
	// Act
	_, err := ParseFieldAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingMissingAnnotation)
}

// ParseStructAnnotation tests
func TestParseStructAnnotation_SingleRecord(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=R", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "R", result.Tag)
	assert.Empty(t, result.Attributes)
}
func TestParseStructAnnotation_Array(t *testing.T) {
	// Arrange
	type inputType struct {
		RecordArray []Record `astm:"TAG=R"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("RecordArray")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=R", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "R", result.Tag)
	assert.Empty(t, result.Attributes)
}
func TestParseStructAnnotation_AnnotatedArray(t *testing.T) {
	// Arrange
	type inputType struct {
		RecordArray []Record `astm:"TAG=R;ATR=optional"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("RecordArray")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=R;ATR=optional", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "R", result.Tag)
	assert.Contains(t, result.Attributes, constants.AttributeOptional)
}
func TestParseStructAnnotation_FreeTagOrder(t *testing.T) {
	// Arrange
	type inputType struct {
		RecordArray []Record `astm:"ATR=optional;TAG=R"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("RecordArray")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "ATR=optional;TAG=R", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "R", result.Tag)
	assert.Contains(t, result.Attributes, constants.AttributeOptional)
}
func TestParseStructAnnotation_Group(t *testing.T) {
	// Arrange
	type inputType struct {
		Group Group `astm:"GROUP"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Group")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "GROUP", result.Raw)
	assert.Equal(t, true, result.IsGroup)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "", result.Tag)
	assert.Empty(t, result.Attributes)
}
func TestParseStructAnnotation_GroupArray(t *testing.T) {
	// Arrange
	type inputType struct {
		Group []Group `astm:"GROUP"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Group")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "GROUP", result.Raw)
	assert.Equal(t, true, result.IsGroup)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "", result.Tag)
	assert.Empty(t, result.Attributes)
}
func TestParseStructAnnotation_InvalidKey(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R;INV=value"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseStructAnnotation_EmptyKey(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R;"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseStructAnnotation_InvalidAttribute(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R;ATR=invalid"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseStructAnnotation_IllegalCombination(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"GROUP;TAG=R"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingIllegal)
}
func TestParseStructAnnotation_MissingTag(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"ATR=optional"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingIllegal)
}
func TestParseStructAnnotation_AttributeValue(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R;ATR=subname:S"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	result, err := ParseStructAnnotation(field, config)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=R;ATR=subname:S", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "R", result.Tag)
	assert.Contains(t, result.Attributes, constants.AttributeSubname)
	assert.Equal(t, "S", result.Attributes[constants.AttributeSubname])
}
func TestParseStructAnnotation_TooManyAttributeValues(t *testing.T) {
	// Arrange
	type inputType struct {
		Record Record `astm:"TAG=R;ATR=subname:ONE:TWO"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	_, err := ParseStructAnnotation(field, config)
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
}
func TestParseStructAnnotation_SingleRecordHL7(t *testing.T) {
	// Arrange
	var input Group
	field, _ := reflect.TypeOf(input).FieldByName("Record")
	// Act
	result, err := ParseStructAnnotation(field, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=REC", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, false, result.IsArray)
	assert.Equal(t, "REC", result.Tag)
	assert.Empty(t, result.Attributes)
}
func TestParseStructAnnotation_ComplexHL7(t *testing.T) {
	// Arrange
	type inputType struct {
		RecordArray []Record `hl7:"TAG=REC;ATR=optional,subname:SUB"`
	}
	var input inputType
	field, _ := reflect.TypeOf(input).FieldByName("RecordArray")
	// Act
	result, err := ParseStructAnnotation(field, configHL7)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "TAG=REC;ATR=optional,subname:SUB", result.Raw)
	assert.Equal(t, false, result.IsGroup)
	assert.Equal(t, true, result.IsArray)
	assert.Equal(t, "REC", result.Tag)
	assert.Contains(t, result.Attributes, constants.AttributeOptional)
	assert.Contains(t, result.Attributes, constants.AttributeSubname)
	assert.Equal(t, "SUB", result.Attributes[constants.AttributeSubname])
}

// parseAnnotationElements tests
func TestParseAnnotationElements_Multiple(t *testing.T) {
	// Arrange
	input := "one-data1/two/three-data3"
	// Act
	result, err := parseAnnotationElements(input, "/", "-", []string{})
	// Assert
	assert.Nil(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, "data1", result["one"])
	assert.Equal(t, "", result["two"])
	assert.Equal(t, "data3", result["three"])
}
func TestParseAnnotationElements_MessageStruct(t *testing.T) {
	// Arrange
	input := "GROUP;ATR=optional"
	// Act
	result, err := parseAnnotationElements(input, ";", "=", []string{
		constants.AnnotationElementGroup,
		constants.AnnotationElementTag,
		constants.AnnotationElementAttribute,
	})
	// Assert
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "", result[constants.AnnotationElementGroup])
	assert.Equal(t, "optional", result[constants.AnnotationElementAttribute])
}
func TestParseAnnotationElements_Attributes(t *testing.T) {
	// Arrange
	input := "required,length:3"
	// Act
	result, err := parseAnnotationElements(input, ",", ":", []string{
		constants.AttributeRequired,
		constants.AttributeDate,
		constants.AttributeLength,
	})
	// Assert
	assert.Nil(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "", result[constants.AttributeRequired])
	assert.Equal(t, "3", result[constants.AttributeLength])
}
func TestParseAnnotationElements_InvalidKey(t *testing.T) {
	// Arrange
	input := "GROUP;ATR=optional"
	// Act
	_, err := parseAnnotationElements(input, ";", "=", []string{
		constants.AnnotationElementGroup,
	})
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseAnnotationElements_EmptyKey(t *testing.T) {
	// Arrange
	input := "ATR=optional;"
	// Act
	_, err := parseAnnotationElements(input, ";", "=", []string{
		constants.AnnotationElementAttribute,
	})
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElementKey)
}
func TestParseAnnotationElements_InvalidFormat(t *testing.T) {
	// Arrange
	input := "ATR=optional=invalid"
	// Act
	_, err := parseAnnotationElements(input, ";", "=", []string{
		constants.AnnotationElementAttribute,
	})
	// Assert
	assert.ErrorIs(t, err, errdef.ErrAnnotationParsingInvalidElement)
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
