package log

// NoDebugLogger embeds a Logger, but in calls to debug functions it does nothing.
// It avoids doing fmt.Sprintf() for those calls as they will be discarded anyways.
// This makes those calls like 50 times faster (see benchmark file)
type NoDebugLogger struct {
	Logger
}

func (NoDebugLogger) Debug(args ...interface{})                 {}
func (NoDebugLogger) Debugf(format string, args ...interface{}) {}
func (NoDebugLogger) Debugln(args ...interface{})               {}
