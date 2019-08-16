// Package template provides a Logur adapter for TEMPLATE.
package template

import (
	"github.com/goph/logur"
)

// Logger is a Logur adapter for TEMPLATE.
type Logger struct {
}

// New returns a new Logur logger.
// If logger is nil, a default instance will be created.
func New(logger interface{}) *Logger {
	if logger == nil {
		return &Logger{}
	}

	return &Logger{}
}

// Trace implements the logur.Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {

}

// Debug implements the logur.Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {

}

// Info implements the logur.Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {

}

// Warn implements the logur.Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {

}

// Error implements the logur.Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {

}

// LevelEnabled implements logur.LevelEnabler interface.
func (l *Logger) LevelEnabled(level logur.Level) bool {
	switch level {
	case logur.Trace:
		return true
	case logur.Debug:
		return true
	case logur.Info:
		return true
	case logur.Warn:
		return true
	case logur.Error:
		return true
	}

	return true
}
