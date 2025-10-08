package functions

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/models"
	"github.com/blutspende/go-parser/parserconfig"
)

func ParseFieldAnnotation(input reflect.StructField, config *parserconfig.Configuration) (result models.FieldAnnotation, err error) {
	// Get the correct annotation key for the protocol
	key, err := getAnnotationKeyForProtocol(config)
	if err != nil {
		return models.FieldAnnotation{}, err
	}
	// Get the tag value and check if it is empty
	raw := input.Tag.Get(key)
	if raw == "" {
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingMissingAnnotation
	}
	// Parse the annotation string
	result, err = parseFieldAnnotationString(raw)
	if err != nil {
		return models.FieldAnnotation{}, err
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
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingIllegalComponentArray
	}
	if result.IsComponent && result.IsSubstructure {
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingIllegalComponentSubstructure
	}
	// All okay, return the result and no error
	return result, nil
}

func parseFieldAnnotationString(input string) (result models.FieldAnnotation, err error) {
	// Save the input into the raw
	result.Raw = input
	// Separate attributes and the field definition
	fieldDef, attributes := splitByFirst(input, ",")
	// Parse and save attributes
	result.Attributes, err = parseAttributes(attributes, []string{
		constants.AttributeRequired,
		constants.AttributeLongdate,
		constants.AttributeLength,
	})
	if err != nil {
		return models.FieldAnnotation{}, err
	}
	// Split field and component (if any) and parse them
	segments := strings.Split(fieldDef, ".")
	if len(segments) > 2 {
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingInvalidAnnotation
	}
	if len(segments) == 2 {
		result.IsComponent = true
		result.ComponentPos, err = strconv.Atoi(segments[1])
		if err != nil {
			return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingInvalidAnnotation
		}
	}
	result.FieldPos, err = strconv.Atoi(segments[0])
	if err != nil {
		return models.FieldAnnotation{}, errmsg.ErrAnnotationParsingInvalidAnnotation
	}
	// Return the result
	return result, nil
}

func ParseStructAnnotation(input reflect.StructField, config *parserconfig.Configuration) (result models.StructAnnotation, err error) {
	// Get the correct annotation key for the protocol
	key, err := getAnnotationKeyForProtocol(config)
	if err != nil {
		return models.StructAnnotation{}, err
	}
	// Get the tag value
	raw := input.Tag.Get(key)
	// Save the raw annotation
	result.Raw = raw
	// Determine if the struct is composite (no tag) or not
	result.IsComposite = raw == ""
	// Determine if the field is an array or not
	result.IsArray = input.Type.Kind() == reflect.Slice || input.Type.Kind() == reflect.Array
	// Composite has no tag so further parsing is not needed
	if result.IsComposite {
		return result, nil
	}
	// Separate attributes and the struct name, and save the name
	attributes := ""
	result.StructName, attributes = splitByFirst(raw, ",")
	// Parse and save attributes
	result.Attributes, err = parseAttributes(attributes, []string{
		constants.AttributeOptional,
		constants.AttributeSubname,
	})
	// Return the result
	return result, err
}

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

func parseAttributes(input string, valids []string) (result map[string]string, err error) {
	// Initialize the result map
	result = make(map[string]string)
	// Check for empty input (if empty still return a usable empty map)
	if input == "" {
		return result, nil
	}
	// Split the input string by commas
	attributes := strings.Split(input, ",")
	// Iterate over the attributes and parse them
	for _, attribute := range attributes {
		// Split each attribute by the colon
		attributeParts := strings.Split(attribute, ":")
		if len(attributeParts) > 2 {
			return nil, errmsg.ErrAnnotationParsingInvalidAttributeFormat
		}
		// Check if the attribute is valid
		if !isInList(attributeParts[0], valids) {
			return nil, errmsg.ErrAnnotationParsingInvalidAttribute
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

func splitByFirst(input string, delimiter string) (before string, after string) {
	index := strings.Index(input, delimiter) // Find the first occurrence of the comma
	if index == -1 {
		return input, "" // No comma, return whole string and empty second part
	}
	return input[:index], input[index+1:] // Split at the first comma
}
func isInList(target string, list []string) bool {
	set := make(map[string]struct{})
	for _, item := range list {
		set[item] = struct{}{}
	}
	_, exists := set[target]
	return exists
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
