package log

import (
	"fmt"
	"io"
	"sync"
)

type Color int

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

var LevelColors = map[Level]Color{
	CRITICAL: MAGENTA,
	ERROR:    RED,
	WARNING:  YELLOW,
	NOTICE:   GREEN,
	INFO:     WHITE,
	DEBUG:    BLUE,
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
	if rec.Level > h.Level {
		return ""
	}
	return h.Formatter.Format(rec)
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