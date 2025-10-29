package parser

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/parserconfig"
)

func Marshal(sourceStruct interface{}, config *parserconfig.Configuration) (result [][]byte, err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return nil, err
	}
	// Build the lines from the source structure
	lines, err := functions.BuildStruct(sourceStruct, config)
	if err != nil {
		return nil, err
	}
	// Convert UTF8 string array to encoding
	result, err = encoding.ConvertArrayFromUtf8ToEncoding(lines, config.Encoding)
	if err != nil {
		return nil, err
	}
	// Return the result and no error if everything went well
	return result, nil
}
