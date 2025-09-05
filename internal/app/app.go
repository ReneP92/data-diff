package app

import (
	"github.com/renepersau/data-diff/internal/commands"
	"github.com/renepersau/data-diff/internal/config"
	"github.com/sirupsen/logrus"
)

// App represents the main application
type App struct {
	config *config.Config
	logger *logrus.Logger
}

// New creates a new application instance
func New(cfg *config.Config, log *logrus.Logger) *App {
	return &App{
		config: cfg,
		logger: log,
	}
}

// Run starts the application
func (a *App) Run() error {
	a.logger.Info("Starting data-diff application")

	// Create root command
	rootCmd := commands.NewRootCommand(a.config, a.logger)

	// Execute the command
	return rootCmd.Execute()
}

