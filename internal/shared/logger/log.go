package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-svc/internal/shared/config"
	"user-svc/internal/shared/open_search"
)

type FieldMap = map[string]interface{}

type Logger interface {
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	WithFields(fields FieldMap) *logrus.Entry
}

type LoggerWrapper struct {
	logger           *logrus.Logger
	openSearchClient *open_search.OpenSearchClient
}

func (l *LoggerWrapper) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *LoggerWrapper) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *LoggerWrapper) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LoggerWrapper) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LoggerWrapper) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LoggerWrapper) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LoggerWrapper) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *LoggerWrapper) WithFields(fields FieldMap) *logrus.Entry {
	return l.logger.WithFields(fields)
}

func NewLogger(config *config.Config, openSearchClient *open_search.OpenSearchClient) *LoggerWrapper {
	// Config log
	path := config.Log.File.FileLocation
	isLogFileEnable := config.Log.File.Enable
	maxAge := config.Log.File.MaxAge
	isCompressed := config.Log.File.Compress

	// Config openSearch
	indexName := config.App.Name
	isOpenSearchEnable := config.Log.OpenSearch.Enable

	out := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    0, // megabytes
		MaxBackups: 0,
		MaxAge:     maxAge,       // days
		Compress:   isCompressed, // disabled by default
		LocalTime:  true,
	}

	level := logrus.TraceLevel
	logger := logrus.New()
	logger.SetReportCaller(false)
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  time.RFC3339Nano,
		DisableTimestamp: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
	})

	if isLogFileEnable {
		logger.SetOutput(out)
	}

	if isOpenSearchEnable {
		// Create the OpenSearch hook
		openSearchHook := open_search.NewOpenSearchHook(openSearchClient.Client, indexName)

		// Add the OpenSearch hook to the logger
		logger.AddHook(openSearchHook)
	}
	// Calculate the time until the next day
	now := time.Now()
	nextDay := now.Truncate(24 * time.Hour).Add(24 * time.Hour)
	sleepDuration := time.Until(nextDay)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP)

	go func() {
		for {
			select {
			case <-signalChan:
				err := out.Rotate()
				if err != nil {
					return
				}
			case <-time.After(sleepDuration):
				err := out.Rotate()
				if err != nil {
					return
				}
				sleepDuration = 24 * time.Hour
			}
		}
	}()

	return &LoggerWrapper{
		logger:           logger,
		openSearchClient: openSearchClient,
	}
}
