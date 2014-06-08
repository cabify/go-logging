// +build darwin freebsd linux netbsd openbsd

package log

func init() {
	StdoutHandler.Colorize = true
	StderrHandler.Colorize = true
}
