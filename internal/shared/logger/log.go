package logger

import (
	"fmt"
	"io"
	"os"
	"time"
	"user-svc/internal/shared/config"

	"github.com/sirupsen/logrus"
)

const (
	logFormat    = "2006-01-02"
	logExtension = ".log"
)

type Fields = map[string]interface{}

type Logger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	WithFields(fields Fields) logrus.FieldLogger
}

type LogWrapper struct {
	logger *logrus.Logger
}

func (l *LogWrapper) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *LogWrapper) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LogWrapper) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LogWrapper) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LogWrapper) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogWrapper) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LogWrapper) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *LogWrapper) WithFields(fields Fields) logrus.FieldLogger {
	return l.logger.WithFields(fields)
}

func NewLogger(config *config.Config) *LogWrapper {
	fileLocation := config.Log.FileLocation
	stdout := config.Log.Stdout

	var out io.Writer
	if stdout {
		out = os.Stdout
	} else {
		filename := fileLocation + "/" + time.Now().Format(logFormat) + logExtension
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			if err := os.MkdirAll(fileLocation, 0755); err != nil {
				logrus.Error(fmt.Sprintf("failed to create logger directory: %s", err))
			}
			if _, err := os.Create(filename); err != nil {
				logrus.Error(fmt.Sprintf("failed to create logger file: %s", err))
			}
		}
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Error(fmt.Sprintf("failed to open logger file: %s", err))
		}
		out = io.MultiWriter(file, os.Stdout)
		go func() {
			for {
				if fileInfo, err := file.Stat(); err == nil && fileInfo.Size() > 1024*1024*30 {
					// max size reached, rotate logger file
					file.Close()
					filename := fileLocation + "/" + time.Now().Format(logFormat) + logExtension
					file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
					if err != nil {
						// failed to create new logger file, ignore and continue
						continue
					}
					out = io.MultiWriter(file, os.Stdout)
				}
				time.Sleep(time.Minute)
			}
		}()
	}

	level := logrus.TraceLevel
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetOutput(out)
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  time.RFC3339Nano,
		DisableColors:    false,
		DisableTimestamp: false,
	})

	defer func() {
		if err := recover(); err != nil {
			logger.WithField("panic", err).Error("Recovered panic")
		}
	}()

	return &LogWrapper{logger: logger}
}
