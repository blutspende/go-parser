package functions

import (
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/parserconfig"
)

func getAnnotationKeyForProtocol(config *parserconfig.Configuration) (key string, err error) {
	switch config.Protocol {
	case parserconfig.ASTM:
		return "astm", nil
	case parserconfig.HL7:
		return "hl7", nil
	default:
		return "", errmsg.ErrAnnotationParsingInvalidProtocol
	}
}

func ParseStructAnnotation(input reflect.StructField, config *parserconfig.Configuration) (result models.StructAnnotation, err error) {
	// Get the correct annotation key for the protocol
	key, err := getAnnotationKeyForProtocol(config)
	if err != nil {
		return models.StructAnnotation{}, err
	}
	// Extract and save the raw value
	result.Raw = input.Tag.Get(key)
	if result.Raw == "" {
		return models.StructAnnotation{}, errmsg.ErrAnnotationParsingMissingAnnotation
	}
	// Parse the annotation
	elements, err := parseAnnotationElements(result.Raw, ";", "=", []string{
		constants.AnnotationElementGroup,
		constants.AnnotationElementTag,
		constants.AnnotationElementAttribute,
	})
	// Extract if the struct is a group
	_, result.IsGroup = elements[constants.AnnotationElementGroup]
	// Extract the tag if present
	var hasTag bool
	result.Tag, hasTag = elements[constants.AnnotationElementTag]
	// Check for illegal combinations
	if result.IsGroup == hasTag {
		return models.StructAnnotation{}, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingIllegal, "structure must be either a group or have a tag")
	}
	// Determine if the field is an array or not
	result.IsArray = input.Type.Kind() == reflect.Slice || input.Type.Kind() == reflect.Array
	// Extract attributes if any
	if attributes, hasAttributes := elements[constants.AnnotationElementAttribute]; hasAttributes {
		// Parse and save attributes
		result.Attributes, err = parseAnnotationElements(attributes, ",", ":", []string{
			constants.AttributeOptional,
			constants.AttributeSubname,
		})
		if err != nil {
			return models.StructAnnotation{}, err
		}
	}
	// Return the result with no error
	return result, nil
}

func ParseFieldAnnotation(input reflect.StructField, config *parserconfig.Configuration) (result models.FieldAnnotation, err error) {
	// Get the correct annotation key for the protocol
	key, err := getAnnotationKeyForProtocol(config)
	if err != nil {
		return models.FieldAnnotation{}, err
	}
	// Extract and save the raw value
	result.Raw = input.Tag.Get(key)
	if result.Raw == "" {
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingMissingAnnotation
	}
	// Parse the annotation
	elements, err := parseAnnotationElements(result.Raw, ";", "=", []string{
		constants.AnnotationElementPosition,
		constants.AnnotationElementAttribute,
	})
	// Extract and parse the position
	if posString, hasPos := elements[constants.AnnotationElementTag]; hasPos {
		// Prepare the error for any parsing issue
		errParse := fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingInvalidElement, posString)
		// Split field and component (if any)
		segments := strings.Split(posString, ".")
		if len(segments) > 2 {
			return models.FieldAnnotation{}, errParse
		}
		// Parse component position if present
		if len(segments) == 2 {
			result.IsComponent = true
			result.ComponentPos, err = strconv.Atoi(segments[1])
			if err != nil {
				return models.FieldAnnotation{}, errParse
			}
		}
		// Parse field position
		result.FieldPos, err = strconv.Atoi(segments[0])
		if err != nil {
			return models.FieldAnnotation{}, errParse
		}
	} else {
		return models.FieldAnnotation{}, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingIllegal, "field must have a position")
	}
	// Determine if the field is an array or not
	result.IsArray = input.Type.Kind() == reflect.Slice || input.Type.Kind() == reflect.Array
	// Determine if the field is a substructure or not (excluding the time.Time type)
	var checkType reflect.Type
	if result.IsArray {
		checkType = input.Type.Elem()
	} else {
		checkType = input.Type
	}
	result.IsSubstructure = checkType.Kind() == reflect.Struct && checkType != reflect.TypeOf(time.Time{})
	// Check illegal combinations
	if result.IsComponent && result.IsArray {
		return models.FieldAnnotation{}, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingIllegal, "field can not be component and array")
	}
	if result.IsComponent && result.IsSubstructure {
		return models.FieldAnnotation{}, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingIllegal, "field can not be component and substructure")
	}
	// Extract attributes if any
	if attributes, hasAttributes := elements[constants.AnnotationElementAttribute]; hasAttributes {
		// Parse and save attributes
		result.Attributes, err = parseAnnotationElements(attributes, ",", ":", []string{
			constants.AttributeRequired,
			constants.AttributeLongdate,
			constants.AttributeLength,
			constants.AttributeSequence,
		})
		if err != nil {
			return models.FieldAnnotation{}, err
		}
	}
	// Return the result with no error
	return result, nil
}

func parseAnnotationElements(input, elemSep, partSep string, valids []string) (result map[string]string, err error) {
	// Initialize the result map
	result = make(map[string]string)
	// Check for empty input (if empty still return a usable empty map)
	if input == "" {
		return result, nil
	}
	// Split the input string by commas
	attributes := strings.Split(input, elemSep)
	// Iterate over the attributes and parse them
	for _, attribute := range attributes {
		// Split each attribute by the colon
		attributeParts := strings.Split(attribute, partSep)
		if len(attributeParts) > 2 {
			return nil, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingInvalidElement, input)
		}
		// Check if the attribute is valid (empty valids means: anything goes)
		if len(valids) > 0 && !slices.Contains(valids, attributeParts[0]) {
			return nil, fmt.Errorf("%w: %s", errmsg.ErrAnnotationParsingInvalidElementKey, attributeParts[0])
		}
		// Save the attribute name and value (if present)
		if len(attributeParts) == 2 {
			result[attributeParts[0]] = attributeParts[1]
		} else {
			result[attributeParts[0]] = ""
		}
	}
	// Return the result map and no error
	return result, nil
}

func ProcessStructReflection(inputStruct interface{}) (outputTypes []reflect.StructField, outputValues []reflect.Value, length int, err error) {
	// Ensure the inputStruct is a pointer to a struct
	targetPtrValue := reflect.ValueOf(inputStruct)
	if targetPtrValue.Kind() != reflect.Ptr {
		// If inputStruct is not a pointer, take its address
		targetPtrValue = reflect.New(reflect.TypeOf(inputStruct))
		targetPtrValue.Elem().Set(reflect.ValueOf(inputStruct))
	}
	if targetPtrValue.Elem().Kind() != reflect.Struct {
		// inputStruct must be a pointer to a struct
		return nil, nil, 0, errmsg.ErrAnnotationParsingInvalidInputStruct
	}
	// Get the underlying struct
	targetValue := targetPtrValue.Elem()
	targetType := targetPtrValue.Type().Elem()
	// Allocate the results
	outputTypes = make([]reflect.StructField, targetValue.NumField())
	outputValues = make([]reflect.Value, targetType.NumField())
	length = targetType.NumField()
	// Iterate and save outputTypes and outputValues
	for i := 0; i < targetType.NumField(); i++ {
		outputTypes[i] = targetType.Field(i)
		outputValues[i] = targetValue.Field(i)
	}
	// Return the results
	return outputTypes, outputValues, length, nil
}
