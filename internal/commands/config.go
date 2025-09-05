package commands

import (
	"fmt"
	"os"

	"github.com/renepersau/data-diff/internal/config"
	"github.com/spf13/cobra"
)

// NewConfigCommand creates the config command
func NewConfigCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration management",
		Long:  "Manage configuration settings for data-diff",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			showConfig(cfg)
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Initialize configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
	})

	return cmd
}

func showConfig(cfg *config.Config) {
	fmt.Println("Current configuration:")
	fmt.Printf("  Log Level: %s\n", cfg.LogLevel)
	fmt.Printf("  Log Format: %s\n", cfg.LogFormat)
	fmt.Printf("  Debug: %t\n", cfg.Debug)
	fmt.Printf("  Input File: %s\n", cfg.InputFile)
	fmt.Printf("  Output File: %s\n", cfg.OutputFile)
	fmt.Printf("  Format: %s\n", cfg.Format)

	if config.IsConfigFileUsed() {
		fmt.Printf("  Config File: %s\n", config.GetConfigFile())
	} else {
		fmt.Println("  Config File: none (using defaults and environment variables)")
	}
}

func initConfig() error {
	configDir := os.Getenv("HOME") + "/.data-diff"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configFile := configDir + "/config.yaml"
	if _, err := os.Stat(configFile); err == nil {
		return fmt.Errorf("config file already exists: %s", configFile)
	}

	// Create default config file
	defaultConfig := `# data-diff configuration file
log_level: info
log_format: json
debug: false
format: json
`

	if err := os.WriteFile(configFile, []byte(defaultConfig), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Printf("Configuration file created: %s\n", configFile)
	return nil
}
