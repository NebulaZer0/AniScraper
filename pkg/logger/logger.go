package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = &logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	},
}
