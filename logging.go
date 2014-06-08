// Package logging is an alternative to log package in standard library.
package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type (
	Color int
	Level int
)

// Colors for different log levels.
const (
	BLACK Color = (iota + 30)
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

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

var LevelColors = map[Level]Color{
	CRITICAL: MAGENTA,
	ERROR:    RED,
	WARNING:  YELLOW,
	NOTICE:   GREEN,
	INFO:     WHITE,
	DEBUG:    CYAN,
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

// Handler handles the output.
type Handler interface {
	SetFormatter(Formatter)
	SetLevel(Level)

	// Handle single log record.
	Handle(*Record)

	// Close the handler.
	Close()
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

func (l *logger) SetLevel(level Level) {
	l.Level = level
}

func (l *logger) SetHandler(b Handler) {
	l.Handler = b
}

func (l *logger) SetCallDepth(n int) {
	l.calldepth = n
}

// Fatal is equivalent to Critical() followed by a call to os.Exit(1).
func (l *logger) Fatal(args ...interface{}) {
	l.Critical(args...)
	os.Exit(1)
}

// Fatalf is equivalent to Criticalf() followed by a call to os.Exit(1).
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Criticalf(format, args...)
	os.Exit(1)
}

// Panic is equivalent to Critical() followed by a call to panic().
func (l *logger) Panic(args ...interface{}) {
	l.Critical(args...)
	panic(fmt.Sprint(args...))
}

// Panicf is equivalent to Criticalf() followed by a call to panic().
func (l *logger) Panicf(format string, args ...interface{}) {
	l.Criticalf(format, args...)
	panic(fmt.Sprintf(format, args...))
}

// Critical sends a critical level log message to the handler.
func (l *logger) Critical(args ...interface{}) { l.log(CRITICAL, nil, args...) }

// Criticalf sends a critical level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Criticalf(format string, args ...interface{}) { l.log(CRITICAL, &format, args...) }

// Error sends a error level log message to the handler.
func (l *logger) Error(args ...interface{}) { l.log(ERROR, nil, args...) }

// Errorf sends a error level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Errorf(format string, args ...interface{}) { l.log(ERROR, &format, args...) }

// Warning sends a warning level log message to the handler.
func (l *logger) Warning(args ...interface{}) { l.log(WARNING, nil, args...) }

// Warningf sends a warning level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Warningf(format string, args ...interface{}) { l.log(WARNING, &format, args...) }

// Notice sends a notice level log message to the handler.
func (l *logger) Notice(args ...interface{}) { l.log(NOTICE, nil, args...) }

// Noticef sends a notice level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Noticef(format string, args ...interface{}) { l.log(NOTICE, &format, args...) }

// Info sends a info level log message to the handler.
func (l *logger) Info(args ...interface{}) { l.log(INFO, nil, args...) }

// Infof sends a info level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Infof(format string, args ...interface{}) { l.log(INFO, &format, args...) }

// Debug sends a debug level log message to the handler.
func (l *logger) Debug(args ...interface{}) { l.log(DEBUG, nil, args...) }

// Debugf sends a debug level log message to the handler. Arguments are handled in the manner of fmt.Printf.
func (l *logger) Debugf(format string, args ...interface{}) { l.log(DEBUG, &format, args...) }

func (l *logger) log(level Level, format *string, args ...interface{}) {
	if level > l.Level {
		return
	}

	_, file, line, ok := runtime.Caller(l.calldepth + 2)
	if !ok {
		file = "???"
		line = 0
	}

	rec := &Record{
		format:      format,
		args:        args,
		LoggerName:  l.Name,
		Level:       level,
		Time:        time.Now(),
		Filename:    file,
		Line:        line,
		ProcessName: procName(),
		ProcessID:   os.Getpid(),
	}

	if format != nil {
		rec.Message = fmt.Sprintf(*format, args...)
	} else {
		rec.Message = fmt.Sprint(args...)
	}

	// Add missing newline at the end.
	if !strings.HasSuffix(rec.Message, "\n") {
		rec.Message += "\n"
	}

	l.Handler.Handle(rec)
}

// procName returns the name of the current process.
func procName() string { return filepath.Base(os.Args[0]) }

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

/////////////////
//             //
// BaseHandler //
//             //
/////////////////

type BaseHandler struct {
	Level     Level
	Formatter Formatter
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{
		Level:     DefaultLevel,
		Formatter: DefaultFormatter,
	}
}

func (h *BaseHandler) SetLevel(l Level) {
	h.Level = l
}

func (h *BaseHandler) SetFormatter(f Formatter) {
	h.Formatter = f
}

func (h *BaseHandler) FilterAndFormat(rec *Record) string {
	if h.Level >= rec.Level {
		return h.Formatter.Format(rec)
	}
	return ""
}

///////////////////
//               //
// WriterHandler //
//               //
///////////////////

// WriterHandler is a handler implementation that writes the logging output to a io.Writer.
type WriterHandler struct {
	*BaseHandler
	w        io.Writer
	Colorize bool
}

func NewWriterHandler(w io.Writer) *WriterHandler {
	return &WriterHandler{
		BaseHandler: NewBaseHandler(),
		w:           w,
	}
}

func (b *WriterHandler) Handle(rec *Record) {
	message := b.BaseHandler.FilterAndFormat(rec)
	if message == "" {
		return
	}
	if b.Colorize {
		b.w.Write([]byte(fmt.Sprintf("\033[%dm", LevelColors[rec.Level])))
	}
	fmt.Fprint(b.w, message)
	if b.Colorize {
		b.w.Write([]byte("\033[0m")) // reset color
	}
}

func (b *WriterHandler) Close() {}

//////////////////
//              //
// MultiHandler //
//              //
//////////////////

// MultiHandler sends the log output to multiple handlers concurrently.
type MultiHandler struct {
	handlers []Handler
}

func NewMultiHandler(handlers ...Handler) *MultiHandler {
	return &MultiHandler{handlers: handlers}
}

func (b *MultiHandler) SetFormatter(f Formatter) {
	for _, h := range b.handlers {
		h.SetFormatter(f)
	}
}

func (b *MultiHandler) SetLevel(l Level) {
	for _, h := range b.handlers {
		h.SetLevel(l)
	}
}

func (b *MultiHandler) Handle(rec *Record) {
	wg := sync.WaitGroup{}
	wg.Add(len(b.handlers))
	for _, handler := range b.handlers {
		go func(handler Handler) {
			handler.Handle(rec)
			wg.Done()
		}(handler)
	}
	wg.Wait()
}

func (b *MultiHandler) Close() {
	wg := sync.WaitGroup{}
	wg.Add(len(b.handlers))
	for _, handler := range b.handlers {
		go func(handler Handler) {
			handler.Close()
			wg.Done()
		}(handler)
	}
	wg.Wait()
}
