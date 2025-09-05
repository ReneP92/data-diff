package diff

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Options contains configuration for the diff operation
type Options struct {
	IgnoreCase    bool
	IgnoreFields  []string
	ShowUnchanged bool
	Format        string
}

// Result represents the result of a comparison
type Result struct {
	Source      string       `json:"source"`
	Target      string       `json:"target"`
	Equal       bool         `json:"equal"`
	Differences []Difference `json:"differences,omitempty"`
	Summary     Summary      `json:"summary"`
}

// Difference represents a single difference between source and target
type Difference struct {
	Path    string      `json:"path"`
	Type    string      `json:"type"` // "added", "removed", "modified"
	Source  interface{} `json:"source,omitempty"`
	Target  interface{} `json:"target,omitempty"`
	Message string      `json:"message"`
}

// Summary provides statistics about the comparison
type Summary struct {
	TotalFields     int `json:"total_fields"`
	EqualFields     int `json:"equal_fields"`
	DifferentFields int `json:"different_fields"`
	AddedFields     int `json:"added_fields"`
	RemovedFields   int `json:"removed_fields"`
}

// Compare performs a comparison between two data sources
func Compare(source, target string, options *Options) (*Result, error) {
	// Load source data
	sourceData, err := loadData(source)
	if err != nil {
		return nil, fmt.Errorf("failed to load source data: %w", err)
	}

	// Load target data
	targetData, err := loadData(target)
	if err != nil {
		return nil, fmt.Errorf("failed to load target data: %w", err)
	}

	// Perform comparison
	result := &Result{
		Source: source,
		Target: target,
	}

	// Simple comparison logic (you can expand this based on your needs)
	if compareValues(sourceData, targetData, options) {
		result.Equal = true
		result.Summary = Summary{
			TotalFields:     1,
			EqualFields:     1,
			DifferentFields: 0,
			AddedFields:     0,
			RemovedFields:   0,
		}
	} else {
		result.Equal = false
		result.Differences = []Difference{
			{
				Path:    "root",
				Type:    "modified",
				Source:  sourceData,
				Target:  targetData,
				Message: "Data structures differ",
			},
		}
		result.Summary = Summary{
			TotalFields:     1,
			EqualFields:     0,
			DifferentFields: 1,
			AddedFields:     0,
			RemovedFields:   0,
		}
	}

	return result, nil
}

// Write outputs the result in the specified format
func (r *Result) Write(writer io.Writer, format string) error {
	switch strings.ToLower(format) {
	case "json":
		encoder := json.NewEncoder(writer)
		encoder.SetIndent("", "  ")
		return encoder.Encode(r)
	case "yaml":
		return yaml.NewEncoder(writer).Encode(r)
	case "table":
		return r.writeTable(writer)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

// loadData loads data from a file or returns a placeholder
func loadData(source string) (interface{}, error) {
	// Check if it's a file
	if _, err := os.Stat(source); err == nil {
		data, err := os.ReadFile(source)
		if err != nil {
			return nil, err
		}

		// Try to parse as JSON first
		var jsonData interface{}
		if err := json.Unmarshal(data, &jsonData); err == nil {
			return jsonData, nil
		}

		// Try to parse as YAML
		var yamlData interface{}
		if err := yaml.Unmarshal(data, &yamlData); err == nil {
			return yamlData, nil
		}

		// Return as string if neither JSON nor YAML
		return string(data), nil
	}

	// For now, return the source string as placeholder
	// In a real implementation, you might support URLs, databases, etc.
	return source, nil
}

// compareValues compares two values with the given options
func compareValues(source, target interface{}, options *Options) bool {
	// Simple string comparison for now
	sourceStr := fmt.Sprintf("%v", source)
	targetStr := fmt.Sprintf("%v", target)

	if options.IgnoreCase {
		return strings.EqualFold(sourceStr, targetStr)
	}

	return sourceStr == targetStr
}

// writeTable writes the result in table format
func (r *Result) writeTable(writer io.Writer) error {
	fmt.Fprintf(writer, "Comparison Result\n")
	fmt.Fprintf(writer, "=================\n\n")
	fmt.Fprintf(writer, "Source: %s\n", r.Source)
	fmt.Fprintf(writer, "Target: %s\n", r.Target)
	fmt.Fprintf(writer, "Equal: %t\n\n", r.Equal)

	if !r.Equal && len(r.Differences) > 0 {
		fmt.Fprintf(writer, "Differences:\n")
		fmt.Fprintf(writer, "------------\n")
		for _, diff := range r.Differences {
			fmt.Fprintf(writer, "Path: %s\n", diff.Path)
			fmt.Fprintf(writer, "Type: %s\n", diff.Type)
			fmt.Fprintf(writer, "Message: %s\n", diff.Message)
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "Summary:\n")
	fmt.Fprintf(writer, "--------\n")
	fmt.Fprintf(writer, "Total Fields: %d\n", r.Summary.TotalFields)
	fmt.Fprintf(writer, "Equal Fields: %d\n", r.Summary.EqualFields)
	fmt.Fprintf(writer, "Different Fields: %d\n", r.Summary.DifferentFields)
	fmt.Fprintf(writer, "Added Fields: %d\n", r.Summary.AddedFields)
	fmt.Fprintf(writer, "Removed Fields: %d\n", r.Summary.RemovedFields)

	return nil
}

