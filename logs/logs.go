package logs

import (
	"io"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func new() XLogger {

	newlog := XLogger{}
	lr := *logrus.New()

	logformat := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "060102∙150405",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	//this.toDisk = false
	basePath, _ := os.Getwd()
	//if err != nil {
	//	this.logger.Fatal(err)
	//}
	//this.path = basePath

	lr.SetFormatter(logformat.Formatter)
	newlog.log = lr
	newlog.path = basePath
	newlog.toDisk = false

	return newlog
}

func (l *XLogger) toFileAndConsole(name string) *XLogger {
	l.log.WithField("name", name).Info("Logging to file and console")
	l.toDisk = true
	l.setOutput(name, true)
	return l
}

func (l *XLogger) toConsole() *XLogger {
	l.log.Info("Logging to Console")
	l.toDisk = false
	l.log.SetOutput(os.Stdout)
	return l
}

func (l *XLogger) toFile(name string) {
	l.log.WithField("name", name).Info("Logging to file")
	l.toDisk = true
	l.setOutput(name, false)
}

func (l *XLogger) setOutput(name string, both bool) *XLogger {
	if name == "" {
		name = "default"
	}

	// get current os user name
	username, err := getUserName()

	filename := l.path + string(os.PathSeparator) + name + "_" + username + "_" + time.Now().Format("20060102") + ".log"
	l.Warn("Logging to file: " + filename)

	// Create a io.Writer instance
	l.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Fatal(err)
		return nil
	}

	if both {
		// Create a new MultiWriter
		mw := io.MultiWriter(os.Stdout, l.file)
		l.log.SetOutput(mw)
		return l
	}
	l.log.SetOutput(l.file)
	return l
}

func (l *XLogger) setPath(path string) *XLogger {
	if path == "" {
		path = "./"
	}
	l.path = path
	l.log.WithField("path", path).Info("Logging path set")
	return l
}

func getUserName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		usr.Username = "unknown"
	}

	user := usr.Username

	// if username contains "\"" then return the last part

	// The code `if len(user) > 0 && strings.Contains(user, "\")` checks if the length of the `user`
	// string is greater than 0 and if it contains the backslash character "\".
	if len(user) > 0 && strings.Contains(user, "\\") {
		user = user[strings.LastIndex(user, "\\")+1:]
	}

	if len(user) == 0 {
		user = usr.Name
	}

	if len(user) == 0 {
		user = "unknown"
	}
	return user, err
}
