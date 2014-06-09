// +build darwin freebsd linux netbsd openbsd

package log

func init() { stderrHandler.Colorize = true }
