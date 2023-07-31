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

type commondata struct {
	logger logrus.Logger
	toDisk bool
	file   io.Writer
	path   string
}

var this commondata

func init() {

	this.logger = *logrus.New()

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
	this.toDisk = false
	basePath, err := os.Getwd()
	if err != nil {
		this.logger.Fatal(err)
	}
	this.path = basePath
	this.logger.SetFormatter(logformat.Formatter)

}

func toFileAndConsole(name string) {
	this.logger.WithField("name", name).Info("Logging to file and console")
	this.toDisk = true
	setOutput(name, true)
}

func toConsole() {
	this.logger.Info("Logging to Console")
	this.toDisk = false
	this.logger.SetOutput(os.Stdout)
}

func toFile(name string) {
	this.logger.WithField("name", name).Info("Logging to file")
	this.toDisk = true
	setOutput(name, false)
}

func setOutput(name string, both bool) {

	if name == "" {
		name = "default"
	}

	// get current os user name
	username, err := getUserName()

	filename := this.path + string(os.PathSeparator) + name + "_" + username + "_" + time.Now().Format("20060102") + ".log"
	this.logger.Warn("Logging to file: " + filename)

	// Create a io.Writer instance
	this.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		this.logger.Fatal(err)
	}

	if both {
		// Create a new MultiWriter
		mw := io.MultiWriter(os.Stdout, this.file)
		this.logger.SetOutput(mw)
		return
	}
	this.logger.SetOutput(this.file)
}

func getUserName() (string, error) {
	usr, err := user.Current()
	if err != nil {
		this.logger.Fatal(err)
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
		this.logger.Fatal("Could not get username")
	}
	return user, err
}

func setPath(path string) {
	if path == "" {
		path = "./"
	}
	this.path = path
	this.logger.WithField("path", path).Info("Logging path set")
}
