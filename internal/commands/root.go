package commands

import (
	"fmt"

	"github.com/renepersau/data-diff/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRootCommand creates the root command for the CLI
func NewRootCommand(cfg *config.Config, logger *logrus.Logger) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "data-diff",
		Short: "A tool for comparing data structures",
		Long: `data-diff is a command-line tool for comparing and analyzing differences
between data structures. It supports various input formats and provides
detailed comparison reports.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Bind flags to viper
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return fmt.Errorf("failed to bind flags: %w", err)
			}
			return nil
		},
	}

	// Add persistent flags
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.data-diff/config.yaml)")
	rootCmd.PersistentFlags().String("log-level", cfg.LogLevel, "log level (debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().String("log-format", cfg.LogFormat, "log format (json, text)")
	rootCmd.PersistentFlags().Bool("debug", cfg.Debug, "enable debug mode")

	// Add subcommands
	rootCmd.AddCommand(NewCompareCommand(cfg, logger))
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(NewConfigCommand(cfg))

	return rootCmd
}
