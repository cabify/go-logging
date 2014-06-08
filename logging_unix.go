// +build darwin freebsd linux netbsd openbsd

package log

func init() { strerrHandler.Colorize = true }
