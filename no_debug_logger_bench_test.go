package log

import (
	"testing"
)

/*
$ go test -bench=. ./...
     | goos: darwin
     | goarch: amd64
     | pkg: github.com/cabify/go-logging
     | BenchmarkPlainLoggerDebugSpeed-4      	30000000	        45.9 ns/op
     | BenchmarkNoDebugLoggerDebugSpeed-4    	2000000000	         0.33 ns/op
     | BenchmarkPlainLoggerDebugfSpeed-4     	20000000	        80.8 ns/op
     | BenchmarkNoDebugLoggerDebugfSpeed-4   	50000000	        36.5 ns/op
*/

type someStruct struct {
	value float64
}

var (
	structValue           = someStruct{value: 10.0}
	plainLoggerInstance   = NewLogger("test")
	noDebugLoggerInstance = NoDebugLogger{Logger: plainLoggerInstance}
)

func BenchmarkPlainLoggerDebugSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainLoggerInstance.Debug("Something")
	}
}

func BenchmarkNoDebugLoggerDebugSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noDebugLoggerInstance.Debug("Something")
	}
}

func BenchmarkPlainLoggerDebugfSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainLoggerInstance.Debugf("Something %v, %d", structValue, i)
	}
}

func BenchmarkNoDebugLoggerDebugfSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noDebugLoggerInstance.Debugf("Something %v, %d", structValue, i)
	}
}
