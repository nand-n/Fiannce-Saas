package utils

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Log = logrus.New()

func InitLogger() {
    // Set the output to stdout instead of the default stderr
    Log.Out = os.Stdout

    // Set the log level (e.g., DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel)
    Log.SetLevel(logrus.InfoLevel)

    // Set the formatter
    Log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
}
