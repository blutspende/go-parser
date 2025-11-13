package astm

import (
	"bytes"
	"testing"

	"github.com/blutspende/go-parser/pconfig"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// Configuration struct for tests
var config *pconfig.Configuration

// Reset config to default values
func teardown() {
	config = pconfig.NewDefaultConfiguration(pconfig.ASTM)
	_ = pconfig.InitConfig(config)
}

// Setup default config and run all tests
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}

// Encoding helper function
func helperEncode(charMap *charmap.Charmap, data []byte) []byte {
	e := charMap.NewEncoder()
	var b bytes.Buffer
	writer := transform.NewWriter(&b, e)
	_, _ = writer.Write(data)
	result := b.Bytes()
	_ = writer.Close()
	return result
}
