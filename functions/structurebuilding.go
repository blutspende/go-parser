package functions

import (
	"github.com/blutspende/go-parser/constants"
	"github.com/blutspende/go-parser/errmsg"
	"github.com/blutspende/go-parser/parserconfig"
)

func BuildStruct(sourceStruct interface{}, sequenceNumber int, depth int, config *parserconfig.Configuration) (result []string, err error) {
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
			return nil, err
		}
		// Save the source value pointer
		sourceValue := sourceValues[i].Addr().Interface()

		// Source is an array it is iterated
		if sourceStructAnnotation.IsArray {
			for j := 0; j < sourceValues[i].Len(); j++ {
				if sourceStructAnnotation.IsGroup {
					// Composite source: recursively build the composite structure
					subResult, err := BuildStruct(sourceValues[i].Index(j).Addr().Interface(), j+1, depth+1, config)
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
				subResult, err := BuildStruct(sourceValue, sequenceNumber, depth+1, config)
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

	// Return the result and no error if everything went well
	return result, nil
}
