package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatterHasRightDateFormat(t *testing.T) {
	ts := time.Date(2018, 6, 11, 12, 35, 18, 123000000, time.Local)
	rec := Record{
		Level:   INFO,
		Time:    ts,
		Message: "Hello World!",
	}
	line := DefaultFormatter.Format(&rec)
	assert.Equal(t, "2018-06-11 12:35:18.123 [] INFO     Hello World!", line)
}
