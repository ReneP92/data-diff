package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/renepersau/data-diff/internal/config"
	"github.com/renepersau/data-diff/pkg/diff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewCompareCommand creates the compare command
func NewCompareCommand(cfg *config.Config, logger *logrus.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare [flags] <source> <target>",
		Short: "Compare two data sources",
		Long: `Compare two data sources and show the differences.
The sources can be files, URLs, or other data formats supported by the tool.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompare(cmd, args, cfg, logger)
		},
	}

	// Add flags
	cmd.Flags().StringP("output", "o", "", "output file (default: stdout)")
	cmd.Flags().StringP("format", "f", "json", "output format (json, yaml, table)")
	cmd.Flags().Bool("ignore-case", false, "ignore case when comparing strings")
	cmd.Flags().StringSlice("ignore-fields", []string{}, "fields to ignore during comparison")
	cmd.Flags().Bool("show-unchanged", false, "show unchanged fields in output")

	return cmd
}

func runCompare(cmd *cobra.Command, args []string, cfg *config.Config, logger *logrus.Logger) error {
	source := args[0]
	target := args[1]

	output, _ := cmd.Flags().GetString("output")
	format, _ := cmd.Flags().GetString("format")
	ignoreCase, _ := cmd.Flags().GetBool("ignore-case")
	ignoreFields, _ := cmd.Flags().GetStringSlice("ignore-fields")
	showUnchanged, _ := cmd.Flags().GetBool("show-unchanged")

	logger.WithFields(logrus.Fields{
		"source":         source,
		"target":         target,
		"output":         output,
		"format":         format,
		"ignore_case":    ignoreCase,
		"ignore_fields":  ignoreFields,
		"show_unchanged": showUnchanged,
	}).Info("Starting comparison")

	// Create diff options
	options := &diff.Options{
		IgnoreCase:    ignoreCase,
		IgnoreFields:  ignoreFields,
		ShowUnchanged: showUnchanged,
		Format:        format,
	}

	// Perform the comparison
	result, err := diff.Compare(source, target, options)
	if err != nil {
		return fmt.Errorf("comparison failed: %w", err)
	}

	// Output the result
	var writer io.Writer = os.Stdout
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer file.Close()
		writer = file
	}

	if err := result.Write(writer, format); err != nil {
		return fmt.Errorf("failed to write result: %w", err)
	}

	logger.Info("Comparison completed successfully")
	return nil
}

