package log

import (
	"os"
)

// Config defines the logging configuration
type Config struct {
	Level  string `default:"info"`
	Output string `default:"stdout"`
}

// ConfigureDefaultLogger configures loggers for your service, optionally adding log message counters with your favorite
// metrics system
func ConfigureDefaultLogger(name string, cfg Config, logCounters ...CountLogMessage) {
	if logLevel, ok := logLevelMap[cfg.Level]; ok {
		SetLevel(logLevel) // This sets the default level for all future
		DefaultLevel = logLevel
	} else {
		Warningf("Unknown log level configured: %s", cfg.Level)
	}

	var handler Handler = NewFileHandler(getLoggerOutput(cfg.Output))

	if len(logCounters) > 0 {
		handler = &metricsAgentLoggingHandler{
			Handler:     handler,
			logCounters: logCounters,
		}
	}
	handler.SetFormatter(DefaultFormatter)

	logger := NewLogger(name)
	if cfg.Level != logLevelDebug {
		logger = NoDebugLogger{
			Logger: logger,
		}
	}
	logger.SetHandler(handler)

	DefaultLogger = logger
	Infof("Configured default logger %s with log level %s", name, cfg.Level)
}

func getLoggerOutput(outputName string) *os.File {
	switch outputName {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	default:
		Warningf("Unknown logger output defined in the config: '%s'", outputName)
		return os.Stderr
	}
}
