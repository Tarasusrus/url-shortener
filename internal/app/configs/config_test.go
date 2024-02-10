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

	address := config.GetAddress()
	if address != "127.0.0.1:8888" {
		t.Errorf("GetAddress was incorrect, got: %s, want: %s", address, "127.0.0.1:8888")
	}

	baseURL := config.GetBaseURL()
	if baseURL != "http://example.com/" {
		t.Errorf("GetBaseURL was incorrect, got: %s, want: %s", baseURL, "http://example.com/")
	}
}
