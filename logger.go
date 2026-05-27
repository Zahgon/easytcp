package easytcp

import (
	"log"
)

var _ Logger = &DefaultLogger{}

// _log is the instance of Logger interface.
var _log Logger = newDiscardLogger()

// Logger is the generic interface for log recording.
type Logger interface {
	Errorf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
}

func newDiscardLogger() *DefaultLogger { _ = "STUB: not implemented"; return nil }

// DefaultLogger is the default logger instance for this package.
// DefaultLogger uses the built-in log.Logger.
type DefaultLogger struct {
	rawLogger *log.Logger
}

// Errorf implements Logger Errorf method.
func (d *DefaultLogger) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Tracef implements Logger Tracef method.
func (d *DefaultLogger) Tracef(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Log returns the package logger.
func Log() Logger {
	_ = "STUB: not implemented"

	// SetLogger sets the package logger.
	return *new(Logger)
}

func SetLogger(lg Logger) { _ = "STUB: not implemented"; return }
