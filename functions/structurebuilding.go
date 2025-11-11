package functions

import (
	"errors"
	"regexp"

	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/pconfig"
)

func BuildStruct(sourceStruct interface{}, config *pconfig.Configuration) (result []string, err error) {
	return buildStructRecursive(sourceStruct, 1, 0, config)
}

func buildStructRecursive(sourceStruct interface{}, sequenceNumber int, depth int, config *pconfig.Configuration) (result []string, err error) {
	// Check for maximum depth
	if depth >= constants.MaxDepth {
		return nil, errmsg.ErrStructureParsingMaxDepthReached
	}

	// Process the source structure
	sourceTypes, sourceValues, _, err := ProcessStructReflection(sourceStruct)
	if err != nil {
		return nil, err
	}

	// Iterate over the inputFields of the sourceStruct struct
	for i, sourceType := range sourceTypes {
		// Parse the sourceStruct field sourceFieldAnnotation
		sourceStructAnnotation, err := ParseStructAnnotation(sourceType, config)
		if err != nil {
			if errors.Is(err, errmsg.ErrAnnotationParsingMissingAnnotation) {
				// If the annotation is missing, skip this field
				continue
			} else {
				return nil, err
			}
		}
		// Save the source value pointer
		sourceValue := sourceValues[i].Addr().Interface()

		// Source is an array it is iterated
		if sourceStructAnnotation.IsArray {
			for j := 0; j < sourceValues[i].Len(); j++ {
				if sourceStructAnnotation.IsGroup {
					// Composite source: recursively build the composite structure
					subResult, err := buildStructRecursive(sourceValues[i].Index(j).Addr().Interface(), j+1, depth+1, config)
					if err != nil {
						return nil, err
					}
					result = append(result, subResult...)
				} else {
					// Non-composite source: build the single line
					lineResult, err := BuildLine(sourceValues[i].Index(j).Addr().Interface(), sourceStructAnnotation.Tag, j+1, config)
					if err != nil {
						return nil, err
					}
					result = append(result, lineResult)
				}
			}
		} else {
			// Source is a single element
			if sourceStructAnnotation.IsGroup {
				// Composite source: recursively build the composite structure
				subResult, err := buildStructRecursive(sourceValue, sequenceNumber, depth+1, config)
				if err != nil {
					return nil, err
				}
				result = append(result, subResult...)
			} else {
				// Only the first element is inheriting the sequence number
				seqNum := 1
				if i == 0 {
					seqNum = sequenceNumber
				}
				// Non-composite source: build the single line
				lineResult, err := BuildLine(sourceValue, sourceStructAnnotation.Tag, seqNum, config)
				if err != nil {
					return nil, err
				}
				result = append(result, lineResult)
			}
		}
	}

	// Filter empty lines if configured so (only in HL7)
	if config.DropEmptyOutputRecords && config.Protocol == pconfig.HL7 {
		return filterEmptyLines(result, config), nil
	} else {
		// Or return the raw result with all the lines
		return result, nil
	}
}

func filterEmptyLines(lines []string, config *pconfig.Configuration) (filteredLines []string) {
	fieldSep := config.Delimiters.Field
	if fieldSep == `\` {
		fieldSep = `\\`
	}
	compSep := config.Delimiters.Component
	emptyLine := regexp.MustCompile(`^[A-Z]([A-Z1-9]{2})?([` + fieldSep + `][1-9][0-9]?)?[` + fieldSep + compSep + `]*$`)
	filteredLines = make([]string, 0, len(lines))
	for _, line := range lines {
		if !emptyLine.MatchString(line) {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}
