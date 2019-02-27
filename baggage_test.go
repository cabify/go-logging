package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	t.Run("has correct length", func(t *testing.T) {
		id := NewID()
		assert.Equal(t, loggerRequestIDLength, len(id))
	})
}
