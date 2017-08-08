package main

import (
	"fmt"

	"github.com/cabify/go-logging"
)

func main() {

	// Default logger
	log.Debug("Debug")
	log.Info("Info")
	log.Notice("Notice")
	log.Warning("Warning")
	log.Error("Error")
	log.Critical("Critical")

	// Custom logger with default handler
	l := log.NewLogger("test")

	l.Debug("Debug")
	l.Info("Info")
	l.Notice("Notice")
	l.Warning("Warning")
	l.Error("Error")
	l.Critical("Critical")

	// Custom logger with custom handler
	l2 := log.NewLogger("test2")
	h := NewMyHandler("!!!")
	h.SetLevel(log.WARNING)
	l2.SetHandler(h)

	l2.Debug("Debug")
	l2.Info("Info")
	l2.Notice("Notice")
	l2.Warning("Warning")
	l2.Error("Error")
	l2.Critical("Critical")
}

// Adds prefix to log messages
type MyHandler struct {
	*log.BaseHandler
	prefix string
}

func NewMyHandler(prefix string) *MyHandler {
	return &MyHandler{
		BaseHandler: log.NewBaseHandler(),
		prefix:      prefix,
	}
}

func (h *MyHandler) Handle(rec *log.Record) {
	message := h.BaseHandler.FilterAndFormat(rec)
	if message == "" {
		return
	}
	fmt.Println(h.prefix, message)
}

func (h *MyHandler) Close() error { return nil }
