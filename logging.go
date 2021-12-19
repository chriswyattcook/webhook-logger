package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLogging() {
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
}
