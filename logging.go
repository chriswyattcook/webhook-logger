package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

func setupLogging() {
	logger = logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	// newLogger.SetReportCaller(true)
	logger.SetFormatter(&ecslogrus.Formatter{
		DataKey: "labels",
	})
	logger.SetOutput(os.Stdout)
}
