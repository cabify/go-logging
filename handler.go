package log

import (
	"fmt"
	"io"
	"sync"
)

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
