package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {

	logformat := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "060102150405",
			FullTimestamp:   true,
			ForceFormatting: false,
		},
	}

	logrus.SetFormatter(logformat.Formatter)

}
