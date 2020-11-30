package logger

import (
	"github.com/sirupsen/logrus"
	"syscall"
)

func Init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func Debug(message string, fields logrus.Fields) {
	if fields != nil {
		logrus.WithFields(fields).Debug(message)
	} else {
		logrus.Debug(message)
	}
}

func Info(message string, fields logrus.Fields) {
	if fields != nil {
		logrus.WithFields(fields).Info(message)
	} else {
		logrus.Info(message)
	}
}

func Warn(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Warn(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Warn(message)
	} else {
		logrus.Warn(message)
	}
}

/**
 * Alias for logger.Warn
 */
func Warning(message string, err error, fields logrus.Fields) {
	Warn(message, err, fields)
}

func Error(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Error(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Error(message)
	} else {
		logrus.Error(message)
	}
}

func Fatal(message string, err error, fields logrus.Fields) {
	if fields != nil {
		fields["error"] = err
		logrus.WithFields(fields).Fatal(message)
	} else if err != nil {
		fields := logrus.Fields{"error": err}
		logrus.WithFields(fields).Fatal(message)
	} else {
		logrus.Fatal(message)
	}
	syscall.Exit(1)
}
