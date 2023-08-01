package logs

import (
	"io"

	"github.com/sirupsen/logrus"
)

type XLogger struct {
	log logrus.Logger
	//	Fields logrus.Fields
	toDisk bool
	file   io.Writer
	path   string
	prefix string
}

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// l is the internal logger
func New() XLogger {
	return new()
}

// Info logs a message at level Info on the standard shared.logger.
func (l *XLogger) Info(args ...interface{}) {
	l.log.Info(args...)
}

// Warn logs a message at level Warn on the standard shared.logger.
// Deprecated: Use Warn instead
func (l *XLogger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

// Trace logs a message at level Trace on the standard shared.logger.
func (l *XLogger) Trace(args ...interface{}) {
	l.log.Trace(args...)
}

// Debug logs a message at level Debug on the standard shared.logger.
func (l *XLogger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (l *XLogger) WithField(key string, value interface{}) *logrus.Entry {
	return l.log.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (l *XLogger) WithFields(fields Fields) *logrus.Entry {
	x := make(logrus.Fields, len(fields))
	for k, v := range fields {
		x[k] = v
	}
	return l.log.WithFields(x)
}

// Panic logs a message at level Panic on the standard shared.logger.
func (l *XLogger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l *XLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

// Warning logs a message at level Warn on the standard shared.logger.
func (l *XLogger) Warning(args ...interface{}) {
	l.log.Warning(args...)
}

// Println logs a message at level Info on the standard shared.logger.
func (l *XLogger) Println(args ...interface{}) {
	l.log.Println(args...)
}

// Printf logs a message at level Info on the standard shared.logger.
func (l *XLogger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}

// Error logs a message at level Error on the standard shared.logger.
func (l *XLogger) Error(args ...interface{}) {
	l.log.Error(args...)
}

// ToFile sets the logging to file
func (l *XLogger) ToFile(name string) {
	l.toFile(name)
}

// ToConsole sets the logging to console
func (l *XLogger) ToConsole() {
	l.toConsole()
}

// ToFileAndConsole sets the logging to file and console
func (l *XLogger) ToFileAndConsole(name string) {
	l.toFileAndConsole(name)
}

// Start starts the logging
func (l *XLogger) Start() *XLogger {
	l.toConsole()
	l.log.Info("Logging Started")
	return l
}

// SetPath sets the path for the log file
func (l *XLogger) SetPath(path string) *XLogger {
	l.setPath(path)
	return l
}
