package configs

import (
	"os"
	"testing"
)

func TestNewFlagConfig(t *testing.T) {
	// установим значения флагов для тестирования
	os.Args = append(os.Args, "-a=127.0.0.1:8888")
	os.Args = append(os.Args, "-b=http://example.com/")

	config := NewFlagConfig()

	address := config.Address()
	if address != "127.0.0.1:8888" {
		t.Errorf("Address was incorrect, got: %s, want: %s", address, "127.0.0.1:8888")
	}

	baseURL := config.BaseURL()
	if baseURL != "http://example.com/" {
		t.Errorf("BaseURL was incorrect, got: %s, want: %s", baseURL, "http://example.com/")
	}
}
