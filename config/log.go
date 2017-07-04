package config

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

// NewLogger create logger
func NewLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	logger = logrus.New()
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}

	return logger
}
