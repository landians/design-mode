package decorator

import "testing"

func Test_Decoration(t *testing.T) {
	fnTestlogger := func(factory ILoggerFactory) {
		logger := factory.GetLogger("testing.Test_Decoration")
		logger.Debug("This is a Debug msg")
		logger.Info("This is an Info msg")
	}

	fnTestlogger(ConsoleLoggerFactory)
	fnTestlogger(JsonLoggerFactory)
}
