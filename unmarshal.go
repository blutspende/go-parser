package parser

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/go-parser/functions"
	"github.com/blutspende/go-parser/pconfig"
)

func Unmarshal(messageData []byte, targetStruct interface{}, config *pconfig.Configuration) (err error) {
	// Init configuration
	err = pconfig.InitConfig(config)
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
	err = functions.ParseStruct(lines, targetStruct, config)
	if err != nil {
		return err
	}
	// Return nil if everything went well
	return nil
}
