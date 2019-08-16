package logrus_test

import (
	"github.com/sirupsen/logrus"

	logrusadapter "logur.dev/adapter/logrus"
)

func ExampleNew() {
	logger := logrusadapter.New(logrus.New())

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := logrusadapter.New(nil)

	// Output:
	_ = logger
}

// The logger can be instantiated with a Logrus entry.
// It's useful if you want to create a logger from an already populated Logrus instance.
func ExampleNewFromEntry() {
	entry := logrus.New().WithField("key", "value")

	logger := logrusadapter.NewFromEntry(entry)

	// Output:
	_ = logger
}

// If entry is nil, a default instance is created.
func ExampleNewFromEntry_default() {
	logger := logrusadapter.NewFromEntry(nil)

	// Output:
	_ = logger
}
