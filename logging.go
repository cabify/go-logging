// Package logging is an alternative to log package in standard library.
package log

import (
	"fmt"
	"os"
	"time"
)

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

var LevelNames = map[Level]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARNING:  "WARNING",
	NOTICE:   "NOTICE",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
}

var (
	DefaultLogger    Logger    = NewLogger(procName())
	DefaultLevel     Level     = INFO
	DefaultHandler   Handler   = StderrHandler
	DefaultFormatter Formatter = &defaultFormatter{}
	StdoutHandler              = NewWriterHandler(os.Stdout)
	StderrHandler              = NewWriterHandler(os.Stderr)
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

	// Fatal is equivalent to l.Critical followed by a call to os.Exit(1).
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	// Panic is equivalent to l.Critical followed by a call to panic().
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	// Critical logs a message using CRITICAL as log level.
	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})

	// Error logs a message using ERROR as log level.
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	// Warning logs a message using WARNING as log level.
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	// Notice logs a message using NOTICE as log level.
	Notice(args ...interface{})
	Noticef(format string, args ...interface{})

	// Info logs a message using INFO as log level.
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	// Debug logs a message using DEBUG as log level.
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

// Record contains all of the information about a single log message.
type Record struct {
	format      *string       // Format string
	args        []interface{} // Arguments to format string
	Message     string        // Formatted log message
	LoggerName  string        // Name of the logger module
	Level       Level         // Level of the record
	Time        time.Time     // Time of the record (local time)
	Filename    string        // File name of the log call (absolute path)
	Line        int           // Line number in file
	ProcessID   int           // PID
	ProcessName string        // Name of the process
}

// Formatter formats a record.
type Formatter interface {
	// Format the record and return a message.
	Format(*Record) (message string)
}

///////////////////////
//                   //
// Default Formatter //
//                   //
///////////////////////

type defaultFormatter struct{}

// Format outputs a message like "2014-02-28 18:15:57 [example] INFO     something happened"
func (f *defaultFormatter) Format(rec *Record) string {
	return fmt.Sprintf("%s [%s] %-8s %s", fmt.Sprint(rec.Time)[:19], rec.LoggerName, LevelNames[rec.Level], rec.Message)
}

///////////////////
//               //
// DefaultLogger //
//               //
///////////////////

// Fatal is equivalent to Critical() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) { DefaultLogger.Fatal(args...) }

// Fatal is equivalent to Criticalf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) { DefaultLogger.Fatalf(format, args...) }

// Panic is equivalent to Critical() followed by a call to panic().
func Panic(args ...interface{}) { DefaultLogger.Panic(args...) }

// Panic is equivalent to Criticalf() followed by a call to panic().
func Panicf(format string, args ...interface{}) { DefaultLogger.Panicf(format, args...) }

// Critical prints a critical level log message to the stderr.
func Critical(args ...interface{}) { DefaultLogger.Critical(args...) }

// Criticalf prints a critical level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Criticalf(format string, args ...interface{}) { DefaultLogger.Criticalf(format, args...) }

// Error prints a error level log message to the stderr.
func Error(args ...interface{}) { DefaultLogger.Error(args...) }

// Errorf prints a error level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) { DefaultLogger.Errorf(format, args...) }

// Warning prints a warning level log message to the stderr.
func Warning(args ...interface{}) { DefaultLogger.Warning(args...) }

// Warningf prints a warning level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) { DefaultLogger.Warningf(format, args...) }

// Notice prints a notice level log message to the stderr.
func Notice(args ...interface{}) { DefaultLogger.Notice(args...) }

// Noticef prints a notice level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Noticef(format string, args ...interface{}) { DefaultLogger.Noticef(format, args...) }

// Info prints a info level log message to the stderr.
func Info(args ...interface{}) { DefaultLogger.Info(args...) }

// Infof prints a info level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) { DefaultLogger.Infof(format, args...) }

// Debug prints a debug level log message to the stderr.
func Debug(args ...interface{}) { DefaultLogger.Debug(args...) }

// Debugf prints a debug level log message to the stderr. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) { DefaultLogger.Debugf(format, args...) }
