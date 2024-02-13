package configs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFlagConfig(t *testing.T) {
	// Set flag values for testing
	os.Args = append(os.Args, "-a=127.0.0.1:8888")
	os.Args = append(os.Args, "-b=http://example.com/")

	config, _ := NewFlagConfig()

	address := config.GetAddress()
	assert.Equal(t, address, "127.0.0.1:8888", "GetAddress was incorrect")

	baseURL := config.GetBaseURL()
	assert.Equal(t, baseURL, "http://example.com/", "GetBaseURL was incorrect")
}
