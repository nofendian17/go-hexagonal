package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
	"user-svc/internal/shared/logger"
)

type LoggingMiddleware interface {
	LogRequestAndResponse(next echo.HandlerFunc) echo.HandlerFunc
}

type loggingMiddleware struct {
	logger *logger.LogWrapper
}

func NewLoggingMiddleware(logger *logger.LogWrapper) LoggingMiddleware {
	return &loggingMiddleware{
		logger: logger,
	}
}

func (l *loggingMiddleware) LogRequestAndResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Capture start time
		startTime := time.Now()

		// Wrap the next handler with the body dump middleware
		bodyDumpMiddleware := middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
			Skipper: nil,
			Handler: func(c echo.Context, reqBody, resBody []byte) {
				// Combine request and response into a single log entry
				logEntry := logger.Fields{
					"method":        c.Request().Method,
					"uri":           c.Request().RequestURI,
					"remote_ip":     c.RealIP(),
					"user_agent":    c.Request().UserAgent(),
					"headers":       c.Request().Header,
					"request_body":  string(reqBody),
					"response_body": string(resBody),
					"response_time": time.Since(startTime).Milliseconds(),
				}

				// Log the combined entry
				l.logger.WithFields(logEntry).Info("Request and response")
			},
		})

		// Invoke the wrapped handler
		return bodyDumpMiddleware(next)(c)
	}
}
