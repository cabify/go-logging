package main

import (
	"fmt"

	"github.com/cenkalti/log"
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
	l2.SetHandler(&MyHandler{})

	l2.Debug("Debug")
	l2.Info("Info")
	l2.Notice("Notice")
	l2.Warning("Warning")
	l2.Error("Error")
	l2.Critical("Critical")
}

type MyHandler struct {
	log.BaseHandler
}

func (h *MyHandler) Handle(rec *log.Record) {
	fmt.Print(rec.Message)
}

func (h *MyHandler) Close() {
}
