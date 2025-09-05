package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// New creates a new logger instance with the specified level and format
func New(level, format string) *logrus.Logger {
	logger := logrus.New()

	// Set log level
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	// Set log format
	switch format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	// Set output to stdout
	logger.SetOutput(os.Stdout)

	return logger
}

// NewWithFields creates a new logger with predefined fields
func NewWithFields(level, format string, fields logrus.Fields) *logrus.Entry {
	return New(level, format).WithFields(fields)
}

