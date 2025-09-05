package config

import (
	"os"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				LogLevel:  "info",
				LogFormat: "json",
				Debug:     false,
			},
			wantErr: false,
		},
		{
			name: "invalid log level",
			config: &Config{
				LogLevel:  "invalid",
				LogFormat: "json",
				Debug:     false,
			},
			wantErr: true,
		},
		{
			name: "invalid log format",
			config: &Config{
				LogLevel:  "info",
				LogFormat: "invalid",
				Debug:     false,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	// Test with environment variables
	os.Setenv("DATA_DIFF_LOG_LEVEL", "debug")
	os.Setenv("DATA_DIFF_LOG_FORMAT", "text")
	defer func() {
		os.Unsetenv("DATA_DIFF_LOG_LEVEL")
		os.Unsetenv("DATA_DIFF_LOG_FORMAT")
	}()

	config, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if config.LogLevel != "debug" {
		t.Errorf("Load() LogLevel = %v, expected debug", config.LogLevel)
	}

	if config.LogFormat != "text" {
		t.Errorf("Load() LogFormat = %v, expected text", config.LogFormat)
	}
}

