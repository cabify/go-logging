package log

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

type baggageLogger struct {
	Logger
	ctx context.Context
}

func baggageString(b map[string]interface{}) string {
	var kvPairs []string
	for key, value := range b {
		kvPairs = append(kvPairs, fmt.Sprintf("%s:%v", key, value))
	}

	sort.Strings(kvPairs)

	return strings.Join(kvPairs, ": ")
}

// Factory provides context aware loggers.
type Factory struct {
	baseLogger Logger
}

// NewFactory instantiates a factory with the default logger.
func NewFactory() Factory {
	return Factory{
		baseLogger: DefaultLogger,
	}
}

// For provides a logger which is aware of the passed context and will prepend the context baggage values.
func (f Factory) For(ctx context.Context) Logger {
	return newBaggageLogger(ctx, f.baseLogger)
}

func newBaggageLogger(ctx context.Context, base Logger) baggageLogger {
	return baggageLogger{
		Logger: base,
		ctx:    ctx,
	}
}

func (l baggageLogger) getContextString() string {
	baggage, ok := l.ctx.Value(BaggageContextKey).(map[string]interface{})
	if !ok {
		return ""
	}

	return baggageString(baggage) + ": "
}

func (l baggageLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(l.getContextString()+format, args...)
}

func (l baggageLogger) Fatalln(args ...interface{}) {
	l.Logger.Fatalln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Panic(args ...interface{}) {
	l.Logger.Panic(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(l.getContextString()+format, args...)
}

func (l baggageLogger) Panicln(args ...interface{}) {
	l.Logger.Panicln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Critical(args ...interface{}) {
	l.Logger.Critical(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Criticalf(format string, args ...interface{}) {
	l.Logger.Criticalf(l.getContextString()+format, args...)
}

func (l baggageLogger) Criticalln(args ...interface{}) {
	l.Logger.Criticalln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Error(args ...interface{}) {
	l.Logger.Error(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(l.getContextString()+format, args...)
}

func (l baggageLogger) Errorln(args ...interface{}) {
	l.Logger.Errorln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Warning(args ...interface{}) {
	l.Logger.Warning(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Warningf(format string, args ...interface{}) {
	l.Logger.Warningf(l.getContextString()+format, args...)
}

func (l baggageLogger) Warningln(args ...interface{}) {
	l.Logger.Warningln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Notice(args ...interface{}) {
	l.Logger.Notice(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Noticef(format string, args ...interface{}) {
	l.Logger.Noticef(l.getContextString()+format, args...)
}

func (l baggageLogger) Noticeln(args ...interface{}) {
	l.Logger.Noticeln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Info(args ...interface{}) {
	l.Logger.Info(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(l.getContextString()+format, args...)
}

func (l baggageLogger) Infoln(args ...interface{}) {
	l.Logger.Infoln(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Debug(args ...interface{}) {
	l.Logger.Debug(append([]interface{}{l.getContextString()}, args...)...)
}

func (l baggageLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(l.getContextString()+format, args...)
}

func (l baggageLogger) Debugln(args ...interface{}) {
	l.Logger.Debugln(append([]interface{}{l.getContextString()}, args...)...)
}
