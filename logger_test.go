package template

import (
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/logtesting"
)

// nolint: gochecknoglobals
var levelMap = map[logur.Level]string{
	logur.Trace: "trace",
	logur.Debug: "debug",
	logur.Info:  "info",
	logur.Warn:  "warn",
	logur.Error: "error",
}

func newTestSuite() *logtesting.LoggerTestSuite {
	return &logtesting.LoggerTestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, func() []logur.LogEvent) {
			var logger interface{}
			_ = levelMap

			return New(logger), func() []logur.LogEvent {
				return []logur.LogEvent{}
			}
		},
	}
}

func TestLoggerSuite(t *testing.T) {
	t.Skip("implement me")

	newTestSuite().Execute(t)
}
