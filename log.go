// Package log is an alternative to log package in standard library.
package log

import "os"

type Level int

// Logging levels.
const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var (
	DefaultLogger    Logger    = NewLogger(procName)
	DefaultLevel     Level     = INFO
	DefaultHandler   Handler   = NewFileHandler(os.Stderr)
	DefaultFormatter Formatter = defaultFormatter{}
)

///////////////////
//               //
// DefaultLogger //
//               //
///////////////////

// SetLevel changes the level of DefaultLogger and DefaultHandler.
func SetLevel(l Level) {
	DefaultLogger.SetLevel(l)
	DefaultHandler.SetLevel(l)
}

func Fatal(args ...interface{})                    { DefaultLogger.Fatal(args...) }
func Fatalf(format string, args ...interface{})    { DefaultLogger.Fatalf(format, args...) }
func Fatalln(args ...interface{})                  { DefaultLogger.Fatalln(args...) }
func Panic(args ...interface{})                    { DefaultLogger.Panic(args...) }
func Panicf(format string, args ...interface{})    { DefaultLogger.Panicf(format, args...) }
func Panicln(args ...interface{})                  { DefaultLogger.Panicln(args...) }
func Critical(args ...interface{})                 { DefaultLogger.Critical(args...) }
func Criticalf(format string, args ...interface{}) { DefaultLogger.Criticalf(format, args...) }
func Criticalln(args ...interface{})               { DefaultLogger.Criticalln(args...) }
func Error(args ...interface{})                    { DefaultLogger.Error(args...) }
func Errorf(format string, args ...interface{})    { DefaultLogger.Errorf(format, args...) }
func Errorln(args ...interface{})                  { DefaultLogger.Errorln(args...) }
func Warning(args ...interface{})                  { DefaultLogger.Warning(args...) }
func Warningf(format string, args ...interface{})  { DefaultLogger.Warningf(format, args...) }
func Warningln(args ...interface{})                { DefaultLogger.Warningln(args...) }
func Notice(args ...interface{})                   { DefaultLogger.Notice(args...) }
func Noticef(format string, args ...interface{})   { DefaultLogger.Noticef(format, args...) }
func Noticeln(args ...interface{})                 { DefaultLogger.Noticeln(args...) }
func Info(args ...interface{})                     { DefaultLogger.Info(args...) }
func Infof(format string, args ...interface{})     { DefaultLogger.Infof(format, args...) }
func Infoln(args ...interface{})                   { DefaultLogger.Infoln(args...) }
func Debug(args ...interface{})                    { DefaultLogger.Debug(args...) }
func Debugf(format string, args ...interface{})    { DefaultLogger.Debugf(format, args...) }
func Debugln(args ...interface{})                  { DefaultLogger.Debugln(args...) }

const (
	logLevelCritical = "critical"
	logLevelError    = "error"
	logLevelWarning  = "warning"
	logLevelNotice   = "notice"
	logLevelInfo     = "info"
	logLevelDebug    = "debug"
)

var logLevelMap = map[string]Level{
	logLevelCritical: CRITICAL,
	logLevelError:    ERROR,
	logLevelWarning:  WARNING,
	logLevelNotice:   NOTICE,
	logLevelInfo:     INFO,
	logLevelDebug:    DEBUG,
}

var logLevelNameMap = make(map[Level]string)

func init() {
	for name, value := range logLevelMap {
		logLevelNameMap[value] = name
	}
}
