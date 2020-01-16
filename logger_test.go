package logrus

import (
	"testing"

	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"logur.dev/logur/conformance"

	"logur.dev/logur"
)

// nolint: gochecknoglobals
var levelMap = map[logur.Level]logrus.Level{
	logur.Trace: logrus.TraceLevel,
	logur.Debug: logrus.DebugLevel,
	logur.Info:  logrus.InfoLevel,
	logur.Warn:  logrus.WarnLevel,
	logur.Error: logrus.ErrorLevel,
}

func TestLogger(t *testing.T) {
	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			logrusLogger, hook := logrustest.NewNullLogger()
			logrusLogger.SetLevel(levelMap[level])

			return New(logrusLogger), conformance.TestLoggerFunc(func() []logur.LogEvent {
				entries := hook.AllEntries()

				events := make([]logur.LogEvent, len(entries))

				for key, entry := range entries {
					level, _ := logur.ParseLevel(entry.Level.String())

					events[key] = logur.LogEvent{
						Line:   entry.Message,
						Level:  level,
						Fields: entry.Data,
					}
				}

				return events
			})
		},
	}

	suite.Run(t)
}
