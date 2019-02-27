package log

import (
	"context"

	"github.com/cabify/go-logging/internal"
)

type baggageLogger struct {
	Logger
	ctx context.Context
}

func newBaggageLogger(ctx context.Context, base Logger) baggageLogger {
	return baggageLogger{
		Logger: base,
		ctx:    ctx,
	}
}

func (l baggageLogger) getContextString() string {
	baggage, ok := l.ctx.Value(internal.BaggageContextKey).(internal.Baggage)
	if !ok {
		return ""
	}

	return baggage.String() + ": "
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