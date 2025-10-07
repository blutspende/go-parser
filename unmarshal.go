package parser

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/parserconfig"
)

func Unmarshal(messageData []byte, targetStruct interface{}, config *parserconfig.Configuration) (err error) {
	// Init configuration
	err = parserconfig.InitConfig(config)
	if err != nil {
		return err
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUtf8(messageData, config.Encoding)
	if err != nil {
		return err
	}
	// Split the message data into lines
	lines, err := functions.SliceLines(utf8Data, config)
	if err != nil {
		return err
	}
	// Parse the lines into the target structure
	lineIndex := 0
	err = functions.ParseStruct(lines, targetStruct, &lineIndex, 1, 0, config)
	if err != nil {
		return err
	}
	// Return nil if everything went well
	return nil
}
