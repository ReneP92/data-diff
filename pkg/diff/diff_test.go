package diff

import (
	"strings"
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		target   string
		options  *Options
		expected bool
	}{
		{
			name:     "identical strings",
			source:   "test",
			target:   "test",
			options:  &Options{},
			expected: true,
		},
		{
			name:     "different strings",
			source:   "test1",
			target:   "test2",
			options:  &Options{},
			expected: false,
		},
		{
			name:     "case insensitive match",
			source:   "Test",
			target:   "test",
			options:  &Options{IgnoreCase: true},
			expected: true,
		},
		{
			name:     "case sensitive mismatch",
			source:   "Test",
			target:   "test",
			options:  &Options{IgnoreCase: false},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Compare(tt.source, tt.target, tt.options)
			if err != nil {
				t.Fatalf("Compare() error = %v", err)
			}

			if result.Equal != tt.expected {
				t.Errorf("Compare() Equal = %v, expected %v", result.Equal, tt.expected)
			}
		})
	}
}

func TestResult_Write(t *testing.T) {
	result := &Result{
		Source: "source",
		Target: "target",
		Equal:  true,
		Summary: Summary{
			TotalFields:     1,
			EqualFields:     1,
			DifferentFields: 0,
		},
	}

	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "JSON format",
			format:   "json",
			expected: `"source"`,
		},
		{
			name:     "YAML format",
			format:   "yaml",
			expected: `source:`,
		},
		{
			name:     "Table format",
			format:   "table",
			expected: `Comparison Result`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := result.Write(&buf, tt.format)
			if err != nil {
				t.Fatalf("Write() error = %v", err)
			}

			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Write() output = %v, expected to contain %v", output, tt.expected)
			}
		})
	}
}

func TestCompareValues(t *testing.T) {
	tests := []struct {
		name     string
		source   interface{}
		target   interface{}
		options  *Options
		expected bool
	}{
		{
			name:     "identical values",
			source:   "test",
			target:   "test",
			options:  &Options{},
			expected: true,
		},
		{
			name:     "different values",
			source:   "test1",
			target:   "test2",
			options:  &Options{},
			expected: false,
		},
		{
			name:     "case insensitive match",
			source:   "Test",
			target:   "test",
			options:  &Options{IgnoreCase: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compareValues(tt.source, tt.target, tt.options)
			if result != tt.expected {
				t.Errorf("compareValues() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

