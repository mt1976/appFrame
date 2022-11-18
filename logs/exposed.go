package logs

import (
	"github.com/sirupsen/logrus"
)

// Fields type, used to pass to `WithFields`.
type Fields logrus.Fields

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Information(args ...interface{}) {
	Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(fields))
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	logrus.Warning(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	logrus.Println(args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	logrus.Error(args...)
}
