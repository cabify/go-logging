package log

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultFormatterHasRightDateFormat(t *testing.T) {
	// Mon, 11 Jun 2018 12:35:18.123 UTC
	ts := time.Unix(1528713318, 123000000)
	rec := Record{
		Level:   INFO,
		Time:    ts,
		Message: "Hello World!",
	}
	line := DefaultFormatter.Format(&rec)
	assert.Equal(t, "2018-06-11 12:35:18.123 [] INFO     Hello World!", line)
}
