package log

// CountLogMessage offers a method to count log messages logged by the logger
type CountLogMessage func(level string)

type metricsAgentLoggingHandler struct {
	Handler
	logCounters []CountLogMessage
}

// Handle decorates Handler's Handle method counting the log messages
func (h *metricsAgentLoggingHandler) Handle(record *Record) {
	level := "unknown"
	if levelName, ok := logLevelNameMap[record.Level]; ok {
		level = levelName
	}
	for _, countLogMessage := range h.logCounters {
		countLogMessage(level)
	}
	h.Handler.Handle(record)
}
