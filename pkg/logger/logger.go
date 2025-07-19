// Package logger provides a wrapper around zerolog with file management capabilities.
package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

// logger wraps zerolog.Logger and provides additional functionality for managing log files.
// It embeds zerolog.Logger to provide all standard logging methods while adding
// file management capabilities through the logFile field.
type logger struct {
	zerolog.Logger
	logFile *os.File
}

// Close closes the underlying log file if it exists.
// It returns an error if the file cannot be closed properly.
// If no log file is associated with the logger, it returns nil.
func (l *logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// NewLogger creates a new logger instance with specified console and file log levels.
// The levelConsole parameter sets the minimum log level for console output,
// while levelFile sets the minimum log level for file output.
// The function creates a logs directory if it doesn't exist and opens/creates
// an "app.log" file for logging. It sets up a multi-writer that outputs to both
// console (with timestamp format "15:04:05") and file.
// The global log level is set to the more restrictive (lower) of the two levels.
// Returns an error if log levels are invalid, directory creation fails, or
// log file cannot be opened.
func NewLogger(levelConsole, levelFile string) (*logger, error) {
	parsedConsoleLevel, err := zerolog.ParseLevel(levelConsole)
	if err != nil {
		return nil, fmt.Errorf("invalid console log level: %w", err)
	}
	parsedFileLevel, err := zerolog.ParseLevel(levelFile)
	if err != nil {
		return nil, fmt.Errorf("invalid file log level: %w", err)
	}
	var globalLevel zerolog.Level
	if parsedConsoleLevel < parsedFileLevel {
		globalLevel = parsedConsoleLevel
	} else {
		globalLevel = parsedFileLevel
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("error creating logs directory: %w", err)
	}
	consoleWriter.NoColor = false // Enable colors in console output
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %w", err)
	}
	multiWriter := zerolog.MultiLevelWriter(consoleWriter, logFile)
	zLogger := zerolog.New(multiWriter).Level(globalLevel).With().Timestamp().Logger()

	return &logger{
		Logger:  zLogger,
		logFile: logFile,
	}, nil
}
