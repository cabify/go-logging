package log

import (
	"fmt"
	"io"
	"strings"
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
	NOCOLOR = -1
)

var LevelColors = map[Level]Color{
	CRITICAL: MAGENTA,
	ERROR:    RED,
	WARNING:  YELLOW,
	NOTICE:   GREEN,
	INFO:     NOCOLOR,
	DEBUG:    BLUE,
}

// WriterHandler is a handler implementation that writes the logging output to a io.Writer.
type WriterHandler struct {
	*BaseHandler
	w        io.Writer
	m        sync.Mutex
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
	if !strings.HasSuffix(message, "\n") {
		message += "\n"
	}
	if b.Colorize && LevelColors[rec.Level] != NOCOLOR {
		message = fmt.Sprintf("\033[%dm%s\033[0m", LevelColors[rec.Level], message)
	}
	b.m.Lock()
	fmt.Fprint(b.w, message)
	b.m.Unlock()
}

func (b *WriterHandler) Close() {}
