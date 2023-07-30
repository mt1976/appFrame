package logs

import "github.com/sirupsen/logrus"

// Fields type, used to pass to `WithFields`.
type Fields logrus.Fields

// Info logs a message at level Info on the standard shared.logger.
func Info(args ...interface{}) {
	this.logger.Info(args...)
}

// Information logs a message at level Info on the standard shared.logger.
// Deprecated: Use Info instead
func Information(args ...interface{}) {
	Info(args...)
}

// Warn logs a message at level Warn on the standard shared.logger.
func Warn(args ...interface{}) {
	this.logger.Warn(args...)
}

// Trace logs a message at level Trace on the standard shared.logger.
func Trace(args ...interface{}) {
	this.logger.Trace(args...)
}

// Debug logs a message at level Debug on the standard shared.logger.
func Debug(args ...interface{}) {
	this.logger.Debug(args...)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return this.logger.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	return this.logger.WithFields(logrus.Fields(fields))
}

// Panic logs a message at level Panic on the standard shared.logger.
func Panic(args ...interface{}) {
	this.logger.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	this.logger.Fatal(args...)
}

// Warning logs a message at level Warn on the standard shared.logger.
func Warning(args ...interface{}) {
	this.logger.Warning(args...)
}

// Println logs a message at level Info on the standard shared.logger.
func Println(args ...interface{}) {
	this.logger.Println(args...)
}

// Printf logs a message at level Info on the standard shared.logger.
func Printf(format string, args ...interface{}) {
	this.logger.Printf(format, args...)
}

// Error logs a message at level Error on the standard shared.logger.
func Error(args ...interface{}) {
	this.logger.Error(args...)
}

// ToFile sets the logging to file
func ToFile(name string) {
	toFile(name)
}

// ToConsole sets the logging to console
func ToConsole() {
	toConsole()
}

func ToFileAndConsole(name string) {
	toFileAndConsole(name)
}

func Start() {
	toConsole()
	this.logger.Info("Logging Started")
}

func SetPath(path string) {
	setPath(path)
}
