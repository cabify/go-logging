package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Logger is the interface for outputing log messages in different levels.
// A new Logger can be created with NewLogger() function.
// You can changed the output handler with SetHandler() function.
type Logger interface {
	// SetLevel changes the level of the logger. Default is logging.Info.
	SetLevel(Level)

	// SetHandler replaces the current handler for output. Default is logging.StderrHandler.
	SetHandler(Handler)

	// SetCallDepth sets the parameter passed to runtime.Caller().
	// It is used to get the file name from call stack.
	// For example you need to set it to 1 if you are using a wrapper around
	// the Logger. Default value is zero.
	SetCallDepth(int)

	// Fatal is equivalent to Logger.Critical followed by a call to os.Exit(1).
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	// Panic is equivalent to Logger.Critical followed by a call to panic().
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	// Log functions
	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})
	Criticalln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})
	Noticeln(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
}

///////////////////////////
//                       //
// Logger implementation //
//                       //
///////////////////////////

// logger is the default Logger implementation.
type logger struct {
	Name      string
	Level     Level
	Handler   Handler
	calldepth int
}

// NewLogger returns a new Logger implementation. Do not forget to close it at exit.
func NewLogger(name string) Logger {
	return &logger{
		Name:    name,
		Level:   DefaultLevel,
		Handler: DefaultHandler,
	}
}

func (l *logger) SetLevel(level Level) { l.Level = level }
func (l *logger) SetHandler(b Handler) { l.Handler = b }
func (l *logger) SetCallDepth(n int)   { l.calldepth = n }

func (l *logger) log(level Level, args ...interface{}) {
	if level > l.Level {
		return
	}
	l.logMsg(level, fmt.Sprint(args...))
}

func (l *logger) logf(level Level, format string, args ...interface{}) {
	if level > l.Level {
		return
	}
	l.logMsg(level, fmt.Sprintf(format, args...))
}

func (l *logger) logln(level Level, args ...interface{}) {
	if level > l.Level {
		return
	}
	l.logMsg(level, fmt.Sprintln(args...))
}

func (l *logger) logMsg(level Level, message string) {
	if level > l.Level {
		return
	}

	_, file, line, ok := runtime.Caller(l.calldepth + 2)
	if !ok {
		file = "???"
		line = 0
	}

	rec := &Record{
		Message:     message,
		LoggerName:  l.Name,
		Level:       level,
		Time:        time.Now(),
		Filename:    file,
		Line:        line,
		ProcessName: procName,
		ProcessID:   pid,
	}

	l.Handler.Handle(rec)
}

// procName returns the name of the current process.
// func procName() string { return filepath.Base(os.Args[0]) }
var procName = filepath.Base(os.Args[0])
var pid = os.Getpid()

func (l *logger) Fatal(args ...interface{}) {
	l.log(CRITICAL, args...)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logf(CRITICAL, format, args...)
	os.Exit(1)
}

func (l *logger) Fatalln(args ...interface{}) {
	l.logln(CRITICAL, args...)
	os.Exit(1)
}

func (l *logger) Panic(args ...interface{}) {
	l.log(CRITICAL, args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logf(CRITICAL, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Panicln(args ...interface{}) {
	l.logln(CRITICAL, args...)
	panic(fmt.Sprintln(args...))
}

func (l *logger) Critical(args ...interface{}) {
	l.log(CRITICAL, args...)
}

func (l *logger) Criticalf(format string, args ...interface{}) {
	l.logf(CRITICAL, format, args...)
}

func (l *logger) Criticalln(args ...interface{}) {
	l.logln(CRITICAL, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log(ERROR, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

func (l *logger) Errorln(args ...interface{}) {
	l.logln(ERROR, args...)
}

func (l *logger) Warning(args ...interface{}) {
	l.log(WARNING, args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.logf(WARNING, format, args...)
}

func (l *logger) Warningln(args ...interface{}) {
	l.logln(WARNING, args...)
}

func (l *logger) Notice(args ...interface{}) {
	l.log(NOTICE, args...)
}

func (l *logger) Noticef(format string, args ...interface{}) {
	l.logf(NOTICE, format, args...)
}

func (l *logger) Noticeln(args ...interface{}) {
	l.logln(NOTICE, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log(INFO, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

func (l *logger) Infoln(args ...interface{}) {
	l.logln(INFO, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.log(DEBUG, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

func (l *logger) Debugln(args ...interface{}) {
	l.logln(DEBUG, args...)
}
