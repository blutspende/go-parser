package e2e

import (
	"bytes"
	"testing"

	"github.com/blutspende/go-astm/v3/parserconfig"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// Configuration struct for tests
var config *parserconfig.Configuration

// Reset config to default values
func teardown() {
	config = parserconfig.NewDefaultConfiguration(parserconfig.ASTM)
	_ = parserconfig.InitConfig(config)
}

// Setup default config and run all tests
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}

// Encoding helper function
func helperEncode(charmap *charmap.Charmap, data []byte) []byte {
	e := charmap.NewEncoder()
	var b bytes.Buffer
	writer := transform.NewWriter(&b, e)
	writer.Write(data)
	resultdata := b.Bytes()
	writer.Close()
	return resultdata
}
